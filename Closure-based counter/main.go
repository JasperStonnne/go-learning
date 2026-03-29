package main

import "fmt"

func Counter() func() int {
	count := 0
	return func() int {
		count++
		return count

	}
}
func main() {
	fmt.Println("====================================")
	c := Counter()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

}
