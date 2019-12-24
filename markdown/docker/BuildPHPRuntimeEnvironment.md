# ECS搭建Docker之PHP运行环境

## 使用docker逐一构建

### Step.1 拉取镜像

```shell
$ docker pull php:7.2-fpm
$ docker pull nginxf
$ docker pull mysql:5.7f
$ docker pull redis:3.2f
```

使用`docker images`查看已拉取的所有镜像

```shell
$ docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
php                 7.2-fpm             4f240ecdaf51        2 days ago          398MB
mysql               5.7                 1e4405fe1ea9        4 weeks ago         437MB
nginx               latest              231d40e811cd        4 weeks ago         126MB
redis               3.2                 87856cc39862        14 months ago       76MB
```

好了，该有的镜像都已经有了，接下来开始构建各个镜像所对应的容器。

### Step.2 拉取完成镜像后运行并构建容器

关于Docker选项的简要说明

| 选项                | 功能描述                                                     |
| ------------------- | ------------------------------------------------------------ |
| `-i`                | 表示允许我们对容器进行操作                                   |
| `-t`                | 表示在新容器内指定一个为终端                                 |
| `-d`                | 表示容器在后台执行                                           |
| `/bin/bash`         | 这将在容器内启动bash shell                                   |
| `-p`                | 为宿主机和容器创建端口映射                                   |
| `--name`            | 为容器指定一个名字                                           |
| `-v`                | 将容器内路径挂载到宿主机路径(`e.g:-v HostPath:ContainerPath`)  |
| `--privileged=true` | 给容器特权,在挂载目录后容器可以访问目录以下的文件或者目录    |
| `--link`            | 可以用来链接2个容器，使得源容器（被链接的容器）和接收容器（主动去链接的容器）之间可以互相通信，解除了容器之间通信对容器IP的依赖 |
| `--restart always`  | 重启Docker服务时总是自动运行 |
| `-e`                | 设置环境变量 |

#### 运行并构建mysql容器

关于MYSQL镜像的详细使用情况请参阅官方发布的[`mysql image`](https://hub.docker.com/_/mysql/)

```shell
$ docker run -d --name mysql --restart always -p 3306:3306 -e MYSQL_ROOT_PASSWORD=database-secret -v /data/database/mysql:/var/lib/mysql mysql:5.7
```

`MYSQL_ROOT_PASSWORD=database-secret` 给mysql数据库root用户设置初始登录密码，并且登录密码为你设置的`database-secret`；`-v` 指定将容器目录 `/var/lib/mysql` 挂载到宿主机 `/data/database/mysql`。

#### 运行并构建redis容器


关于该Redis镜像的详细使用情况请参阅官方发布的[`redis image`](https://hub.docker.com/_/redis/)

```shell
$ docker run -d --name redis --restart always -p 6379:6379 -v /data/database/redis:/data redis:3.2 redis-server --appendonly yes --requirepass cache-secret
```

`--requirepass cache-secret` 指定此redis服务需要密码登录，并且登录密码为你设置的`cache-secret`。

#### 运行并构建PHP容器

```shell
$ docker run -d --name php --restart always -p 9000:9000 -v /server/www:/var/www/html --link mysql:mysql --link redis:redis --privileged=true php:7.2-fpm
```

#### 运行并构建NGINX容器

```shell
$ docker run -d --name nginx --restart always -p 80:80 -v /server/www:/usr/share/nginx/html -v /server/nginx:/etc/nginx -v /server/logs/nginx.logs:/var/log/nginx --link php:php --privileged=true nginx
```

