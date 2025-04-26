package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("test conditions")

	var num int = 10

	if num%2 == 0 {
		fmt.Println("even")
	}

	var num2 int = 0

	if num2 == 0 {
		fmt.Println("zero")
	} else if num2 < 0 {
		fmt.Println("negative")
	} else {
		fmt.Println("positive")
	}

	var option = 1
	switch option {
	case 1:
		fmt.Println("option1")

	case 2:
		fmt.Println("option 2")

	default:
		fmt.Println("end")
	}

	switch time.Now().Weekday() {
	case time.Wednesday:
		fmt.Println("wed")
	}

	vartype := func(i interface{}) {

		switch i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")

		}

	}

	vartype("test")
	vartype(1)
	vartype(true)
}
