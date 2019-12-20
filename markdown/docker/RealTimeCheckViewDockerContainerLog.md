# 实时查看docker容器日志

```shell
$ docker logs -f -t --tail (行数) (容器ID)
```

如实时查看ContinerID为 `1707a34511f4` 的容器最后100行日志，可以执行：

```shell
$ docker logs -f -t --tail 100 1707a34511f4
```