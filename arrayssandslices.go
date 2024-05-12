package main

import "fmt"

func main() {
	// arrays size are predifined
	//whenever we passes to  function it takes as copy but in slices it takes as a references
	//in array you have to define the size
	//for an example
	//all values are initialised with 0
	//memory allocation  is sequantial and contagious.
	var arr = [5]int8{}
	for i := 0; i < 5; i++ {
		arr[i] = arr[i] * 2
		fmt.Println("pop array values", arr[i], &arr[i])
	}
	//output
	//   hrushikesh@Hrushikeshs-MBP GoByExample % go run  arrayssandslices.go
	//   pop array values 0 0x1400000e092
	//   pop array values 0 0x1400000e093
	//   pop array values 0 0x1400000e094
	//   pop array values 0 0x1400000e095
	//   pop array values 0 0x1400000e096

	//slices in Go
	//we can declare the slice even without initialising the size
	//these are dynamic and can be filled and popped and at any time
	//you also can resize the slice by using append function
	var slices = []int8{}
	slices = make([]int8, 5)
	for i := 0; i < 10; i++ {
		if i < 5 {
			slices[i] = int8(i)
		} else {
			slices = append(slices, int8(i))
		}
		fmt.Println("slices popping the values and memory address", slices[i], &slices[i])
	}

	//output
	// slices poping the values and memory address 0 0x14000120017
	// slices poping the values and memory address 1 0x14000120018
	// slices poping the values and memory address 2 0x14000120019
	// slices poping the values and memory address 3 0x1400012001a
	// slices poping the values and memory address 4 0x1400012001b
	// slices poping the values and memory address 5 0x14000120035
	// slices poping the values and memory address 6 0x14000120036
	// slices poping the values and memory address 7 0x14000120037
	// slices poping the values and memory address 8 0x14000120038
	// slices poping the values and memory address 9 0x14000120039

	//lets remove the element will see the memory allocation how happens in slices
	slices = append(slices[:3], slices[4:]...)
	fmt.Println(slices)
	//output
	//[0 1 2 4 5 6 7 8 9]
	for i := 0; i < 9; i++ {
		fmt.Println("lets look at the memory allocation after deallocation of 3rd index", slices[i], &slices[i])
	}
	//output
	// 	lets look at the memory allocation after deallocation of 3rd index 0 0x14000112030
	// lets look at the memory allocation after deallocation of 3rd index 1 0x14000112031
	// lets look at the memory allocation after deallocation of 3rd index 2 0x14000112032
	// lets look at the memory allocation after deallocation of 3rd index 4 0x14000112033
	// lets look at the memory allocation after deallocation of 3rd index 5 0x14000112034
	// lets look at the memory allocation after deallocation of 3rd index 6 0x14000112035
	// lets look at the memory allocation after deallocation of 3rd index 7 0x14000112036
	// lets look at the memory allocation after deallocation of 3rd index 8 0x14000112037
	// lets look at the memory allocation after deallocation of 3rd index 9 0x14000112038

	//there is a difference after removing of any element it deallocating the memory and creating the array with new memory allocation

}
