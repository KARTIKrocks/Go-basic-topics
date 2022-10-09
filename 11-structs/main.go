package main

import "fmt"

func main() {
	fmt.Println("welcome")
	//  no inheritance in golang; No super or parent

	kartik := User{"Kartik","kartik@gmail.com",true,10}
	fmt.Println(kartik)
	fmt.Printf("kartik details are: %+v\n",kartik)
	fmt.Printf("Name is %v and email is %v.",kartik.Name,kartik.Email)
}


type User struct {
	Name string
	Email string
	Status bool
	Age int
}