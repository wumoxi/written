# CentOS7之VIM快速替换所有匹配字符串


例如有一个文件 `nginx.conf` 如下所示：

```nginx
server {
    listen      80;
    server_name linkto.com www.linkto.com;
    root        /data/www/linkto/public/;
    index       index.php; 
    charset     utf-8;
    # ssl 	on;

    location / {
        try_files $uri $uri/ /index.php?$query_string;
    }

    location = /favicon.ico { access_log off; log_not_found off; }
    location = /robots.txt  { access_log off; log_not_found off; }

    access_log /data/logs/nginx/laravel/linkto.com.com-access.log;
    error_log  /data/logs/nginx/laravel/linkto.com-error.log error;
    sendfile   off;

    client_max_body_size 8m;

    location ~ \.php$ {
        fastcgi_split_path_info ^(.+\.php)(/.+)$;
        fastcgi_pass             127.0.0.1:9000;
        fastcgi_index            index.php;
        fastcgi_param            SCRIPT_FILENAME /data/www/linkto/public/$fastcgi_script_name;
        fastcgi_intercept_errors off;
        fastcgi_buffer_size      16k;
        fastcgi_buffers          4 16k;
        fastcgi_connect_timeout  300;
        fastcgi_send_timeout     300;
        fastcgi_read_timeout     300;
        include                  fastcgi_params;
    }

    location ~ /\.ht {
        deny all;
    }

    #Redirect non-https traffic to https
    if ($scheme != "https") {
    	return 301 https://$host$request_uri;
    }

    listen 443 ssl; # managed by Certbot
    ssl_certificate /etc/letsencrypt/live/linkto.com/fullchain.pem; # managed by Certbot
    ssl_certificate_key /etc/letsencrypt/live/linkto.com/privkey.pem; # managed by Certbot
    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot

}
```

现在需要将 `linkto` 修改为 `lt`，就可以像下面这样操作：

- 打开文件

```shell
# 打开目标文件
$ vim nginx.conf
```

- 在命令模式执行替换操作即可

```shell
:%s/linkto/lt/g
```

