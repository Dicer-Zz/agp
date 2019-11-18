package main

import (
	"fmt"
)

func main() {
	/*
		select语句与switch语句类似，
		在仅有一个通信可用的时候会选择该case，如果有多个则会公平的选择一个
		如果没有case可用：
			如果有dafault语句则执行
			否则死锁
	*/
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		//time.Sleep(time.Second)
		ch1 <- 100
	}()
	go func() {
		//time.Sleep(time.Second)
		ch2 <- 200
	}()
	select {
	case <-ch1:
		fmt.Println("case1可用")
	case <-ch2:
		fmt.Println("case2可用")
	default:
		fmt.Println("执行default...")
	}
	fmt.Println("main...over...")
}
