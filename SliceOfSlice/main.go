package main

import "fmt"

func main() {
	var test [][]string
	foo := []string{"xxx", "yyy"}
	test = append(test, foo)
	fmt.Println("test size:", len(test))
	for _, v := range test[0] {
		fmt.Println(v)
	}
}
