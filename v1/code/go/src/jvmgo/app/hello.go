package main

import "fmt"

//Hello should be ...
type Hello struct {
	greet    string
	sayHello func(greet string)
}

func sayHi(greet string) {
	fmt.Println(greet)
}

func main() {
	hello := &Hello{
		greet:    "hi",
		sayHello: sayHi,
	}
	hello.sayHello("hello world")

	var a interface{} = 3.1415926
	fmt.Println(a.(float64))

	b := make(map[string]interface{})
	b["hello"] = "world"

	val,ok := b["hello"]
	fmt.Println(val)
	fmt.Println(ok)

	fmt.Println(int(a.(float64)))

	c := "1"
	d := "2"
	swapVal(&c, &d)
	fmt.Println(c)
	fmt.Println(d)
}

func swapVal(a *string, b *string){
	 temp := a
	 a = b
	 b = temp
}
