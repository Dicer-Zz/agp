/*
通过通信机制模拟排队接水

假设所有人只能排队等待，且只有一条队伍
接水顺序按照排队顺序
同一个水龙头只能让一个人接水
一个人只能使用一个水龙头

实现思路：
	1. 为排队的人设置一个channel
	2. 通过goroutine并发多个水龙头
	3. 通过channel和goroutine竞争解决排队

一个有趣的现实：
	本来以为水龙头的数量可能对程序的效率有很大的影响，但结果并没有。
	水龙头数量从10到1000在1000个人的情况下都是6.2s，
	根据概率知道其中大约5s是在睡眠，
	取消睡眠后大概只有20ms的执行时间，可以说明计算随机数和执行睡眠使用了大概1s。

	结果很容易理解，因为Golang的并发机制并不是真并行，而只是通过时间片轮转造成多个协程轮转执行，
	所以时间并不会快，更可能的结果是变慢，因为需要切换协程。
	即使睡眠也不会多协程同时睡眠。

	但是还有一个问题没有解决，在只有一个水龙头的时候。
	执行时间意外的长，达到了7.8s。
	在将水龙头数量设置为2时，时间就变长了惊人的6.3s，然后增加水龙头的数量，对程序效率就没有什么影响了。
	当去掉睡眠之后就没有这个问题了，大约都是20ms。

	Amaze.
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	tapNumber       int = 2
	waitTaperNumber int = 1000
)

var wgp sync.WaitGroup

func main() {
	rand.Seed(time.Now().Unix())
	start := time.Now()
	defer func() {
		fmt.Printf("cost: %v", time.Since(start))
	}()
	taps := [tapNumber]func(int){}
	for i := 0; i < tapNumber; i++ {
		taps[i] = genTap(i + 1)
	}
	waitTaper := make(chan int)

	wgp.Add(waitTaperNumber)
	go waiting(waitTaper)

	//tapNumber个水龙头并发
	for i := 0; i < tapNumber; i++ {
		go takeWater(waitTaper, taps[i])
	}

	wgp.Wait()
	fmt.Println("main...over...")
}

//利用闭包返回一个 输出信息的水龙头 函数
func genTap(id int) func(int) {
	return func(num int) {
		sleepier := rand.Int63n(10) + 1 // [1, 10]
		fmt.Printf("第%d个人正在%d号水龙头接水\n", num, id)
		time.Sleep(time.Duration(sleepier) * time.Millisecond)
		fmt.Printf("在%d毫秒之后，第%d个人在%d号水龙头接完了水\n", sleepier, num, id)
	}
}

//按时间来到水房
func waiting(ch chan<- int) {
	for i := 1; i <= waitTaperNumber; i++ {
		sleepier := rand.Int63n(10) + 1 // [1, 10]
		time.Sleep(time.Duration(sleepier) * time.Millisecond)
		fmt.Printf("第%d个人在%v来到了水房等待接水...\n", i, time.Now())
		ch <- i
	}
	close(ch)
}

func takeWater(ch <-chan int, f func(int)) {
	for {
		num, ok := <-ch
		if !ok {
			break
		}
		//此处defer不能上移，避免提前结束其他goroutine
		//如果上移，会多增加一次并不应该做的wg.Done()
		//可能造成最后几个人还没有打完水就被强制停止
		//panic: sync: negative WaitGroup counter
		defer wg.Done()
		f(num)
	}
}
