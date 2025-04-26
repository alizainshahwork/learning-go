package main

import "fmt"

func main() {

	fmt.Println("for loop")

	for i := 0; i < 10; i++ {

		fmt.Println(i)

	}
	for i := range 5 {

		fmt.Println(i)

	}
}
