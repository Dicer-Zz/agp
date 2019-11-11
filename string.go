package main

import "fmt"

func main() {
	s := "世界abc"
	for k, v := range s {
		fmt.Printf("%d %x\n", k, v)
	}
}
