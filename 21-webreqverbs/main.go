package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("welcome")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "https://www.google.co.in"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}	
	defer response.Body.Close()
	fmt.Println("status code: ",response.StatusCode)
	fmt.Println("content length: ",response.ContentLength)

	var responseString strings.Builder
	content,_:=ioutil.ReadAll(response.Body)
	byteCount, _:=responseString.Write(content)

	fmt.Println("bytecount is: ",byteCount)
	fmt.Println(responseString.String())
	fmt.Println(string(content))
}
