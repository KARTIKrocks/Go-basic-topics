package main

import "fmt"
func main() {
	fmt.Println("welcome")

	days := []string{"monday","tuesday","wednesday","thursday","friday","saturday","sunday"}
	fmt.Println(days)

	// for d := 0 ; d<len(days) ; d++{
	// 	fmt.Println(days[d])
	// }

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	// for index , day := range days{
	// 	fmt.Printf("index is %v and value is %v\n",index,day)
	// }

	// for _ , day := range days{
	// 	fmt.Printf("index is  and value is %v\n",day)
	// }

	rougueValue :=1
	for rougueValue <10{

		if rougueValue==5 {
			rougueValue++
			continue
		}
		if rougueValue==7 {
			goto lco
		}
		fmt.Println("value is: ",rougueValue)
		rougueValue++

	}
lco:
	fmt.Println("im good")
	
}