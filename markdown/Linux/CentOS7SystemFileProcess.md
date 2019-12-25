# CentOS7文件处理

## 查看文件的前几行

```shell
$ head -n number file
```

## 查看文件的后几行

```shell
$ tail -n number file
```

## 从文件后几行开始实时查看文件内容

这个很常用，一般文件都是追加写，而不是从头插入写，尤其是调试、监控时很有用

```shell
$ tail -f -n number file
```

## 目录
[BACK](../../README.md)