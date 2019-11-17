package main

import (
	"fmt"
	"math"
)

func main() {
	var c Circle
	c.radius = 5
	Print(c)

	t := Triangle{3, 4, 5}
	Print(t)

	var s Shape
	s = c
	test(c)
	test(t)
	test(s)
}

func test(s Shape) {
	if instance, ok := s.(Triangle); ok {
		fmt.Printf("三角形。。。%T %p %v\n", instance, instance, instance)
	} else if instance, ok := s.(Circle); ok {
		fmt.Printf("圆形。。。%v\n", instance)
	} else {
		fmt.Printf("UNKNOW...")
	}
}
func Print(s Shape) {
	fmt.Println(s.perimeter(), s.area())
}

type Shape interface {
	perimeter() float64
	area() float64
}

type Triangle struct {
	a, b, c float64
}

type Circle struct {
	radius float64
}

func (t Triangle) perimeter() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) area() float64 {
	p := t.perimeter() / 2
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}

func (c Circle) perimeter() float64 {
	return c.radius * 2 * math.Pi
}

func (c Circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}
