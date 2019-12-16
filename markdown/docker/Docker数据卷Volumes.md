# Docker数据卷Volumes

## 1、数据卷

数据卷是一个可供一个或多个容器使用的特殊目录，它绕过 UFS，可以提供很多有用的特性：

- 数据卷可以在容器之间共享和重用。
- 对数据卷的修改会立马生效。
- 对数据卷的更新，不会影响镜像。
- 卷会一直存在，直到没有容器使用。

> Top: 数据卷的使用，类似于Linux下对目录或文件进行 `挂载(mount)` 操作。

### 1.1 创建数据卷

在用 `docker run` 命令的时候，使用 `-v` 标记来创建一个数据卷并挂载到容器里。在一次 run 中多次使用可以挂载多个数据卷。

下面创建一个 web 容器，并加载一个数据卷到容器的 /usr/share/nginx/html 目录。

```shell
$ docker run -d -p 8859:80 --name nginx -v /usr/share/nginx/html nginx
```

> Top: 也可以在 `Dockerfile` 中使用 `VOLUME` 来添加一个或者多个新的卷到由该镜像创建的任意容器。

### 1.2 挂载一个主机目录作为数据卷

使用 `-v` 标记也可以指定挂载一个本地主机的目录到容器中去。

```shell
$ docker run -d -p 8859:80 --name nginx -v /usr/local/data/nginx/html:/usr/share/nginx/html nginx
```

上面的命令将加载主机的 `/usr/local/data/nginx/html` 目录到容器的 `/usr/share/nginx/html` 目录。这个功能在进行测试的时候十分方便，比如用户可以放置一些程序到本地目录中，来查看容器是否正常工作。本地目录的路径必须是绝对路径，如果目录不存在 Docker 会自动为你创建它。

> Top: Dockerfile 中不支持这种用法，这是因为 Dockerfile 是为了移植和分享用的。然而，不同操作系统的路径格式不一样，所以目前还不能支持。

Docker 挂载数据卷的默认权限是读写，用户也可以通过 `:ro` 指定为只读。加了 `:ro` 之后，就挂载为只读了。

```shell
$ docker run -d -p 8859:80 --name nginx -v /usr/local/data/nginx/html:/usr/share/nginx/html:ro nginx
```

## 2、数据卷容器

如果你有一些持续更新的数据需要在容器之间共享，最好创建数据卷容器。

数据卷容器，其实就是一个正常的容器，专门用来提供数据卷供其它容器挂载的。

首先，创建一个命名的数据卷容器 database：

```shell
$ docker create -v /usr/local/data/database:/var/lib/mysql --name database centos
```

然后，在其他容器中使用 `--volumes-from` 来挂载 database 容器中的数据卷。


```shell
$ docker run --volumes-from database --name mysql1 -p 3406:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7
```

```shell
$ docker run --volumes-from database --name mysql2 -p 3506:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7
```

还可以使用多个 `--volumes-from` 参数来从多个容器挂载多个数据卷。 也可以从其他已经挂载了数据卷的容器来挂载数据卷。

```shell
$ docker run --volumes-from mysql2 --name mysql3 -p 3506:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7
```

> Top: 使用 --volumes-from 参数所挂载数据卷的容器自己并不需要保持在运行状态。

## 目录
[Back](../../README.md)
