package main

import (
	"fmt"
)

func main() {
	// simple switch statement
	// i := 3
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("default")
	// }

	// multiple swith conditions
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday : 
	// fmt.Println("Weekendddd");
	// default : 
	// fmt.Println("Weekday yaar")
	// }


	// type switch
	whoAmI := func (i interface{}){
		switch i.(type) {
		case int:
			fmt.Println("Its an integer")
		case string:
			fmt.Println("Its a string")
		case bool:
			fmt.Println("Its a boolean")
		default:
			fmt.Println("Other type of variable")
		}

	}
	whoAmI("golang")
}
