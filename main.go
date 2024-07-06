package main

import (
	"fmt"
	"sort"
)

func main() {

	var a = []int{1, 1, 1, 1, 1, 1}

	fmt.Println(jump(a))

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

func majorityElement(nums []int) int {
	var majorElement = nums[0]
	var counter = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == majorElement {
			counter++
		} else {
			counter--
		}
		if counter == 0 {
			majorElement = nums[i]
			counter = 1
		}

	}
	return majorElement

}

func rotate(nums []int, k int) {
	fmt.Println(nums)
	if k > len(nums) {
		k = k % len(nums)
	}
	var nums2 = append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	copy(nums, nums2)
	// for i := 0; i < len(nums2); i++ {
	// 	nums[i] = nums2[i]
	// }
}

func maxProfit(prices []int) int {
	var profit = 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			profit += prices[i+1] - prices[i]
		}
	}
	return profit
}

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	var buff = nums[0]
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			return true
		}
		buff = max(buff, nums[i])
		if buff == 0 {
			return false
		}
		buff--
	}
	return true
}

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var minJumps = 1
	var buffJumpCounter = nums[0]
	var maxNextJump = nums[0]

	for i := 1; i < len(nums)-1; i++ {
		buffJumpCounter--
		maxNextJump--
		maxNextJump = max(maxNextJump, nums[i])
		if buffJumpCounter == 0 {
			minJumps++
			buffJumpCounter = maxNextJump
		}

	}
	return minJumps
}
