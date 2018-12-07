package main


type MyInterface interface {
	sayHello()
}

func main() {
	var hello MyInterface = &World{}
	hello = hello.(*Hello)
	hello.sayHello()
}
