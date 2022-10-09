package main

import (
	"fmt"
	"sort"
)
func main() {
	fmt.Println("welcome")
	var fruitList = []string{"apple","tomato","peach"}
	fmt.Printf("types of fruitlist is %T \n",fruitList)
	fruitList = append(fruitList, "mango" , "banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 874
	highScore[2] = 195
	highScore[3] = 514
	// highScore[4] = 777

	fmt.Println(highScore)

	highScore = append(highScore, 777 , 888 , 999)

	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)
	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))

	// how to remove a avalue from slices based on index
	 var courses = []string{"reactjs", "javascript", "swift" , "python" , "ruby"}
	 fmt.Println(courses)

	 var index int =2
	 courses = append(courses[:index],courses[index+1:]...)
	 fmt.Println(courses)

	 
}