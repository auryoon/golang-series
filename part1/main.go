package main

import (
	"fmt"
)

func main() {
	/*
		fmt itu package untuk ngeprint
		fmt.Print("Hello world!")
	*/
	var input string
	fmt.Print("Hallo dunia isekai!\n")
	fmt.Println("Hallo dunia non isekai!")
	fmt.Scan(&input)
	fmt.Print(input)
}
