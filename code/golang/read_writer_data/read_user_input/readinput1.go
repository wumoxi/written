package main

import "fmt"

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func main() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	// 输出结果：Hi 武 沫汐
	fmt.Printf("Hi %s %s!\n", firstName, lastName)

	fmt.Sscanf(input, format, &f, &i, &s)
	// 输出结果：From the string we read: 56.12 5212 Go
	fmt.Println("From the string we read: ", f, i, s)
}
