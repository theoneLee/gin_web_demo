package playerground

import (
	"fmt"
)

type Phone interface {
	Call()   // 包外可调用
	listen() // 包内调用
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) Call() {
	fmt.Println("I am Nokia, I can call you!")
}

func (n NokiaPhone) listen() {
	fmt.Println("nokia listen:", n)
}

type IPhone struct {
}

func (iPhone IPhone) Call() {
	fmt.Println("I am iPhone, I can call you!")
}
