package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		func After(d Duration) <-chan Time
			返回一个会发送d时间后时间的通道
	*/
	ch := time.After(3 * time.Second)

	fmt.Println(time.Now())
	fmt.Print(<-ch)
}
