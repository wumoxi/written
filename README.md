## 小笔头开发日志

记录在学习或工作中使用过的相关技术，以防止岁月的冲击！

### 支付相关

- [Yii2接入PayPal支付](markdown/pay/yii2_join_up_paypal.md)

### URL去重

- 哈希表
- 计算MD5等哈希，再存哈希表
- 使用bloom filter多重哈希结构
- 使用Redis等KEY-VALUE存储系统实现分布式去重

### CentOS7开放端口号

#### 加入开放端口到配置文件

```shell
$ firewall-cmd --zone=public --add-port=80/tcp --permanent
```

- --zone=public 添加时区
- --add-port=80/tcp 添加端口
- --permanent 永久生效

#### 加载防火墙新配置文件(以 root 身份输入以下命令，重新加载防火墙，并不中断用户连接，即不丢失状态信息.)

```shell
$ firewall-cmd --reload
```
