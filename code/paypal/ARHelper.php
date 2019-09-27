<?php

// namespace common\helpers;

use common\models\Order;
use common\models\SubOrder;

class ARHelper
{
    /**
     * 获取ActiveQuery执行SQL
     *
     * @param \yii\db\ActiveQuery $query
     * @return string
     */
    public static function getSQL($query)
    {
        return $query->createCommand()->getRawSql();
    }
}