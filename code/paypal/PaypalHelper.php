<?php

// namespace common\helpers;

use PayPal\Auth\OAuthTokenCredential;
use PayPal\Rest\ApiContext;

class PaypalHelper
{
    /**
     * Client ID
     *
     * @var string
     */
    private static $clientId = 'ARgKRjRw3dtjv-_CGe0RbyIv9mjkCsMQIekDGj7LUdwQdnf9HvccCpNVIHCxlbu6q5G1PXAee5RKQop3';

    /**
     * Client Secret
     *
     * @var string
     */
    private static $clientSecret = 'EPIjPy10jeVsgmmoAqiRrzBQ4KU4senjMbKAOP7U-GpvBhYMTGHwTh5U61HA-eYAx3ErvmmsTocr1BzL';

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