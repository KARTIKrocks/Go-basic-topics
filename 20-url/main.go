package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456"

func main() {
	fmt.Println("welcome")
	fmt.Println(myurl)

	result,_:=url.Parse(myurl)

	// fmt.Println(result)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("the type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])
	for _, v := range qparams {
		fmt.Println("param is: ",v)
	}

	partsOfUrl:= &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=kartik",
	}

	anotherURL:= partsOfUrl.String()
	fmt.Println(anotherURL)
}
