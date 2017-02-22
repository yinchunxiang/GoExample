package main

import (
	"encoding/json"
	"fmt"
)

type data struct{}

func (this *data) Error() string { return "" }

func bad() bool {
	return true
}

func test() error {
	var p *data = nil
	if bad() {
		return p
	}
	return nil
}

type Foo struct {
	Bar string `json:"bar"`
}

type HelloWorld struct {
	Ai interface{}
}

func main() {
	var e error = test()
	if e == nil {
		fmt.Println("e is nil")

	} else {
		fmt.Println("e is not nil")

	}

	hw := &HelloWorld{&Foo{Bar: "bar"}}
	hwbytes, _ := json.Marshal(hw)
	fmt.Println(string(hwbytes))

}
