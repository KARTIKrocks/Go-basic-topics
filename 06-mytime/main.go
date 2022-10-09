package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of golang ")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2002 monday"))

	createDate := time.Date(2020, time.March, 10,23,23,0,0,time.UTC)

	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Moday"))
}

// Memory Management
// new()
// Allocate memory but no INIT
// you will get a memory address
// zeroed storagge

// make()
// Allocate memory and INIT
// you will get a memory addressnon-zeroed storage

// GC(garbage collection) happens automatically
// Out of Scope or nil

// https://pkg.go.dev/runtime
