package third

import "fmt"

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	numsDva := []int{2, 3, 4, 5}
	fmt.Println(maxArea(nums))
	fmt.Println(maxArea(numsDva))
}

func maxArea(height []int) int {
	p1 := 0
	p2 := len(height) - 1
	area := 0
	for p1 < p2 {
		newArea := min(height[p1], height[p2]) * (p2 - p1)
		if newArea > area {
			area = newArea
		}
		if height[p1] < height[p2] {
			p1++
		} else {
			p2--
		}
	}

	return area
}
