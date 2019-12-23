# Docker下安装ElasticSearch

## 拉取镜像

获取适用于Docker的Elasticsearch镜像其实非常简单，使用 `docker pull` 命令从Elastic Docker注册库直接拉取即可。

```shell
$ docker pull docker.elastic.co/elasticsearch/elasticsearch:7.5.0
```

## 使用Docker启动单个节点集群

要启动用于开发或测试的单节点Elasticsearch集群，请指定单节点发现`-e "discovery.type=single-node"`环境变量以绕过引导检查：

```shell
$ docker run --restart always -d -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" -v /usr/local/data/elasticsearch:/usr/share/elasticsearch/data --name es docker.elastic.co/elasticsearch/elasticsearch:7.5.0
```

**注意**：如果你是在linux系统运行可能会出现以下错误，并且容器启动失败

```shell
caused by: java.nio.file.accessdeniedexception: /usr/share/elasticsearch/data/nodes
```

这是因为你挂载到宿主机Host的目录，在这里也就是 `/usr/local/data/elasticsearch` 没有写入权限，需要将其文件夹的所有者和所属组进行修改，可以执行`chown -R 1000:1000 /usr/local/data/elasticsearch`即可(详细请参考[issues#21](https://github.com/elastic/elasticsearch-docker/issues/21))，重新运行容器。

## [使用Docker Compose启动多节点集群](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html#docker-compose-file)


## 设置elasticsearch密码

### Step.1 进入到容器内部环境

```shell
$ docker exec -it es bash
```

### Step.2 使用密码设置工具进行密码设置

```shell
[root@3aabe7452397 elasticsearch]# elasticsearch-setup-passwords interactive
                                   
Unexpected response code [500] from calling GET http://172.18.0.3:9200/_security/_authenticate?pretty
It doesn't look like the X-Pack security feature is enabled on this Elasticsearch node.
Please check if you have enabled X-Pack security in your elasticsearch.yml configuration file.

ERROR: X-Pack Security is disabled by configuration.
```

这里报错了，不能进行设置密码，它告诉你 `在这个Elasticsearch节点上并没有启用X-Pack安全特性。请到elasticsearch.yml文件检查X-Pack安全特性是否开启`，那么你也不知道 `elasticsearch.yml`文件的具体位置对不对，知道更好！

### Step.3 查找elasticsearch.yml的具体位置

```shell
[root@3aabe7452397 elasticsearch]# find / -name "elasticsearch.yml"
/usr/share/elasticsearch/config/elasticsearch.yml
```

从查找结果上看`elasticsearch.yml`文件的具体位置，位于`/usr/share/elasticsearch/config`文件夹，好了，来吧，修改它！

### Step.4 修改elasticsearch.yml配置文件

具体改什么呢，好了，直接使用`curl`命令访问上面报错给的地址`http://172.18.0.3:9200/_security/_authenticate?pretty`

```shell
[root@3aabe7452397 elasticsearch]# curl -XGET http://172.18.0.3:9200/_security/_authenticate?pretty
{
  "error" : {
    "root_cause" : [
      {
        "type" : "exception",
        "reason" : "Security must be explicitly enabled when using a [basic] license. Enable security by setting [xpack.security.enabled] to [true] in the elasticsearch.yml file and restart the node."
      }
    ],
    "type" : "exception",
    "reason" : "Security must be explicitly enabled when using a [basic] license. Enable security by setting [xpack.security.enabled] to [true] in the elasticsearch.yml file and restart the node."
  },
  "status" : 500
}
```

可以看到它提示我们，将`xpack.security.enabled`设置为`true`, 使用`vi`工具进行修改

```shell
[root@3aabe7452397 elasticsearch]# vi /usr/share/elasticsearch/config/elasticsearch.yml
```

`vi`工具怎么用你应该会吧，docker都能用，linux你不会，哈哈，补补linux吧，不多说！

添加这么一行 `xpack.security.enabled: true` 即可，完整的文件内容如下

```yaml
cluster.name: "docker-cluster"
network.host: 0.0.0.0
xpack.security.enabled: true
```

### Step.5 退出es内部环境

```shell
[root@3aabe7452397 elasticsearch]# exit
```

### Step.6 在宿主机Host重启elasticsearch容器

```shell
$ docker restart es
```

### Step.7 再次进入容器内部并进行修改密码

重复Step.1和Step.2，好了，进入正题，它提示如下，并且提示你 `请确认你是否要继续` 当然了，输入 `y`，开始设置相关密码，直接输入相应的密码即可(密码输入不显示任何内容，这在linux系统下涉及到密码操作的基本上都是这样子，盲输就是，这是也linux安全性的一个表现)

```shell
[root@3aabe7452397 elasticsearch]# elasticsearch-setup-passwords interactive
Initiating the setup of passwords for reserved users elastic,apm_system,kibana,logstash_system,beats_system,remote_monitoring_user.
You will be prompted to enter passwords as the process progresses.
Please confirm that you would like to continue [y/N]y


Enter password for [elastic]: 
Reenter password for [elastic]: 
Enter password for [apm_system]: 
Reenter password for [apm_system]: 
Enter password for [kibana]: 
Reenter password for [kibana]: 
Enter password for [logstash_system]: 
Reenter password for [logstash_system]: 
Enter password for [beats_system]: 
Reenter password for [beats_system]: 
Enter password for [remote_monitoring_user]: 
Reenter password for [remote_monitoring_user]: 
Changed password for user [apm_system]
Changed password for user [kibana]
Changed password for user [logstash_system]
Changed password for user [beats_system]
Changed password for user [remote_monitoring_user]
Changed password for user [elastic]
```

至此密码添加成功，包含以下用户`elastic`、`apm_system`、`kibana`、`logstash_system`、`beats_system`、`remote_monitoring_user`。

### Step.8 验证elasticsearch密码设置

不使用密码访问`localhost:9200`这个elasticsearch服务，你会得到如下如果

```shell
[root@3aabe7452397 elasticsearch]# curl localhost:9200?pretty=true
{
  "error" : {
    "root_cause" : [
      {
        "type" : "security_exception",
        "reason" : "missing authentication credentials for REST request [/?pretty=true]",
        "header" : {
          "WWW-Authenticate" : "Basic realm=\"security\" charset=\"UTF-8\""
        }
      }
    ],
    "type" : "security_exception",
    "reason" : "missing authentication credentials for REST request [/?pretty=true]",
    "header" : {
      "WWW-Authenticate" : "Basic realm=\"security\" charset=\"UTF-8\""
    }
  },
  "status" : 401
}
```

对的，就是这样子的，我们的目的就是不能让这个elasticsearch服务在网络上裸奔，要不然你的服务都成了公共的了，这就不好了，这就是给elasticsearch上了把锁，想进得提供钥匙

```shell
[root@3aabe7452397 elasticsearch]# curl -u elastic:es-secret localhost:9200?pretty=true
{
  "name" : "3aabe7452397",
  "cluster_name" : "docker-cluster",
  "cluster_uuid" : "eOQZ5e5ETHG3-IGtSOqVDA",
  "version" : {
    "number" : "7.5.0",
    "build_flavor" : "default",
    "build_type" : "docker",
    "build_hash" : "e9ccaed468e2fac2275a3761849cbee64b39519f",
    "build_date" : "2019-11-26T01:06:52.518245Z",
    "build_snapshot" : false,
    "lucene_version" : "8.3.0",
    "minimum_wire_compatibility_version" : "6.8.0",
    "minimum_index_compatibility_version" : "6.0.0-beta1"
  },
  "tagline" : "You Know, for Search"
}
```

`-u elastic:es-secret` 指定服务认证的用户名和密码。这里使用 `elastic` 这个用户进行 `elasticsearch` 服务操作。

### Step.9 修改成功退出elasticsearch内部环境

```shell
[root@3aabe7452397 elasticsearch]# exit
```

当然了，在 docker 环境下有了elasticsearch服务容器，那应该有一个更好用的管理工具才对，请右转查看[Docker下安装Kibana](DockerInstallKibana.md)

## 目录
[Back](../../README.md)