package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(200 * time.Millisecond)
	return
}
