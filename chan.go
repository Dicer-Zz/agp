package main

import "fmt"

func main() {
	var ch chan bool

	ch = make(chan bool)

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("subgoroutine:", i)
		}
		ch <- true
		fmt.Println("subgoroutine over")
	}()

	data := <-ch
	fmt.Println("data:", data)
	fmt.Println("main over")
}
