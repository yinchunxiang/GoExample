package main

import (
	"encoding/json"
	"fmt"
)

type B struct {
	BxxField string `json:"bxxField"`
}

type A struct {
	Field1 string `json:"field1"`
	B      *B
}

func main() {
	a := &A{
		Field1: "xxx",
		B:      &B{"yyyy"},
	}
	bytes, _ := json.Marshal(a)
	fmt.Println(string(bytes))
}
