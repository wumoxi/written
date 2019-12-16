# HTTP服务器的性能分析

- `import _ "net/http/pprof"`
- `访问/debug/pprof`
- `使用go tool pprof分析性能`

## 构建一个简单的HTTP服务器

编写一个简单的HTTP服务器，并且引入 `net/http/pprof` 包，对于这个包的引入，也仅仅只是引入而已，在代码中并不会使用该包，只是调用了该包的 init 函数进行引入包的初始化而已！

```go
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.Write([]byte(`{"username": "zhangsan", "age": "18", "sex": "male", "mobile": "13927928289"}`))
	})
	log.Fatal(http.ListenAndServe(":8859", nil))
}
```

运行该HTTP服务器

```shell
$ go fun sampleserver.go
```

通过浏览器访问`http://localhost:8859`，可以看到该HTTP服务首页输出为一段JSON响应数据。

![HTTP服务器首页响应JSON数据](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191217_5.png)

通过浏览器访问`http://localhost:8859/debug/pprof`，可以看到HTTP服务器性能分析数据。

![HTTP服务器性能分析数据](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191217_6.png)

## 使用go tool pprof分析性能

### 使用pprof工具查看堆配置文件

在命令行输入 `go tool pprof http://localhost:8859/debug/pprof/heap` 然后输入 `web`，它会自动打开浏览器显示堆配置文件。

```shell
$ go tool pprof http://localhost:8859/debug/pprof/heap
Fetching profile over HTTP from http://localhost:8859/debug/pprof/heap
Saved profile in /Users/warnerwu/pprof/pprof.alloc_objects.alloc_space.inuse_objects.inuse_space.001.pb.gz
Type: inuse_space
Time: Dec 17, 2019 at 1:52am (CST)
No samples were found with the default sample value type.
Try "sample_index" command to analyze different sample values.
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) web
```

具体显示堆配置文件类似如下所示

![显示堆配置文件](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191217_8.png)


### 使用pprof工具查看30秒的CPU配置文件

在命令行输入 `go tool pprof http://localhost:8859/debug/pprof/profile?seconds=30` 紧接着访问`http://localhost:8859` 提供的服务，然后pprof会捕获这30秒内的CPU利用率使用情况，然后输入 `web`，它会自动打开浏览器显示30秒的CPU配置文件。

```shell
go tool pprof http://localhost:8859/debug/pprof/profile\?seconds\=30
Fetching profile over HTTP from http://localhost:8859/debug/pprof/profile?seconds=30
Saved profile in /Users/warnerwu/pprof/pprof.samples.cpu.003.pb.gz
Type: cpu
Time: Dec 17, 2019 at 2:10am (CST)
Duration: 30s, Total samples = 0 
No samples were found with the default sample value type.
Try "sample_index" command to analyze different sample values.
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) web
```
具体显示30秒的CPU配置文件类似如下所示

![30秒的CPU配置文件](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191217_9.png)