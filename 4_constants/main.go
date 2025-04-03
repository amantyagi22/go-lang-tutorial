package main

import "fmt"

// constants and variables can be defined globally also unlike shorthand statements (name := "Aman")
const name string = "Aman"
var name2 string = "AmanTyagi"
// name3 := "Aman Tyagii ji" //this will give syntax error : non-declaration statement outside function body

func main()  {
	const (
		port = 5000
		server = "localhost"
	)
	
	fmt.Println(port, server, name, name2)
}