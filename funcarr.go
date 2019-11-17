package main

import "fmt"

func main() {
	var funcarr []func()
	funcarr = append(funcarr, func() {
		fmt.Println("This is a func")
	})
	fmt.Println(funcarr)
	funcarr[0]()
}
