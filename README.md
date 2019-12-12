# 小笔头开发日志

记录在学习或工作中使用过的相关技术，以防止岁月的冲击！

## 支付相关

- [Yii2接入PayPal支付](markdown/pay/yii2_join_up_paypal.md)

## URL去重

- 哈希表
- 计算MD5等哈希，再存哈希表
- 使用bloom filter多重哈希结构
- 使用Redis等KEY-VALUE存储系统实现分布式去重

## LINUX

### CentOS7开放端口号

#### 加入开放端口到配置文件

```shell
$ firewall-cmd --zone=public --add-port=80/tcp --permanent
```

- `--zone=public`:添加时区
- `--add-port=80/tcp`:添加端口
- `--permanent`:永久生效

#### 加载防火墙新配置文件(以 root 身份输入以下命令，重新加载防火墙，并不中断用户连接，即不丢失状态信息.)

```shell
$ firewall-cmd --reload
```

### 开机启动服务

```shell
$ chkconfig docker on
```
## Docker

### 开机启动容器

运行容器时添加`--restart`参数:
```shell
$ docker run --restart=always -d -p 80:80 nginx
```

如果容器已经启动过了，但是现在停止了退出了，可以使用 `update` 命令，进行对容器进行更新参数:
```shell
$ docker update --restart=always <ContainerID>
```

### 实时查看docker容器日志

```shell
$ docker logs -f -t --tail (行数) (容器ID)
```

如实时查看ContinerID为 `1707a34511f4` 的容器最后100行日志，可以执行：

```shell
$ docker logs -f -t --tail 100 1707a34511f4
```

