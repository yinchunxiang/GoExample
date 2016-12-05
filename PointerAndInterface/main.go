package main

import (
	"fmt"
)

type IFace interface {
	SetSomeField(newValue string)
	GetSomeField() string
}

type Implementation struct {
	someField string
}

func (i *Implementation) GetSomeField() string {
	return i.someField
}

func (i *Implementation) SetSomeField(newValue string) {
	i.someField = newValue
}

func Create() IFace {
	return &Implementation{someField: "Hello"}
}

func main() {
	///	var a IFace
	a := Create()
	a.SetSomeField("World")
	fmt.Println(a.GetSomeField())
}
