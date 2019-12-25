# Laravel项目重新部署安装依赖


## 修改composer镜像源

这里有一篇博文可以参考一下：[Composer 国内加速，修改镜像源](https://learnku.com/articles/15977/composer-accelerate-and-modify-mirror-source-in-china)

## 安装依赖

直接执行 `composer install` 进行安装依赖

![进行安装依赖](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191225_1.png)

不过不会安装成功的，证书验证失败，这是为什么呢，在网络上找了个遍，都说是安装了`openssl`扩展，具体可以参见composer官方github仓库[issue#3346](https://github.com/composer/composer/issues/3346)，可是不管用，再次运行 `composer install` 依旧是这个样子。那就索性把`composer.lock`给删除了，不它进行上体的版本锁定，于是乎再一次执行`composer install`命令进行依赖的安装

![进行安装依赖](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191225_2.gif)

成功了，那就访问一下网站页面，看一下呗！

![访问一下网站页面](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191225_2.png)

不错，正常了，那就再访问一下后台看看如何？

![访问一下后台看看如何](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191225_8.png)

不对呀！情况不对，这是错报了(当然了，正式线不应该出现这种情况，但是当前只是调试一下，回头修改一个即可，禁止错误抛出给用户，应该提高安全性！)，说是找不到 `content`，
没有重新部署之前都是好好的，现在出现这个情况了，初步怀疑是之前使用的依赖在此次安装依赖时更新了版本，所以就在搜索引擎上搜索了一下 `laravel-admin Undefined variable: content`

![laravel-admin Undefined variable: content](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191225_4.png)

果不其然，第一个就是，同样也有人遇到了这个问题，这就很好说明问题了，后台使用的是 `laravel-admin` 开发的后台管理模块，当前依赖很显然也之前的依赖代码不兼容！

![laravel-admin Undefined variable: content](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191225_5.png)

根据这个最佳回答，需要清理一下模板缓存，那就清理一下呗！清理完之后它依旧是这个样子！

![访问一下后台看看如何](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191225_8.png)

那就不行了，也不能因为升级了依赖而去修改代码，还是用之前的依赖版本，这样就不会有什么问题了，关键是它安装不成功啊，好了，***我要放大招了***，当前不是已经安装了依赖对不对！所以就分析一下当前安装依赖生成的这个(远程)`composer.lock`文件与原有(本地)`composer.lock`有什么区别：

远程服务器(linux) `composer.lock` 前50行内容如下

![远程服务器(linux) `composer.lock` 前50行内容如下](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191225_6.png)

本地工作机(mac) `composer.lock` 前50行内容如下

![本地工作机(mac) `composer.lock` 前50行内容如下](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191225_7.png)

有没有发现什么关键性的不一致，对的正如你看到一样，每个包的`dist.mirrors.url` 不一样，一个是`https://mirrors.aliyun.com/`，一个是`https://dl.laravel-china.org/`，你千万不要说每个包的 `source.reference` 和 `dist.reference` 也不一样，这个你没有办法确定，你要是改这个东西那你真是要跪了！对的，为什么我们只关注这个 `dist.mirrors.url`，原因是当你`composer install` 的时候它会去这些源去获取依赖数据！

![远程服务器(linux) `composer.lock` 前50行内容如下](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191225_9.png)
![本地工作机(mac) `composer.lock` 前50行内容如下](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191225_10.png)

就是这个样子的，将本地工作机(mac) `composer.lock` 文件中的所有 `https://dl.laravel-china.org/` 改为 `https://mirrors.aliyun.com/composer/dists/`，这样就完成OK!

![本地工作机(mac) `composer.lock` 前50行内容如下](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191225_11.png)

好的，将本地修改后的`composer.lock`文件提示到版本库，然后，在远程服务器拉取，再次执行`composer install`，再一次去验证！

![再一次去验证](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191225_13.png)

很好，关于这个composer安装依赖问题也就解决了！

## 目录
[BACK](../../../README.md)



