# CentOS7开放端口号

## 加入开放端口到配置文件

```shell
$ firewall-cmd --zone=public --add-port=80/tcp --permanent
```

- `--zone=public`:添加时区
- `--add-port=80/tcp`:添加端口
- `--permanent`:永久生效

## 加载防火墙新配置文件

以 root 身份输入以下命令，重新加载防火墙，并不中断用户连接，即不丢失状态信息.

```shell
$ firewall-cmd --reload
```

## 目录
[BACK](../../README.md)