package main

import "fmt"

func main() {
	arr := [4]string{"he", "ll", "ow", "world"}
	p := &arr
	pp := &p
	fmt.Printf("%T %T %T\n", arr, p, pp)
	fmt.Printf("%v %v %v\n", arr, p, pp)
	fmt.Println((*p)[0], p[0], (**pp)[0])
}
