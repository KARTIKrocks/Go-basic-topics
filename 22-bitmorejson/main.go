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

	fmt.Println("welcome")
	EncodeJson()
}

func EncodeJson() {

	opCourses := []course{
		{"ReactJS bootcamp", 299, "mywebsite.com", "abc123", []string{"web-dev", "js"}},
		{"mern bootcamp", 199, "mywebsite.com", "def123", []string{"fullstack", "js"}},
		{"angular bootcamp", 399, "mywebsite.com", "ghi123", nil},
	}

	// package this  data as json data
	// finaljson,err:= json.Marshal(opCourses)
	finaljson, err := json.MarshalIndent(opCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finaljson)
}
