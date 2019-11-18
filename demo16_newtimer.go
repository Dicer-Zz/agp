package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		1 func NewTimer(d Duration) *Timer
			返回一个Timer，这个Timer会在d事件后向自己的通道中发送那个时候的时间
			用法就是计时、延迟执行
	*/
	//timer := time.NewTimer(3 * time.Second)
	//fmt.Printf("%T\n", timer)
	//fmt.Println(time.Now())
	//fmt.Println(<-timer.C)

	//让匿名函数3s后执行
	timer := time.NewTimer(3 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("函数执行结束了。。。")
	}()
	time.Sleep(5 * time.Second)
	fmt.Print("main...over...")
}
