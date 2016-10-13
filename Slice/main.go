package main

import "fmt"

func main() {

	// insert to slice head

	a := []string{}
	a = append(a, "")
	copy(a[1:], a[:len(a)-1])
	a[0] = "head"
	fmt.Println(a)

	b := []string{"kkk"}
	b = append(b, "")
	copy(b[1:], b[:len(b)-1])
	b[0] = "head"
	fmt.Println(b)

}
