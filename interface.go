package main

import "fmt"

func main() {
	var m Mouse
	m.name = "logie"
	m.start()
	m.end()

	var f FlashDisk
	f.name = "free"
	f.start()
	f.end()

	var usb USB
	usb = m
	Print(usb)
	Print(m)
	Print(f)
}

type USB interface {
	start()
	end()
}

func Print(usb USB) {
	usb.start()
	usb.end()
}

type Mouse struct {
	name string
}

type FlashDisk struct {
	name string
}

func (m Mouse) start() {
	fmt.Println(m.name, "鼠标开始工作了。。。")
}

func (m Mouse) end() {
	fmt.Println("鼠标没电了。。。")
}

func (f FlashDisk) start() {
	fmt.Println(f.name, "u盘插进去了。。。")
}

func (f FlashDisk) end() {
	fmt.Println("u盘拔出去了。。。")
}
