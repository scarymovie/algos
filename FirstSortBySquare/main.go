package main

import "fmt"

func main() {
	nums := []int{-3, -2, 0, 1, 3, 5}
	fmt.Println(sortedSquares(nums))
}

func sortedSquares(nums []int) []int {
	p1 := 0
	p2 := len(nums) - 1
	result := make([]int, 0, len(nums))
	for p1 <= p2 {
		if abs(nums[p1]) > abs(nums[p2]) {
			result = append(result, nums[p1]*nums[p1])
			p1++
		} else {
			result = append(result, nums[p2]*nums[p2])
			p2--
		}
	}
	return reverseSlice(result)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func reverseSlice(nums []int) []int {
	p1, p2 := 0, len(nums)-1
	for p1 < p2 {
		nums[p1], nums[p2] = nums[p2], nums[p1]
		p1++
		p2--
	}
	return nums
}
