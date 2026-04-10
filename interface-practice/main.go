package main

import "fmt"

type Dog struct {
	name string
}

func (d Dog) Speak() {
	fmt.Println("Wuff")
}

type Cat struct {
	name string
}

func (c Cat) Speak() {
	fmt.Println("miao")
}
func MakeSpeak(a Animal) {
	a.Speak()
}

type Animal interface {
	Speak()
}

func main() {
	dog1 := Dog{name: "旺财"}
	cat1 := Cat{name: "小猫"}
	dog1.Speak()
	cat1.Speak()

}
