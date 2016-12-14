package main

import "fmt"

type A struct {
}

type B struct {
	Aa *A
}

type C struct {
	I interface{}
}

func main() {
	b := &B{}
	if nil == b.Aa {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}

	c := &C{}
	if nil == c.I {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}

}
