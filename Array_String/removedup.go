package main

import "fmt"

// Go through the array, and track index of repeated value seperatly.
// If current value is equal to next value, continue, index of repeated value
// will remain the same. When we encounter next good value, take it and put it
// in position next to the index of repeated value.
// Repeat.

func removeDuplicates(nums []int) int {
	i := 0
	for j := range nums {
		if nums[i] != nums[j] {
			fmt.Print(i, j)
			i += 1
			nums[i] = nums[j]
		}
	}
	fmt.Println(nums[:i+1])
	return i + 1
}

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Print(removeDuplicates(arr))
}
