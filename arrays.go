package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println(a)

	for i := range len(a) {
		a[i] = i
	}

	fmt.Println(a)

	newarr := [5]int{9, 0, 1, 2, 3}

	fmt.Println(newarr)

	newarrautolen := [...]int{1, 2, 3}

	fmt.Println(len(newarrautolen))

	newarr1 := [...]int{1, 6: 4, 2, 4, 5}
	fmt.Println(newarr1)

	var arr2d [3][3]int

	arr2d = [3][3]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}

	fmt.Println(arr2d)
}
