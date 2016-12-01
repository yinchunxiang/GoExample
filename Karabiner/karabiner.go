package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bitly/go-simplejson"
)

const (
	file    = "/Users/yinchunxiang/.karabiner.d/configuration/karabiner.json"
	bakfile = "/Users/yinchunxiang/.karabiner.d/configuration/karabiner.json.bak"
)

type NameMap struct {
	ProfileNameMap map[string]string
}

func NewNameMap() *NameMap {
	nm := new(NameMap)
	nm.ProfileNameMap = make(map[string]string)
	nm.ProfileNameMap["default"] = "Default profile"
	nm.ProfileNameMap["apple"] = "Apple profile"
	return nm
}

func (this *NameMap) GetProfileName(key string) string {
	return this.ProfileNameMap[key]
}

func main() {
	if 2 != len(os.Args) {
		fmt.Println("This script is used to set karabiner")
		fmt.Println("\tdefault: use <Default profile>")
		fmt.Println("\tapple:   use <Apple profile>")
		fmt.Println("Usage", os.Args[0], "apple|default")
		return
	}
	nm := NewNameMap()
	input := nm.GetProfileName(os.Args[1])
	bytes, err := ioutil.ReadFile(file)
	if nil != err {
		fmt.Println("read file error:", err.Error())
		return
	}

	///fmt.Println(string(bytes))

	json, _ := simplejson.NewJson(bytes)
	array, _ := json.Get("profiles").Array()
	for _, v := range array {
		m := v.(map[string]interface{})
		if m["name"] == input {
			m["selected"] = true
		} else {
			m["selected"] = false
		}
	}
	newbytes, _ := json.EncodePretty()
	///fmt.Println(string(newbytes))
	os.Rename(file, bakfile)
	ioutil.WriteFile(file, newbytes, 0644)

}
