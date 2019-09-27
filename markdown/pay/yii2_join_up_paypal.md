## Yii2接入PayPal支付

做过跨境电商开发者或者在海外购物网站逛过的朋友都应该知道PayPal，类似于国内流行的支付宝和微信支付，都属于第三方支付平台。在做外贸支付的时候不得不提到的就是PayPal了，毕竟它在国外的市场占有率摆在哪里，当然据做外贸的朋友介绍在东南亚某些国家和地区支付宝的市场份额反超了PayPal。

### 背景介绍

由于在国内，做的项目需要使用PayPal的并不多，就算有时候需要做外贸项目，很多选择第四方集成。据做外贸的朋友介绍，之所以选择第四方支付是因为虽然在海外PayPal的市场占有率是最高的，然而这种在线支付的普及率远不如国内。而如果分别对接各种第三方支付或者银行支付，将需要对接很多平台，开发成本太高；而做的项目服务不在乎多一点支付的手续费。比如PayPal的手续费是3.9%+0.3美元（每笔），而第四方的相对高个1%左右；从中也可以看得出虽然我们经常吐槽支付宝和微信手续费是暴利，但从此看来国内的支付公司很良心了。

好了，步入正题，对接PayPal支付有多种支付方式，比如按钮支付之类的，而今天我们要跟大家介绍的就是基于PayPal官方提供的PayPal-PHP-SDK进行支付（也就是REST API Samples的方式）的简洁版指南；希望帮助大家能花尽可能少的时间快速实现支付对接。

### 前期准备

其实跟微信支付宝类似，就是注册帐号然后成为开发者；当然PayPal跟支付宝一样有沙箱模式，因此在开发阶段我们可以使用沙箱帐号来做开发，这样可以不用签约也可以进行开发工作。主要步骤如下：

#### 1、首先去官网注册一个paypal账号。

#### 2、申请完毕并登录，进入开发者中心（https://developer.paypal.com）。可以进入沙箱帐号：

![进入沙箱帐号](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20190927_10.png)

即可看到你申请账号时自动分配的两个测试账号，账号类别分别是：卖家帐号和买家帐号，默认就有5000.00 USD美元，不过跟支付宝一样，都是虚拟的用于测试，也可以手工修改余额。

![沙箱测试账号列表](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20190927_11.png)

> **注意**

这个两个测试账号你可以选择修改密码也可以选择不修改密码，不修改密码，可以使用这两个账号自动分配时系统自动生成的密码，都仅仅只是在开发测试阶段使用。后台实现接入时说具体怎么用。

#### 3、创建APP

进入我的应用程序和证书（`My Apps & Credentials`）创建APP（不是app客户端，意思是应用，跟支付宝开放平台的应用概念类似）。

![进入我的应用程序和证书](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20190927_9.png)

然后点击 `sandbox` 下的 `REST API apps` 栏目下面的 `Create App` 按钮，进行创建一个APP。

![进行创建一个APP](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20190927_12.png)

填写一个APP名称，然后选择一个测试账户作为此APP绑定的账号，如果你在上一步没有申请新的测试账号（在沙箱中也是可以另外创建测试的卖家帐号和买家帐号的，比如创建一个美国地区的卖家帐号和一个日本地区的买家帐号），那么这里默认就是选择了卖家帐号。如果不是可以自行确认。

![进行创建一个APP](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20190927_14.png)

点击 `Create App` 即可看到创建成功的应用 `Client ID` 和 `Secret`。

![创建APP成功后信息](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20190927_15.png)

#### 4、沙箱及线上模式

paypal的沙箱模式测试环境域名为sandbox.paypal.com，正式域名为www.paypal.com。如果是使用官方的SDK，那么直接设置mode为sandbox就是沙箱模式，而设置为live，也就是线上的意思。

#### 5、下载对应SDK

##### 官方SDK

官方提供了web开发的SDK，包括Java、.net、node、php、Python、Ruby等，也包含移动客户端开发的ios和安卓的。这里的指南是以php为例。

官方SDK地址为：https://developer.paypal.com/docs/api/rest-sdks/。这个地址包含所有SDK的下载入口，都是托管到github上面的。

##### 使用`composer` 进行安装

只需要执行 `composer require "paypal/rest-api-sdk-php:*"` 即可！

#### 6、设置异步通知地址

推荐的方式是在获取应用的`Client ID`和`Secret`的页面下面部分可以设置回调通知的url，如图所示。也可以在卖家帐号中设置PIN的url，经过测试，如果设置了这个就以PIN为准，而且两个地方paypal异步通知的数据包参数并不一样。关于paypal异步通知设置后续如果有机会跟大家详细介绍。


### 接入过程

#### 具体文件目录结构

```shell
$ tree -L 3
├── common
│   ├── helpers
│   │   ├── PaypalHelper.php
├── frontend
│   ├── controllers
│   │   └── UserController.php
```

根据你现在项目需求及自行调整

> PaypalHelper.php

