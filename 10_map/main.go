package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["key1"] = 2
	m["key2"] = 3
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)
}
