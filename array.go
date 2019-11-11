package main

import "fmt"

func main() {
	var a = [...]int{1, 2, 3}
	var b = &a
	fmt.Println(a[0], a[1])
	fmt.Println(b[0], b[1])

	for i, v := range b {
		fmt.Println(i, v)
	}
	fmt.Printf("a: %T\n", a)
	fmt.Printf("a: %#v\n", a)
	fmt.Printf("b: %T\n", *b)
	fmt.Printf("b: %#v\n", *b)
	fmt.Printf("b: %T\n", b)
	fmt.Printf("b: %#v\n", b)
}
