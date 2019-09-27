<?php

// namespace common\helpers;

use Yii;

class DebugHelper
{
    /**
     * 输出调试日志到文件
     *
     * @param $filename
     * @param $data
     * @param string $info
     * @return string
     */
    public static function log($filename, $data, $info = '')
    {
        $message = [];
        $message['date_time'] = TimeHelper::current();
        if ($info) {
            $message['info'] = $info;
        }
        $prefix = implode("|", $message) . "|";
        $data = $prefix . json_encode($data, JSON_UNESCAPED_UNICODE);
        $filename = Yii::$app->basePath . '/../' . $filename . '.debug.log';
        $data = $data . PHP_EOL;
        file_put_contents($filename, $data, FILE_APPEND);
    }
}