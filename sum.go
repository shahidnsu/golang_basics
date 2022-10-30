package main

import "fmt"

func main() {
	// slices containg interger number
	nums := []int{1, 2, 5, 7}

	var sum int = 0
	i := 1
	for i < len(nums) {
		sum = sum + nums[i]
		i = i + 1
	}
	fmt.Println("sum:", sum)
}
