package main

import "fmt"

func main() {
	var arr []map[string]interface{}
	fmt.Printf("Type: %T, value: %v\n", arr, arr)
	arr2 := []map[string]interface{}{
		{
			"name": "dicer",
			"age": 21,
		},
		{
			"name": "topcoder",
			"age": 123,
		},
	}
	fmt.Printf("Type: %T, value: %v\n", arr2, arr2)
}
