# Golang『启动外部命令和程序』注意点

os 包有一个 `StartProcess` 函数可以调用或启动外部系统命令和二进制可执行文件；它的第一个参数是要运行的进程，第二个参数用来传递选项或参数，第三个参数是含有系统环境基本信息的结构体。

这个函数返回被启动进程的 id(pid)，或者启动失败返回错误。

exec 包中也有同样功能的更简单的结构体和函数；主要是 `exec.Command(name string, arg ...string)` 和 `Run()`。首先需要用系统命令或可执行文件的名字创建一个 `Command` 对象，然后用这个对象作为接收者调用 `Run()`。下面的程序（因为是执行 Linux 命令，只能在 Linux 下面运行）演示了它们的使用：

下面示例展示了 `os.StartProcess` 函数的使用：

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	proAttr := &os.ProcAttr{
		Env: os.Environ(),
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}

	// 1st example: list files
	process, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, proAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err)
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", process)

	// 2nd example: show all processes
	process, err = os.StartProcess("/bin/ps", []string{"ps", "-e", "-opid,ppid,comm"}, proAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err)
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", process)
}
```

运行程序输出类似如下：

```shell
The process id is &{4807 0 0 {{0 0} 0 0 0 0}}总用量 4
-rw-r--r-- 1 root root   0 2月  13 00:47 a.txt
-rw-r--r-- 1 root root   0 2月  13 00:47 b.txt
-rw-r--r-- 1 root root   0 2月  13 00:47 c.txt
-rw-r--r-- 1 root root 658 2月  13 01:00 ListFileForGo.go
The process id is &{4808 0 0 {{0 0} 0 0 0 0}}#                                                                                                                                       ➜  external   PID  PPID COMMAND
    1     0 systemd
    2     0 kthreadd
    4     2 kworker/0:0H
    5     2 kworker/u4:0
    6     2 ksoftirqd/0
    7     2 migration/0
    8     2 rcu_bh
    9     2 rcu_sched
   10     2 lru-add-drain
   11     2 watchdog/0
   12     2 watchdog/1
   13     2 migration/1
   14     2 ksoftirqd/1
   16     2 kworker/1:0H
```

## 目录
[Back](../GolangNotice.md)