package main

import (
	"time"
	"fmt"
)

func thrower (c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Println("Threw >>", i)
	}
}

func catcher(c chan int) {
	for i := 0; i < 5; i++ {
		num := <-c
		fmt.Println("Caught <<",num)
	}
}

// CaughtのほうがThrewより先になっているところは、チャネルへ入れたり取り出したりした後の出力文
// をランタイムがどのようにスケジュールするかということなので重要ではない。
func main() {
	c := make(chan int,3)
	go thrower(c)
	go catcher(c)
	time.Sleep(100 * time.Millisecond)
}