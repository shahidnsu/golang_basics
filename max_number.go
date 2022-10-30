package main

import "fmt"

func main() {
	nums := []int{1, 2, 4, 9, 2}

	i := 0
	var max int
	for i < len(nums) {

		if max <= nums[i] {
			max = nums[i]
		}
		i = i + 1

	}
	fmt.Println(max)
}
