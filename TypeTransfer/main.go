package main

import "fmt"

type DeviceQaRequest struct {
	AgentId   string   `json:"agentId"`
	ClientId  string   `json:"clientId"`
	Questions []string `json:"questions"`
	Answers   []string `json:"answers"`
}

type DeviceQa struct {
	Id        int64
	AgentId   int64    `json:"agentid"`
	ClientId  string   `json:"clientid"`
	Questions []string `json:"questions"`
	Answers   []string `json:"answers"`
}

func main() {
	fmt.Println("vim-go")
	d := DeviceQaRequest{
		AgentId:   "agentid",
		ClientId:  "clientid",
		Questions: []string{"q1", "qa"},
		Answers:   []string{"a1", "a2"},
	}

	fmt.Println(DeviceQa(d).AgentId)

}
