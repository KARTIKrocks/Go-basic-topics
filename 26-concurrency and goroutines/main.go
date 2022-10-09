package main

import (
	"fmt"
	"net/http"
	"sync"
	// "time"
)

var signals = []string{"test"}

var mut sync.Mutex //pointer

var wg sync.WaitGroup //pointers
func main() {
	// go greeter("hello")
	// greeter("world")

	websitelist := []string{
		"https://google.com",
		"https://go.dev",
		"https://facebook.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Second)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("oops in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
