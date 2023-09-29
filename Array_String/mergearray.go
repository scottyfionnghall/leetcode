package main

import (
	"fmt"
)

// Go through array 1 and compare it's last NON 0 value to the
// last value of second array. If value of first array is bigger than
// the value of second array, then take  last value of seocnd array and put it
// in the last position of first array and lower n, so we can go to the next
// value of second array starting from the end.
// When the last value of first array is bigger than the value of second array,
// then take value of first array and put it in the free (0 value) position in
//  itself and lower m, so we can go inspect next values.
// Repeat.

func merge(nums1 []int, m int, nums2 []int, n int) {
	for n != 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
	fmt.Println(nums1)
}

func main() {
	arr1 := []int{1, 2, 3, 0, 0, 0}
	arr2 := []int{2, 5, 6}
	m := 3
	n := 3
	merge(arr1, m, arr2, n)
}
