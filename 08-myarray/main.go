package main

import "fmt"

func main() {
	fmt.Println("welcome bhai")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "tomato"
	fruitList[3] = "peach"

	fmt.Println("fruitlist is: ",fruitList)
	fmt.Println("fruitlist is: ",len(fruitList))

	var vagList = [3]string{"potato" , "beans", "mushroom"}
	fmt.Println("vegy list is: ",len(vagList))
}