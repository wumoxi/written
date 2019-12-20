# 开机启动容器

运行容器时添加`--restart`参数:
```shell
$ docker run --restart=always -d -p 80:80 nginx
```

如果容器已经启动过了，但是现在停止了退出了，可以使用 `update` 命令，进行对容器进行更新参数:
```shell
$ docker update --restart=always <ContainerID>
```

## 目录
[BACK](../../README.md)