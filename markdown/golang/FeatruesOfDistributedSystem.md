# 分布式系统特点

- 👉多个节点
- 👉消息传递
- 👉完成特定需求

## 多个节点

- 容错性
- 可扩展性(性能)
- 固有分布性

## 消息传递

- 节点具有私有存储，节点与节点之间通过消息传递进行通信
- 易于开发
- 可扩展性(功能)
- 对比：并行计算

## 消息传递的方法

- REST
- RPC
- 中间件

## 一般消息传递的方法

- 对外：REST
- 模块内部：RPC
    
- 模块之间：中间件，REST

## RPC
- jsonRPC
- gRPC
- Thrift

## 中间件
- RabbitMQ
- ActiveMQ
- Kafka
- Redis

## 分布式架构VS微服务架构

- 分布式：指导节点之间如何通信
- 微服务：鼓励按业务划分模块
- 微服务架构通过分布式架构来实现

## 多层架构VS微服务架构

- 微服务架构相对多层架构具有更多的 "服务"
    - 服务之间耦合度很松
    - 每个服务有自己独立的代码
    - 每个服务有自己独立的测试流程
    - 每个服务有自己独立的上线流程
    - 每个服务都可以不一样
    - 每个服务在不同时间点都可以发布不同的版本
- 微服务通常需要配合自动化测试，部署，服务发现等
- 目前我们倾向于微服务架构


## 目录
[Back](../../README.md)
