# CentOS7查看端口占用

[netstat](https://man.linuxde.net/netstat)命令用来打印Linux中网络系统的状态信息，可让你得知整个Linux系统的网络情况。

## 查看系统TCP端口占用情况

```shell
$ netstat -lntp
```

## 查看系统UDP端口占用情况

```shell
$ netstat -lnup
```

## 目录
[BACK](../../README.md)