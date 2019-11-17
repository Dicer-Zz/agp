package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int, 10)
	var tap *[4]func(int)
	tap = new([4]func(int))
	for i := 0; i < 4; i++ {
		tap[i] = genfunc(i)
	}
	rand.Seed(time.Now().Unix())

	go waitfunc(ch)
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("所有人都接到水了..")
			return
		}
		rd := rand.Intn(4)
		tap[rd](val)
	}
}

func genfunc(id int) func(int) {
	ret := func(num int) {
		fmt.Printf("第%d个人正在第%d个水龙头接水\n", num, id+1)
	}
	return ret
}

func waitfunc(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second)
		ch <- i
		fmt.Printf("第%d个人正在等待接水\n", i)
	}
	close(ch)
}
