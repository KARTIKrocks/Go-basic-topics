package main

import "fmt"
func main() {
	var username string = "kartik"
	fmt.Println(username)
	fmt.Printf("variable is of type : %T\n",username)
	
	
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type : %T\n",isLoggedIn)
	
	
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type : %T\n",smallVal)
	
	
	var smallFloat float32 = 255.545454545455545454
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type : %T\n",smallFloat)
	
	
	var bigFloat float64 = 255.545454545455545454
	fmt.Println(bigFloat)
	fmt.Printf("variable is of type : %T\n",bigFloat)


	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	 //implicit type
	 var website = "imkartik"
	 fmt.Println(website)

	 //no var style
	 numberOfUser := 12345
	 fmt.Println(numberOfUser)

	 
}