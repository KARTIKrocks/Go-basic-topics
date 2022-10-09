package main

import "fmt"
func main() {
	fmt.Println("welcome")

	language := make(map[string]string)

	language["JS"] =  "javascript"
	language["RB"] =  "ruby"
	language["PY"] =  "python"
	fmt.Println("list of all languages: ",language)
	fmt.Println("JS shorts for : ",language["JS"])

	delete(language,"RB")
	fmt.Println("list of all language: ",language)

	// loops are interesting in golang
	for key, value := range language {
		fmt.Printf("for key %v , value is %v\n",key,value)
	}
}