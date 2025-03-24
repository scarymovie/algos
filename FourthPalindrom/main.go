package fourth

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	var filtered []rune
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			filtered = append(filtered, r)
		}
	}

	p1 := 0
	p2 := len(filtered) - 1
	for p1 < p2 {
		if filtered[p1] != filtered[p2] {
			return false
		}
		p1++
		p2--
	}
	return true
}
