package main

import "fmt"

func main() {
	var a = []int{1, 2, 3}
	for k, v := range a {
		fmt.Printf("%d %d\n", k, v)
	}
	a = append(a, []int{4, 5, 6}...)
	for k, v := range a {
		fmt.Printf("%d %d\n", k, v)
	}
}
