package main

import "fmt"

func main() {
	format := "%s:%d"
	// "ssssssddddd"
	line := "some:3"

	var str string
	var numb int64
	n, err := fmt.Sscanf(line, format, &str, &numb)
	fmt.Println(n, err)
	fmt.Println(str, numb)
}
