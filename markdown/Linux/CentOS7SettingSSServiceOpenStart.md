# CentOS7设置SS服务开机启动

## 配置自启动

新建启动脚本文件 `/etc/systemd/system/shadowsocks.service`，内容如下：

```shell
[Unit]
Description=Shadowsocks

[Service]
TimeoutStartSec=0
ExecStart=/usr/bin/ssserver -c /etc/shadowsocks.json

[Install]
WantedBy=multi-user.target
```

## 启动 shadowsocks 服务

执行以下命令启动 shadowsocks 服务：

```shell
$ systemctl enable shadowsocks
$ systemctl start shadowsocks
```

## 检查 shadowsocks 服务是否已成功启动

```shell
[root@localhost ~]# systemctl status shadowsocks
● shadowsocks.service - Shadowsocks
   Loaded: loaded (/etc/systemd/system/shadowsocks.service; enabled; vendor preset: disabled)
   Active: active (running) since Thu 2020-02-06 03:13:39 UTC; 1h 21min ago
 Main PID: 628 (ssserver)
    Tasks: 1 (limit: 4896)
   Memory: 23.9M
   CGroup: /system.slice/shadowsocks.service
           └─628 /usr/bin/python3.6 /usr/local/bin/ssserver -c /etc/shadowsocks/shadowsocks.json

Feb 06 04:33:19 ip-172-31-34-131.us-east-2.compute.internal ssserver[628]: 2020-02-06 04:33:19 INFO     connecting oauthaccountma>
Feb 06 04:33:24 ip-172-31-34-131.us-east-2.compute.internal ssserver[628]: 2020-02-06 04:33:24 INFO     connecting ogs.google.com>
Feb 06 04:33:24 ip-172-31-34-131.us-east-2.compute.internal ssserver[628]: 2020-02-06 04:33:24 INFO     connecting safebrowsing.g>
Feb 06 04:33:25 ip-172-31-34-131.us-east-2.compute.internal ssserver[628]: 2020-02-06 04:33:25 INFO     connecting www.gstatic.co>
Feb 06 04:33:25 ip-172-31-34-131.us-east-2.compute.internal ssserver[628]: 2020-02-06 04:33:25 INFO     connecting ssl.gstatic.co>
Feb 06 04:33:25 ip-172-31-34-131.us-east-2.compute.internal ssserver[628]: 2020-02-06 04:33:25 INFO     connecting apis.google.co>
Feb 06 04:33:26 ip-172-31-34-131.us-east-2.compute.internal ssserver[628]: 2020-02-06 04:33:26 INFO     connecting play.google.co>
Feb 06 04:33:30 ip-172-31-34-131.us-east-2.compute.internal ssserver[628]: 2020-02-06 04:33:30 INFO     connecting clients4.googl>
Feb 06 04:33:32 ip-172-31-34-131.us-east-2.compute.internal ssserver[628]: 2020-02-06 04:33:32 INFO     connecting clients1.googl>
Feb 06 04:34:09 ip-172-31-34-131.us-east-2.compute.internal ssserver[628]: 2020-02-06 04:34:09 INFO     connecting play.google.co>
lines 1-19/19 (END)
```


这个操作过程其实就是添加服务配置到 systemctl 服务管理中去，然后通过 systemctl 服务管理进行开启开机启动就是这么简单哦！