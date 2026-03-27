package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func listTodos(todos []string) {

	for i, v := range todos {
		fmt.Printf("%d. %s\n", i+1, v)
	}
	fmt.Println("--------------------")
}

func main() {
	fmt.Println("请输入待办事项：")
	todos := []string{}
	data, _ := os.ReadFile("todos.json")
	json.Unmarshal(data, &todos)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.Split(input, " ")

		if parts[0] == "quit" {
			data, _ := json.Marshal(todos)
			os.WriteFile("todos.json", data, 0644)
			break
		}
		if parts[0] == "delete" {
			n, _ := strconv.Atoi(parts[1])
			todos = append(todos[:n-1], todos[n:]...)
		}
		if parts[0] == "add" {
			todos = append(todos, parts[1])
		}
		if parts[0] == "list" {
			listTodos(todos)
		}

	}

}
