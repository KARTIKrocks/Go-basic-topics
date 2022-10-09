package main

import "fmt"

func main() {
	fmt.Println("welcome")
	//  no inheritance in golang; No super or parent

	kartik := User{"Kartik","kartik@gmail.com",true,10}
	fmt.Println(kartik)
	fmt.Printf("kartik details are: %+v\n",kartik)
	fmt.Printf("Name is %v and email is %v.\n",kartik.Name,kartik.Email)
	kartik.GetStatus()
	kartik.NewMail()
	fmt.Printf("Name is %v and email is %v.\n",kartik.Name,kartik.Email)
}


type User struct {
	Name string
	Email string
	Status bool
	Age int
	
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ",u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:" , u.Email)
}