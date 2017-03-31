package main

import (
	"fmt"
	"testing"
)

func TestOuter(t *testing.T) {
	inner := func() {
		fmt.Println("fake inner")
	}
	outer()
}
