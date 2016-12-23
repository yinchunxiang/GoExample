package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("vim-go")
	t := time.Now()
	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Day())
	fmt.Println(t.Hour())
}
