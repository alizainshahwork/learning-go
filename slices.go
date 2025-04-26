package main

import "fmt"

func main() {
	var s []string

	fmt.Println(s)

	var arr = make([]string, 0)

	str := [...]string{"hello", "w", "a", "a", "a", "t", "b"}

	for i := range len(str) {
		arr = append(arr, str[i])
	}

	fmt.Println(arr, len(arr))

	c := make([]string, len(arr))
	copy(c, arr)

	l := arr[5:]

	fmt.Println(l)

}
