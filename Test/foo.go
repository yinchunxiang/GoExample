package main

import "fmt"

func inner() {
	fmt.Println("this is inner")
}

func outer() {
	inner()
}
