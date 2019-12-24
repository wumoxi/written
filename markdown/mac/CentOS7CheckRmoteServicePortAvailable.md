# Mac检测远程服务端口是否可用

检测远程服务端口是否可用，可以使用`nc`命令, 进行操作，具体命令


```shell
$ nc -z -w 1 IP PORT
```


## 安装

### Mac

```shell
brew install nc
```

### Linux

```shell
yum install -y nc
```

## 检测服务端口外部环境能否正常访问

例如要检测IP为 `119.75.216.20` 的主机端口号 `443` 是否访问正常，可以使用以下命令

```shell
$ nc -z -w 2 119.75.216.20 443
Connection to 119.75.216.20 port 443 [tcp/https] succeeded!
```