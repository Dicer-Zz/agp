package main

import "fmt"

func getSum() {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func main() {
	getSum()
}
