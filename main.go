package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	title      string
	desc       string
	isComplete bool
}

func format(title, desc string) (title_d, desc_d string) {
	title_d = strings.TrimSpace(title)
	desc_d = strings.TrimSpace(desc)
	return title_d, desc_d
}

func formatTodo(todo Todo) {
	fmt.Println("Title: ", todo.title)
	fmt.Println("Desc: ", todo.desc)
	fmt.Println("Completed: ", todo.isComplete)
	fmt.Println()
}

var (
	id int = 1
)

func main() {
	todos := make(map[int]Todo)

	reader := bufio.NewReader(os.Stdin)
	for {

		// clear terminal

		fmt.Print("What do you want to do?\n")
		fmt.Print("[A] Add a new ToDO \n")
		fmt.Print("[D] Delete a ToDO \n")
		fmt.Print("[L] List all ToDos \n")
		fmt.Print("[C] Complete a ToDo \n")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Println("error reading line")
		}
		task := strings.ToUpper(strings.TrimSpace(input))
		switch string(task) {
		case "A":
			fmt.Print("Add title:\n")
			title, err := reader.ReadString('\n')
			if err != nil {
				log.Println("error reading line")
			}
			fmt.Print("Add description:\n")
			desc, err := reader.ReadString('\n')
			if err != nil {
				log.Println("error reading line")
			}
			title_d, desc_d := format(title, desc)
			todos[id] = Todo{
				title:      title_d,
				desc:       desc_d,
				isComplete: false,
			}
			fmt.Print("Task added\n")
			formatTodo(todos[id])
			id++
		case "D":
			var i int = 1
			for _, todo := range todos {
				fmt.Printf("[ %v ]", i)
				formatTodo(todo)
				i++
			}

			key, err := reader.ReadString('\n')
			if err != nil {
				log.Println("error reading line")
			}
			strKey, err := strconv.Atoi(key)
			if err != nil {
				fmt.Println("cannot parse to int")
			}
			delete(todos, strKey)
			for _, todo := range todos {
				formatTodo(todo)
			}
		case "L":
			for _, todo := range todos {
				formatTodo(todo)
			}
		case "C":
			fmt.Print("You have typed C\n")
		case "/EXIT":
			return
		}
	}
}
