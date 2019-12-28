package main

import "fmt"

func main() {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for pos, char := range str {
		fmt.Printf("Character on position %d is %c\n", pos, char)
	}

	fmt.Println()
	str2 := "Chinese: 汉语😂"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for pos, char := range str2 {
		fmt.Printf("charcater %c starts at byte position %d\n", char, pos)
	}

	fmt.Println()
	fmt.Println("index\t int(rune)\t rune\t\t char\t bytes")
	for index, char := range str2 {
		fmt.Printf("%-5d\t %-5d\t\t %-7U\t '%c'\t %X\n", index, char, char, char, []byte(string(char)))
	}
}
