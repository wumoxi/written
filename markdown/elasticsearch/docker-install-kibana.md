# Docker下安装Kibana

## 拉取镜像

```shell
$ docker pull docker.elastic.co/kibana/kibana:7.5.1
```

## 在Docker上运行Kibana进行开发环境

```shell
$ docker run --restart always -d -p 5601:5601 --name kibana docker.elastic.co/kibana/kibana:7.5.1
```

查看容器运行情况

```shell
$ docker ps
CONTAINER ID        IMAGE                                                 COMMAND                  CREATED             STATUS              PORTS                                            NAMES
2d7436dddedc        docker.elastic.co/kibana/kibana:7.5.1                 "/usr/local/bin/dumb…"   3 seconds ago       Up 3 seconds        0.0.0.0:5601->5601/tcp                           kibana
3aabe7452397        docker.elastic.co/elasticsearch/elasticsearch:7.5.0   "/usr/local/bin/dock…"   3 hours ago         Up 2 hours          0.0.0.0:9200->9200/tcp, 0.0.0.0:9300->9300/tcp   es
```

你会发现`kibana`容器服务已经起来了，很好，通过`curl`访问一个试试看

```shell
$ curl -XGET http://localhost:5601
Kibana server is not ready yet
```

不好意思，它居然提示`kibana服务器还没有准备好`，什么鬼，哦，想起来了，我们的elastic好像是设置过密码了，那就想办法在启动kibana时，把已经设置过的elasticsearch密码添加到kibana服务的配置文件内，让kibana服务去加载这些配置项，好了，可以将kibana服务容器的相关配置文件拷贝到宿主机，编辑这个配置文件，然后在启动kibana服务容器时将这个宿主机的配置文件挂载到服务容器配置文件夹内，那这件事情，也就算是完成了

### 拷贝kibana服务容器配置文件到宿主机

对于确认这个kibana配置文件的具体位置，有几种方法可以做到

