package main

import "fmt"

func main() {
	fmt.Println("hey")
	var a [2]int
	a = [2]int{43, 12}
	fmt.Println(a)
	fmt.Println("size of the array is", len(a))

}
