# Golang『文件读取』注意点

## 读文件

### 逐行读取

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// Open the file
	file, err := os.Open("code/golang/read_writer_data/file_read_writer/input.dat")
	if err != nil {
		fmt.Printf("An error occurred on opening the input file\n" +
			"Does the file exist?\n" +
			"Have you got access to it?\n")
		return // exit the function on error
	}
	// Close the open file
	defer file.Close()

	// Gets a new reader
	reader := bufio.NewReader(file)

	// Loop reads
	for {
		readString, err := reader.ReadString('\n')
		fmt.Printf("The input was: %s", readString)
		if err != nil && err == io.EOF {
			return
		}
	}
}
```

变量 `file` 是 `*os.File` 类型的。该类型是一个结构，表示一个打开文件的描述符（文件句柄）。然后，使用 `os` 包里的 `Open` 函数来打开一个文件。该函数的参数是文件名，类型为 `string`。在上面的程序中，我们以只读模式打开 `input.dat` 文件。

如果文件不存在或者程序没有足够的权限打开这个文件，Open函数会返回一个错误：`file, err = os.Open("input.dat")`。如果文件打开正常，我们就使用 `defer file.Close()` 语句确保在程序退出前关闭该文件。然后，我们使用 `bufio.NewReader` 来获得一个读取器变量。

通过使用 `bufio` 包提供的读取器（写入器也类似），如上面程序所示，我们可以很方便的操作相对高层的 string 对象，而避免了去操作比较底层的字节。

接着，我们在一个无限循环中使用 `ReadString('\n')` 或 `ReadBytes('\n')` 将文件的内容逐行（行结束符 '\n'）读取出来。

**注意：** 在之前的例子中，我们看到，Unix和Linux的行结束符是 \n，而Windows的行结束符是 \r\n。在使用 `ReadString` 和 `ReadBytes` 方法的时候，我们不需要关心操作系统的类型，直接使用 \n 就可以了。另外，我们也可以使用 `ReadLine()` 方法来实现相同的功能。

一旦读取到文件末尾，变量 `readerError` 的值将变成非空（事实上，其值为常量 `io.EOF`），我们就会执行 `return` 语句从而退出循环。

## 其它类似函数

### 1) 将整个文件的内容读到一个字符串里

如果您想这么做，可以使用 `io/ioutil` 包里的 `ioutil.ReadFile()` 方法，该方法第一个返回值的类型是 `[]byte`, 里面存放读取到的内容，第二个返回值是错误，如果没有错误发生，第二个返回值为 `nil`。请看下面的示例，类似的，函数 `WriteFile()` 可以将 `[]byte` 的值写入文件。

```go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "code/golang/read_writer_data/file_read_writer/products.txt"
	outputFile := "code/golang/read_writer_data/file_read_writer/products_copy.txt"
	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Read File Error: %s\n", err)
	}

	fmt.Printf("%s\n", string(bytes))
	err = ioutil.WriteFile(outputFile, bytes, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Write File Error: %s\n", err)
	}
}
```

运行程序输出如下并且生成一个名为 `products_copy.txt` 的文件！

```shell
京东价：京东价为商品的销售价，是您最终决定是否购买商品的依据。
划线价：商品展示的划横线价格为参考价，并非原价，该价格可能是品牌专柜标价、商品吊牌价或由品牌供应商提供的正品零售价（如厂商指导价、建议零售价等）或该商品在京东平台上曾经展示过的销售价；由于地区、时间的差异性和市场行情波动，品牌专柜标价、商品吊牌价等可能会与您购物时展示的不一致，该价格仅供您参考。
折扣：如无特殊说明，折扣指销售商在原价、或划线价（如品牌专柜标价、商品吊牌价、厂商指导价、厂商建议零售价）等某一价格基础上计算出的优惠比例或优惠金额；如有疑问，您可在购买前联系销售商进行咨询。
异常问题：商品促销信息以商品详情页“促销”栏中的信息为准；商品的具体售价以订单结算页价格为准；如您发现活动商品售价或促销信息有异常，建议购买前先联系销售商咨询。
```

### 2) 带缓冲的读取

在很多情况下，文件的内容是不按行划分的，或者干脆就是一个二进制文件。在这种情况下，`ReadString()` 就无法使用了，我们可以使用 `bufio.Reader` 的 `Read()` 方法，它只接收一个参数：

```go
buf := make([]byte, 1024)
...
n, err := reader.Read(buf)
if n == 0 {
	break
}
```

变量n的值表示读取到的字节数。请看下面的示例：

```go
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// Data source and destination
	inputFile := "code/golang/read_writer_data/file_read_writer/description.txt"
	outputFile := "code/golang/read_writer_data/file_read_writer/description_copy.txt"

	// Open the input file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?" +
			"Have you got access to is?\n")
		return
	}
	defer file.Close()

	// Gets a new reader
	reader := bufio.NewReader(file)

	// Initialize a new buffer
	block := make([]byte, 1)

	// Initialize a buffer of the read
	var buffer bytes.Buffer

	// Loop reads
	for {
		n, err := reader.Read(block)
		if err != nil && err != io.EOF {
			fmt.Fprintf(os.Stderr, "Read File Error: %s\n", err)
		}
		if n == 0 || err == io.EOF {
			break
		}
		buffer.WriteString(string(block))
	}

    // Write to file
	fmt.Printf("%s\n", buffer.String())
	err = ioutil.WriteFile(outputFile, []byte(buffer.String()), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Write File Error: %s\n", err)
	}
}
```

运行程序输出如下并生成一个名为 `description_copy.txt` 的文件。

```shell
京东价：京东价为商品的销售价，是您最终决定是否购买商品的依据。
划线价：商品展示的划横线价格为参考价，并非原价，该价格可能是品牌专柜标价、商品吊牌价或由品牌供应商提供的正品零售价（如厂商指导价、建议零售价等）或该商品在京东平台上曾经展示过的销售价；由于地区、时间的差异性和市场行情波动，品牌专柜标价、商品吊牌价等可能会与您购物时展示的不一致，该价格仅供您参考。
折扣：如无特殊说明，折扣指销售商在原价、或划线价（如品牌专柜标价、商品吊牌价、厂商指导价、厂商建议零售价）等某一价格基础上计算出的优惠比例或优惠金额；如有疑问，您可在购买前联系销售商进行咨询。
异常问题：商品促销信息以商品详情页“促销”栏中的信息为准；商品的具体售价以订单结算页价格为准；如您发现活动商品售价或促销信息有异常，建议购买前先联系销售商咨询。
```

### 3) 按列读取文件中的数据

如果数据是按列排列并用空格分隔的，你可以使用 `fmt` 包提供的以 FScan 开头的一系列函数来读取他们。请看以下程序，我们将 3 列的数据分别读入变量 v1、v2 和 v3 内，然后分别把他们添加到切片的尾部。


```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Open the file(打开文件)
	file, err := os.Open("code/golang/read_writer_data/file_read_writer/products2.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Open the file error: %s\n", err)
		return
	}

	// Close the open file(关闭文件)
	defer file.Close()

	// Initializes the column read result(初始化列读取结果)
	var col1, col2, col3 []string

	// Loop read(循环读取)
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}

	// Print the column read result(打印列读取结果)
	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}
