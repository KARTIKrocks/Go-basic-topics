package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome")
	content := "this needs to go in a file - kartik.com"

	file,err := os.Create("./kartik.txt")
	checkNillErr(err)

	length,err:=io.WriteString(file,content)
	checkNillErr(err)
	fmt.Println("length is: ",length)
	defer file.Close()
	readFile("./kartik.txt")
}

func readFile(filename string)  {
	databyte,err:=ioutil.ReadFile(filename)
	checkNillErr(err)
	fmt.Println("text data inside the file is \n",databyte)
	fmt.Println("text data inside the file is \n",string(databyte))
}

func checkNillErr(err error)  {
	if err!=nil {
		panic(err)
	}
}