package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go fun1()
	go fun2()

	wg.Wait()
	fmt.Println("main...over")
}

func fun1() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("fun1: %d\n", i)
	}
	wg.Done()
}

func fun2() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("fun2: %d\n", i)
	}
	wg.Done()
}
