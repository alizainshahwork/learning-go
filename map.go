package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["a"] = 1

	var test = "hello"

	m[test] = 89

	val, prs := m["a"]

	fmt.Println(val, prs)
	fmt.Println(m)

}
