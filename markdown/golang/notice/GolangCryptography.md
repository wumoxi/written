# Golang『密码学』注意点

通过网络传输的数据必须加密，以防止被 hacker（黑客）读取或篡改，并且保证发出的数据和收到的数据检验一致。 鉴于 Go 母公司的业务，我们毫不惊讶地看到 Go 的标准库为该领域提供了超过 30 个的包：

- `hash` 包：实现了 `adler32`、`crc32`、`crc64` 和 `fnv` 校验；
- `crypto` 包：实现了其它的 hash 算法，比如 `md4`、`md5`、`sha1` 等。以及完整地实现了 `aes`、`blowfish`、`rc4`、`rsa`、`xtea` 等加密算法。

下面的示例用 `sha1` 和 `md5` 计算并输出了一些校验值。


```go
package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

func main() {
	// Gets a new hash
	hasher := sha1.New()

	// Initialize byte slice
	var b []byte

	// Hash write string. (散列写字符串)
	_, err := io.WriteString(hasher, "test")
	if err != nil {
		log.Fatalf("write string error: %s\n", err)
	}

	// Sum appends the current hash to b and returns the resulting slice. (Sum将当前散列附加到b并返回结果片)
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))

	// Hash write byte slice. (散列写字节片)
	hasher.Reset()
	data := []byte("We shall overcome!")
	_, err = hasher.Write(data)
	if err != nil {
		log.Fatalf("write byte slice error: %s\n", err)
	}

	checksum := hasher.Sum(b)
	fmt.Printf("Result: %x\n", checksum)
}
```

运行程序输出结果如下所示：

```shell
Result: a94a8fe5ccb19ba61c4c0873d391e987982fbbd3
Result: [169 74 143 229 204 177 155 166 28 76 8 115 211 145 233 135 152 47 187 211]
Result: e2222bfc59850bbb00a722e764a555603bb59b2a
```

通过调用 `sha1.New()` 创建了一个新的 `hash.Hash` 对象，用来计算 SHA1 校验值。`Hash` 类型实际上是一个接口，它实现了 io.Writer 接口：

```go
// Hash is the common interface implemented by all hash functions.
//
// Hash implementations in the standard library (e.g. hash/crc32 and
// crypto/sha256) implement the encoding.BinaryMarshaler and
// encoding.BinaryUnmarshaler interfaces. Marshaling a hash implementation
// allows its internal state to be saved and used for additional processing
// later, without having to re-write the data previously written to the hash.
// The hash state may contain portions of the input in its original form,
// which users are expected to handle for any possible security implications.
//
// Compatibility: Any future changes to hash or crypto packages will endeavor
// to maintain compatibility with state encoded using previous versions.
// That is, any released versions of the packages should be able to
// decode data written with any previously released version,
// subject to issues such as security fixes.
// See the Go compatibility document for background: https://golang.org/doc/go1compat
type Hash interface {
	// Write (via the embedded io.Writer interface) adds more data to the running hash.
	// It never returns an error.
	io.Writer

	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	Sum(b []byte) []byte

	// Reset resets the Hash to its initial state.
	Reset()

	// Size returns the number of bytes Sum will return.
	Size() int

	// BlockSize returns the hash's underlying block size.
	// The Write method must be able to accept any amount
	// of data, but it may operate more efficiently if all writes
	// are a multiple of the block size.
	BlockSize() int
}
```

通过 io.WriteString 或 hasher.Write 将给定的 []byte 附加到当前的 `hash.Hash` 对象中。

## 练习：检验 md5 算法

依据上面的 sha1 示例，使用 md5 进行检验算法。

```go
package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
)

func main() {
	// Gets a new hash
	hasher := md5.New()

	// Initialize byte slice
	var b []byte

	// Hash write string. (散列写字符串)
	_, err := io.WriteString(hasher, "test")
	if err != nil {
		log.Fatalf("write string error: %s\n", err)
	}

	// Sum appends the current hash to b and returns the resulting slice. (Sum将当前散列附加到b并返回结果片)
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))

	// Hash write byte slice. (散列写字节片)
	hasher.Reset()
	data := []byte("We shall overcome!")
	_, err = hasher.Write(data)
	if err != nil {
		log.Fatalf("write byte slice error: %s\n", err)
	}

	checksum := hasher.Sum(b)
	fmt.Printf("Result: %x\n", checksum)
}
```

运行程序输出如下：

```go
Result: 098f6bcd4621d373cade4e832627b4f6
Result: [9 143 107 205 70 33 211 115 202 222 78 131 38 39 180 246]
Result: 2f77ab6934e5a72f28c8d8511009ab5e
```

## 目录
[Back](../GolangNotice.md)    