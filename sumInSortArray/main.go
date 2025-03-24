package second

import "fmt"

func main() {
	nums := []int{-2, 1, 6, 9, 12}
	nums2 := []int{3, 3, 12}

	target := 18
	target2 := 6
	fmt.Println(twoSum(nums, target))
	fmt.Println(twoSum(nums2, target2))
}

func twoSum(nums []int, target int) []int {
	p1 := 0
	p2 := len(nums) - 1
	result := make([]int, 0, len(nums))
	for p1 <= p2 {
		if nums[p2]+nums[p1] == target {
			result = append(result, p1+1, p2+1)
			break
		}
		if nums[p2]+nums[p1] > target {
			p2--
			continue
		}
		if nums[p2]+nums[p1] < target {
			p1++
			continue
		}

	}
	return result
}
