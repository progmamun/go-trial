package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome json in golang")
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "Learncode.com", "abc123", []string{"web-dev", "js"}},
		{"MARN Bootcamp", 399, "Learncode.com", "abc123", []string{"full-stack", "js"}},
		{"Golang Bootcamp", 199, "Learncode.com", "Mam12", nil},
	}
	//package this data as JSON data

	// finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