1. 搜索引擎搜索、
2. [通过官方文档](https://www.elastic.co/guide/en/kibana/current/docker.html)、
3. 直接查看镜像详情
    - 直接通过执行`docker image inspcet image:[tag]`命令，即可查看镜像的详细信息，在这个详细信息中可以找到`Config.WorkingDir`JSON数据项，这个值一般类似于 `/usr/share/kibana`, 这就是一个服务容器的工作目录，其配置文件一般都会在这里，不过话又说回来，一个没有运行任何容器的镜像，也谈不上从容器拷贝文件到宿主机。
4. 直接查看容器详情
    - 个人认为这种方式是逼格最高的，直接通过执行`docker inspcet (CintainerID)`命令，即可查看运行容器的详细信息，在这个详细信息中可以找到`Config.WorkingDir`JSON数据项，这个值一般类似于 `/usr/share/kibana`, 这就是一个服务容器的工作目录，其配置文件一般都会在这里。

#### 进入kibana服务容器的内部查看具体的配置文件路径

```shell
$ docker exec -it kibana bash
```

查看当前路径你会发现一个很有意思的事情

```shell
bash-4.2$ pwd
/usr/share/kibana
```

对的，正如你所看到一样，从宿主机进行到服务容器运行环境内部默认就是进入到工作区目录，这很好，那就列出所有文件呗

```shell
bash-4.2$ ll
total 1524
-rw-rw-r--    1 kibana root   13675 Dec 16 23:46 LICENSE.txt
-rw-rw-r--    1 kibana root 1453580 Dec 16 23:46 NOTICE.txt
-rw-rw-r--    1 kibana root    4048 Dec 16 23:46 README.txt
drwxrwsr-x    2 kibana root    4096 Dec 17 00:01 bin
drwxrwsr-x    5 kibana root    4096 Dec 17 00:01 built_assets
drwxrwsr-x    1 kibana root    4096 Dec 17 00:06 config
drwxrwsr-x    2 kibana root    4096 Dec 16 23:46 data
drwxrwsr-x    6 kibana root    4096 Dec 17 00:01 node
drwxrwsr-x 1206 kibana root   36864 Dec 17 00:01 node_modules
drwxrwsr-x    1 kibana root    4096 Dec 17 00:01 optimize
-rw-rw-r--    1 kibana root     738 Dec 16 23:46 package.json
drwxrwsr-x    2 kibana root    4096 Dec 16 23:46 plugins
drwxrwsr-x   11 kibana root    4096 Dec 17 00:01 src
drwxrwsr-x    2 kibana root    4096 Dec 17 00:01 webpackShims
drwxrwsr-x    5 kibana root    4096 Dec 17 00:01 x-pack
```

很好，从当前目录文件列表中可以看到有一个 `config` 目录，列出来看看呗

```shell
bash-4.2$ ll config
total 4
-rw-rw-r-- 1 kibana root 240 Dec 17 00:01 kibana.yml
```

好的果不其然，它里面有一个 `kibana.yml` 文件，好的就是它了，不然也可以看看具体文件内容

```shell
bash-4.2$ cat config/kibana.yml 
```

```yaml
bash-4.2$ cat config/kibana.yml 
#
# ** THIS IS AN AUTO-GENERATED FILE **
#

# Default Kibana configuration for docker target
server.name: kibana
server.host: "0"
elasticsearch.hosts: [ "http://elasticsearch:9200" ]
xpack.monitoring.ui.container.elasticsearch.enabled: true
```

#### 创建一个宿主机文件夹用于存放kibana配置文件

```shell
$ mkdir -p /server/docker/kibana/config
```

#### 拷贝kibana服务容器配置文件到宿主机

```shell
$ docker cp kibana:/usr/share/kibana/config/kibana.yml /server/docker/kibana/config
```

### 修改宿主机kibana配置文件

```shell
$ vim /server/docker/kibana/config/kibana.yml
```

添加两行内容

```yaml
elasticsearch.username: "elastic"
elasticsearch.password: "es-secret"
```

关于如何设置elasticsearch密码，请参考[「Docker下安装ElasticSearch#设置elasticsearch密码」](./docker-install-elasticsearch.md#%E8%AE%BE%E7%BD%AEelasticsearch%E5%AF%86%E7%A0%81)

```yaml
#
# ** THIS IS AN AUTO-GENERATED FILE **
#

# Default Kibana configuration for docker target
server.name: kibana
server.host: "0"
elasticsearch.hosts: [ "http://elasticsearch:9200" ]
xpack.monitoring.ui.container.elasticsearch.enabled: true
elasticsearch.username: "elastic"
elasticsearch.password: "es-secret"
```

### 删除现已启动但配置不正确的kibana服务容器

```shell
$ docker rm -f kibana
```

### 启动新的kibana服务容器

```shell
$ docker run --restart always -d -p 5601:5601 -v /server/docker/kibana/config:/usr/share/kibana/config --link es:elasticsearch --name kibana docker.elastic.co/kibana/kibana:7.5.1
```

- `-v /server/docker/kibana/config:/usr/share/kibana/config` 将宿主机`/server/docker/kibana/config`目录挂载到服务容器环境`/usr/share/kibana/config`目录。

- `--link es:elasticsearch` 可以用来链接2个容器，使得源容器（es）和接收容器（kibana）之间可以互相通信，解除了容器之间通信对容器IP的依赖，`--link 源容器ID或容器名称:接收容器在内部使用的容器名称`，还记不记得`/server/docker/kibana/config/kibana.yml`配置文件中有这么一行配置`elasticsearch.hosts: [ "http://elasticsearch:9200" ]`，对的`http://elasticsearch:9200`就是对应于`接收容器在内部使用的容器名称`，它和容器名称为`es`容器对应并关联。

### 测试kibana容器是否可用

```shell
$ curl -XGET http://localhost:5601
```

很好，没有任何，输出也没有错误提示，说明它运行起来了，不信的话可以通过测试进行测试

![测试kibana容器是否可用](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191223_111.png)

果然就是这个样子，**注意**这个`ali.test`域名是我绑定在本地Mac中的一个host配置，IP就是对应于阿里云ECS云服务器IP，输入用户名和密码进行登录认证

![测试kibana容器是否可用](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191223_112.png)

这是一个初始化页面

![测试kibana容器是否可用](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191223_115.png)

上图展示了可以对Index索引库进行管理，kibana的功能很丰富，你可以自己学习怎么使用这些功能！祝好！

## 目录
[Back](../../README.md)




































