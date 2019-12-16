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

## [使用Docker Compose启动多节点集群](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html#docker-compose-file)