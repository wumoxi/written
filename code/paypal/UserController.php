<?php

// namespace frontend\controllers;

use common\helpers\ConstHelper;
use common\helpers\DebugHelper;
use common\helpers\PaypalHelper;
use Exception;
use PayPal\Api\Amount;
use PayPal\Api\Details;
use PayPal\Api\Item;
use PayPal\Api\ItemList;
use PayPal\Api\Payer;
use PayPal\Api\Payment;
use PayPal\Api\PaymentExecution;
use PayPal\Api\RedirectUrls;
use PayPal\Api\Transaction;
use Yii;

class UserController extends BaseController
{
    /**
     *  支付类型【paypal】
     */
    const PAYMENT_TYPE_PAYPAL = 'paypal';

    /**
     * 货币类型【USD(美元)】
     */
    const CURRENCY_TYPE_USD = 'USD';

    /**
     * Paypal支付
     *
     * @return array
     */
    public function actionPaypal()
    {
        /**
         * 获取待付款的商品列表(详细项目信息)
         * @var ItemList $item_list
         * @var Amount $amount
         */
        list($item_list, $amount) = $this->getPaypalItemsList();

        // 运费
        $shipping = 1.2;
        // 税
        $tax = 1.3;

        // 获取待付款金额(金额)
        $amount = $this->getPaypalPayAmount($amount, $shipping, $tax);

        // 获取事务定义支付的契约(透传)
        $transaction = $this->getPaypalTransaction($amount, $item_list);

        // 获取跳转地址列表(支付或取消)
        $redirect_urls = $this->getPaypalRedirectUrls();

        // 获取付款对象
        $payment = $this->getPaypalPayment($redirect_urls, $transaction);

        // 请求数据
        $request = clone $payment;

        // 获取Paypal ApiContext
        $apiContext = PaypalHelper::getApiContext();

        // 创建付款
        try {
            $payment->create($apiContext);
        } catch (Exception $ex) {
            DebugHelper::log('paypal', ['request' => $request->toJSON(), 'error' => $ex->getMessage()], '获取付款授权失败');
            return $this->renderFailJson(ConstHelper::CODE_FAIL, "Failed to obtain payment authorization", $ex->getMessage());
        }

        // 获取授权跳转地址
        $approval_url = $payment->getApprovalLink();
        return $this->renderSuccessJson(["approval_url" => $approval_url]);
        return $this->renderSuccessJson(["payment_id" => $payment->id, "html" => PaypalHelper::buildRequestForm($approval_url)]);
    }

    /**
     * 用户付款或取消
     *
     * @return array
     */
    public function actionPaypalPayment()
    {
        $params = Yii::$app->request->get();
        DebugHelper::log('paypal', $params, 'GET请求参数');

        // ### Approval Status
        // Determine if the user approved the payment or not
        if (isset($params['success']) && $params['success'] == 'true') {

            $apiContext = PaypalHelper::getApiContext();

            // Get the payment Object by passing paymentId
            // payment id was previously stored in session in
            // CreatePaymentUsingPayPal.php
            $paymentId = $params['paymentId'];
            $payment = Payment::get($paymentId, $apiContext);

            // ### Payment Execute
            // PaymentExecution object includes information necessary
            // to execute a PayPal account payment.
            // The payer_id is added to the request query parameters
            // when the user is redirected from paypal back to your site
            $execution = new PaymentExecution();
            $execution->setPayerId($params['PayerID']);

            try {
                // Execute the payment(执行付款)
                // (See bootstrap.php for more on `ApiContext`)
                $result = $payment->execute($execution, $apiContext);
                DebugHelper::log('paypal', $result->toJSON(), '执行付款');
                try {
                    // 显示按ID支付的详细信息
                    $payment = Payment::get($paymentId, $apiContext);
                } catch (Exception $ex) {
                    DebugHelper::log('paypal', $ex->getMessage(), '获取付款详情信息');
                    return $this->renderFailJson(ConstHelper::CODE_FAIL, $ex->getMessage());
                }
            } catch (Exception $ex) {
                DebugHelper::log('paypal', $ex->getMessage(), '执行付款');
                return $this->renderFailJson(ConstHelper::CODE_FAIL, $ex->getMessage());
            }
            DebugHelper::log('paypal', $payment->toJSON(), '支付');
            return $this->renderSuccessJson(json_decode($payment->toJSON(), true));
        } else {
            return $this->renderFailJson(ConstHelper::CODE_FAIL, "User Cancelled the Approval");
        }
    }

