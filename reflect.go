package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p Person) Say(msg string) {
	fmt.Printf("%s say: Hello, %s\n", p.Name, msg)
}

func (p Person) PrintInfo() {
	fmt.Printf("name: %s, age: %d, sex: %s.", p.Name, p.Age, p.Sex)
}

func main() {
	p1 := Person{"dicer", 18, "male"}
	GetMessage(p1)
}

func GetMessage(input interface{}) {
	getType := reflect.TypeOf(input)
	fmt.Println("get Type is: ", getType.Name()) //get Type is:  Person
	fmt.Println("get Kind is: ", getType.Kind()) //get Kind is:  struct
	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is: ", getValue) //get all Fields is:  {dicer 18 male}
}
