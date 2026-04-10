package main

import (
	"fmt"
)

type Stacker interface {
	Push(val int)
	Pop() (int, bool)
	Peek() (int, bool)
}

type Stack struct {
	items []int
}

func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}
func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val, true
}
func (s *Stack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false

	}
	return s.items[len(s.items)-1], true
}
// 值接收者版本的 Push —— 修改的是副本
func (s Stack) PushByValue(val int) {
	s.items = append(s.items, val)
}

// 接收接口，不关心具体是什么 Stack
func PrintTop(s Stacker) {
	val, ok := s.Peek()
	if ok {
		fmt.Println("栈顶是:", val)
	} else {
		fmt.Println("栈是空的")
	}
}

func main() {
	myStack := Stack{}
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)

	val, _ := myStack.Peek()
	fmt.Println("Peek看到的栈顶:", val)

	val, _ = myStack.Pop()
	fmt.Println("Pop弹出的值:", val)

	val, _ = myStack.Peek()
	fmt.Println("Pop后栈顶变成:", val)

	fmt.Println("剩余items:", myStack.items)

	// 用接口调用
	PrintTop(&myStack)

	// 验证值接收者
	fmt.Println("\n--- 验证值接收者 ---")
	myStack.PushByValue(999)
	fmt.Println("PushByValue(999) 之后 items:", myStack.items) // 999 不会出现

}
