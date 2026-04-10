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

type Bird struct {
    name string
}

func (b Bird) Speak() {
    fmt.Println("啾啾")
}

type Animal interface {
    Speak()
}

func MakeSpeak(a Animal) {
    a.Speak()
}

func main() {
    dog1 := Dog{name: "旺财"}
    cat1 := Cat{name: "小猫"}
    bird := Bird{name: "miaomiao"}
    MakeSpeak(dog1)
    MakeSpeak(cat1)
    MakeSpeak(bird)
}