```php
<?php

namespace common\helpers;

use PayPal\Auth\OAuthTokenCredential;
use PayPal\Rest\ApiContext;

class PaypalHelper
{
    /**
     * Client ID
     *
     * @var string
     */
    private static $clientId = 'Your project Client ID';

    /**
     * Secret
     *
     * @var string
     */
    private static $clientSecret = 'Your project Secret';

    /**
     * 构建前端请求表单
     *
     * @param $url
     * @return string
     */
    public static function buildRequestForm($url)
    {
        $html = "<form id='paypal-submit' name='paypal-submit' action='" . $url . "' method='GET'>";
        $html .= "<input type='submit' value='ok' style='display:none;'>";
        $html .= "</form>";
        $html .= "<script>document.forms['paypal-submit'].submit();</script>";
        return $html;
    }

    /**
     * Helper method for getting an APIContext for all calls
     * @param string $clientId Client ID
     * @param string $clientSecret Client Secret
     * @return ApiContext
     */
    public static function getApiContext()
    {
        // #### SDK configuration
        // Register the sdk_config.ini file in current directory
        // as the configuration source.
        if (!defined("PP_CONFIG_PATH")) {
            define("PP_CONFIG_PATH", __DIR__);
        }

        // ### Api context
        // Use an ApiContext object to authenticate
        // API calls. The clientId and clientSecret for the
        // OAuthTokenCredential class can be retrieved from
        // developer.paypal.com
        $apiContext = new ApiContext(
            new OAuthTokenCredential(
                self::$clientId,
                self::$clientSecret
            )
        );

        // Comment this line out and uncomment the PP_CONFIG_PATH
        // 'define' block if you want to use static file
        // based configuration
        $apiContext->setConfig(
            array(
                'mode' => 'sandbox',
                'log.LogEnabled' => true,
                'log.FileName' => '../PayPal.log',
                'log.LogLevel' => 'DEBUG', // PLEASE USE `INFO` LEVEL FOR LOGGING IN LIVE ENVIRONMENTS
                'cache.enabled' => true,
            )
        );

        // Partner Attribution Id
        // Use this header if you are a PayPal partner. Specify a unique BN Code to receive revenue attribution.
        // To learn more or to request a BN Code, contact your Partner Manager or visit the PayPal Partner Portal
        // $apiContext->addRequestHeader('PayPal-Partner-Attribution-Id', '123123123');
        return $apiContext;
    }
}
```

> UserController.php

```php
<?php

namespace frontend\controllers;

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
```

#### 1、创建一个支付，发送到paypal服务器用于获取用户授权url

在上面的 `UesrController.php` 文件中可以看到 `actionPaypal` 方法定义，访问 `http://www.shop.test/user/paypal` 即从paypal服务端获取的用户授权url，注意这并不是类似支付宝那样的支付url，仅仅只是获取用户授权url。

![从paypal服务端获取的用户授权url](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20190927_18.png)

其中在发起获取的用户授权url的时候除了设置商品及金额信息外，还需要设置两个url，一个是支付url，另外一个是取消支付的url。这里的url并不是类似支付宝支付的同步回调地址，支付宝的同步回调地址访问前实际上支付已经完成了，而这里是需要到达支付页面完成支付的。

在上面的 `UesrController.php` 文件中可以看到 `actionPaypal` 方法定义中可以看到调用了方法 `getPaypalRedirectUrls`:

```php
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
```

它返回用户确认支付时往哪跳，取消支付时又往哪跳的具体地址！

> ***注意***

这个地方一定要注意了，这也是跟支付宝支付一个不太一样的地方。支付宝的只要获取支付url，然后去支付宝网站里面就可以完成支付，然后异步通知。而paypal首先去其官网仅仅是获取用户的授权而已，最终支付还是要回到自己的网站再一次请求paypal的支付接口。


#### 2、请求paypal用户授权url

请求paypal用户授权url，进入paypal去登录买家帐号进行用户授权。如果没有帐号需要先登录买家帐号，如果已经登录了并且选择支付类型（余额或者信用卡），就会显示如下的界面：

![请求paypal用户授权url](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20190927_20.png)

点击 `Continue` 进行确认支付, 这时它将会跳转到之前我们获取用户授权url时，设置的支付跳转地址，也就是 `http://www.shop.test/user/paypal-payment?success=true` 完成支付！

![确认支付](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20190927_24.png)

方法 `actionPaypalPayment`，主要完成的工作就是根据请求参数 `paymentId`，完成支付

```php
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
```

#### 关于支付登录说明

开发测试阶段，登录支付授权时只需要使用沙箱账户列表的个人账户即可！

![个人账户](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20190927_26.png)

具体的密码可以点击 `View/Edit Account` 查看！

![账号密码](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20190927_28.png)

### 注意事项

paypal的支付流程跟国内的支付宝支付还是有区别的，总体而言我感觉不管是支付方式还是开发体验还是使用体验都不如国内的支付宝和微信支付。下面把过程中几个可能被大家误解的地方说明下，防止大家入坑。

1. 支付的基本流程是：创建一个支付，发送到paypal服务端并返回一个用户授权地址(在客户端，即我们自己的服务端，并设置一个支付)–>跳转到用户授权地址（paypal网站）–>用户授权完毕（paypal网站，用户登录帐号并同意支付）–>paypal返回到客户端设置的execute地址（这个地址是第一步设置的，在客户端，即支付跳转地址），付款实现。
2. 在上述流程的第一步设置的url，并不是类似微信或支付宝那样的回调url，而是后续返回到客户端完成支付的url。
3. paypal貌似不支持人民币支付，也就是在上述步骤中第一步创建支付的时候设置货币类型，可以设置美元等其他货币，但不能设置人民币（CNY），因此如果是系统货币单位为人民币需要在服务端转化为美元然后再创建支付。据说是由于paypal在中国大陆地区没有获得支付牌照的原因。













