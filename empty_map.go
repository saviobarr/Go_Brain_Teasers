package main

import "fmt"

func main() {
	var m map[string]int

	m = make(map[string]int)

	fmt.Println(m["errors"])

	m["errors"] = 0

	val, ok := m["errors"]

	fmt.Println(val, ok)

}
