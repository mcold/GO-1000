package main

import "fmt"

func factorial(num uint) uint {
	if num < 1 {
		panic(fmt.Sprintf("Not correct number!"))
	}
	if num == 1 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {
	fmt.Println(factorial(5))
	fmt.Println(factorial(4))
	//	fmt.Println(factorial(0))
}
