package main

import (
	"fmt"
	"sort"
)

func main() {

	var a = []int{0, 0, 1, 1, 1, 1, 2, 3, 3, 3}
	removeDuplicates2(a)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = nums1[:m]
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	fmt.Println(nums1)
}

func removeElement(nums []int, val int) int {
	var counter = 0
	var k = 0
	for counter < len(nums) {
		if nums[counter] == val {
			nums[counter] = nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			k++
		} else {
			counter++

		}
	}
	fmt.Print(nums)
	return len(nums)
}

func removeDuplicates(nums []int) int {
	var counter = 1
	if len(nums) < 2 {
		return len(nums)
	}
	for counter < len(nums) {
		if nums[counter] == nums[counter-1] {
			nums = append(nums[:counter-1], nums[counter:]...)
		} else {
			counter++
		}
	}
	fmt.Println(nums)
	return len(nums)
}

func removeDuplicates2(nums []int) int {
	var counter = 1
	var duplicateCounter = 1
	if len(nums) < 2 {
		return len(nums)
	}
	for counter < len(nums) {
		if nums[counter] == nums[counter-1] {
			duplicateCounter++
			if duplicateCounter > 2 {
				nums = append(nums[:counter-1], nums[counter:]...)

			} else {
				counter++
			}
		} else {
			counter++
			duplicateCounter = 1

		}
	}
	fmt.Println(nums)
	return len(nums)
}
