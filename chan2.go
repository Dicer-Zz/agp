package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var ch chan int
	ch = make(chan int)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			fmt.Println("write to channle", i)
			ch <- i
		}
		close(ch)
		fmt.Println("写结束...")
	}()
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			data, ok := <-ch
			if ok {
				fmt.Println("read ok", data)
			} else {
				fmt.Println("read bad")
			}
		}
		fmt.Println("读结束...")
	}()

	wg.Wait()
	fmt.Println("main...over")
}
