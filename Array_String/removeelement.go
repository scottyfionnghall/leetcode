package main

import "fmt"

// Go through array, if current value in the array is value to be removed, lower k
// and don't increase current_index (just contine). After that, when we encounter
// good value we have current_index which at this time will be storing index of
// removed value, using it we replace value in array with next good value.
func removeElement(nums []int, val int) int {
	arr_len := len(nums)
	k := arr_len
	current_index := 0
	for i := 0; i < arr_len; i++ {
		if nums[i] == val {
			k--
			continue
		}
		nums[current_index] = nums[i]
		current_index++
	}
	return k
}

func main() {
	arr1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Print(removeElement(arr1, val))
}
