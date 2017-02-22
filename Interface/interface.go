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

type Void interface{}

type Response struct {
	*Void
}

type Foo struct {
	Bar string `json:"bar"`
}

func main() {
	var e error = test()
	if e == nil {
		fmt.Println("e is nil")

	} else {
		fmt.Println("e is not nil")

	}

	resp := &Response{&Foo{Bar: "xxx"}}
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))

}
