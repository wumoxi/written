package channel_block

import "fmt"

func main() {
	ch := make(chan int)
	go pump(ch)
	fmt.Println(<-ch)
}

func pump(ch chan int) {
	for i := 9; i > 0; i-- {
		ch <- i
	}
}
