package main

import "fmt"

type Hello struct {
	greet string
}

type World struct {
	hello Hello
	greet string
}

func (self *Hello) sayHello() {
	fmt.Println("hello")
}

func (self *World) sayHello() {
	fmt.Println("world")
}