    /**
     * 获取待付款的商品列表
     *
     * @return array
     */
    private function getPaypalItemsList()
    {
        $products = [
            [
                'name' => 'Ground Coffee 40 oz',
                'currency' => 'USD',
                'quantity' => 1,
                'sku' => '123123',
                'price' => 7.5,
            ],
            [
                'name' => 'Granola bars',
                'currency' => 'USD',
                'quantity' => 5,
                'sku' => '123123',
                'price' => 2,
            ],
        ];

        $items = [];
        $amount = 0.0;
        if (!empty($products)) {
            foreach ($products as $product) {
                $item = new Item();
                // 名称
                $item->setName($product['name']);
                // 货币
                $item->setCurrency($product['currency']);
                // 数量
                $item->setQuantity($product['quantity']);
                // 商品属性
                $item->setSku($product['sku']);
                // 价格
                $item->setPrice($product['price']);
                $amount += ($product['quantity'] * $product['price']);
                array_push($items, $item);
            }
        }

        $item_list = new ItemList();
        $item_list->setItems($items);
        return array($item_list, $amount);
    }

    /**
     * 获取待付款金额具体细节
     * @param $amount
     * @param $shipping
     * @param $tax
     * @return Amount
     */
    private function getPaypalPayAmount($amount, $shipping, $tax)
    {
        // 订单总金额
        $total = $amount + $shipping + $tax;

        // 付款详情
        $detail = (new Details())
            // 运费
            ->setShipping($shipping)
            // 税
            ->setTax($tax)
            // 订单金额
            ->setSubtotal($amount);

        // 订单金额及具体详情
        return (new Amount())
            ->setTotal($total)
            ->setCurrency(self::CURRENCY_TYPE_USD)
            ->setDetails($detail);
    }

    /**
     * 获取事务定义支付的契约
     *
     * @param Amount $amount
     * @param $item_list
     * @return Transaction
     */
    private function getPaypalTransaction(Amount $amount, ItemList $item_list)
    {
        return (new Transaction())
            ->setAmount($amount)
            ->setItemList($item_list)
            ->setDescription('Payment description of shipping');
    }

    /**
     * 获取域名
     * @return string
     */
    private function getBaseUrl()
    {
        $protocol = 'http';
        if ($_SERVER['SERVER_PORT'] == 443 || (!empty($_SERVER['HTTPS']) && strtolower($_SERVER['HTTPS']) == 'on')) {
            $protocol .= 's';
        }
        $host = $_SERVER['HTTP_HOST'];
        $request = $_SERVER['PHP_SELF'];
        return dirname($protocol . '://' . $host . $request);
    }

    /**
     * 获取跳转地址列表
     *
     * @return RedirectUrls
     */
    private function getPaypalRedirectUrls()
    {
        $base_url = $this->getBaseUrl();
        return (new RedirectUrls())
            ->setReturnUrl("$base_url/user/paypal-payment?success=true")
            ->setCancelUrl("$base_url/user/paypal-payment?success=false");
    }

    /**
     * 获取付款对象
     *
     * @param RedirectUrls $redirect_urls
     * @param Transaction $transaction
     * @return Payment
     */
    private function getPaypalPayment(RedirectUrls $redirect_urls, Transaction $transaction)
    {
        return (new Payment())
            ->setIntent('sale')
            ->setPayer((new Payer())->setPaymentMethod(self::PAYMENT_TYPE_PAYPAL))
            ->setRedirectUrls($redirect_urls)
            ->setTransactions(array($transaction));
    }
}