```

运行程序输出结果如下：

```shell
[张三 李四 春花]
[男 男 女]
[23 18 18]
```

**注意**： `path` 包里包含一个子包叫 `filepath`，这个子包提供了跨平台的函数，用于处理文件名和路径。例如 `Base()` 函数用于获得路径中的最后一个元素（不包含后面的分隔符）：


```go
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := "code/golang/read_writer_data/file_read_writer/products2.txt"
	fmt.Printf("file path of the base: %s\n", filepath.Base(path))
}
```

运行程序输出结果如下：

```shell
file path of the base: products2.txt
```

### 练习：读取类CSV格式文件内容到结构体并输出


文件 books.txt 的内容如下：

```text
"The ABC of Go";25.5;1500
"Functional Programming with Go";56;280
"Go for It";45.9;356
"The Go Way";55;500
```

每行的第一个字段为 title，第二个字段为 price，第三个字段为 quantity。内容的格式基本与示例 [按列读取文件中的数据](./GolangReadFile.md#按列读取文件中的数据) 的相同，除了分隔符改成了分号。请读取出文件的内容，创建一个结构用于存取一行的数据，然后使用结构的切片，并把数据打印出来。

关于解析 CSV 文件，`encoding/csv` 包提供了相应的功能。具体请参考 [http://golang.org/pkg/encoding/csv/](http://golang.org/pkg/encoding/csv/)

## compress包：读取压缩文件

`compress` 包提供了读取压缩文件的功能，支持的压缩文件格式为：bzip2、flate、gzip、lzw 和 zlib。

下面的程序展示了如何读取一个 gzip 文件。

~~这个地方留坑回头回填！~~

