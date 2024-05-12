package main

import (
	"fmt"
	"slices"
)

func main() {
	//string creation with the size

	var a []string
	//string will be empty
	//note object are default value will be nil(null)
	a = []string{"h", "r", "u", "s", "h", "i"}
	fmt.Println("string size is going to be", len(a))
	fmt.Println("slice between one and two", slices.Delete(a, 1, 3))

}
