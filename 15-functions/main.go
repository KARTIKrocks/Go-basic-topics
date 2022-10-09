package main

import "fmt"

func main() {
	fmt.Println("welcome")
	greeter()
	greeterTwo()
	
	result := adder(3,5)
	fmt.Println(result)

	result1 , result2:= proAdder(3,5,4,5,6,7,7)
	fmt.Println(result1,result2)
}

// func ()  {
// 	fmt.Println("annoynims function")
// }()

func greeterTwo()  {
	fmt.Println("another method")
}

func greeter()  {
	fmt.Println("namaste")
}

func adder(valOne int, valTwo int) int {
	return valOne+valTwo
}

func proAdder(values ...int) (int,string){
	total :=0

	for _, v := range values {
		total+=v
	}
	return total,"a function can return more than one value"
}