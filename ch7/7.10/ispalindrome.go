package main

import (
	"fmt"
	"sort"
)

type palindrome []int

func (p palindrome) Len() int {
	return len(p)
}

func (p palindrome) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p palindrome) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func ispalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len()/2; i++ {
		x, y := i, s.Len()-1-i
		if !(!s.Less(x, y) && !s.Less(y, x)) {
			return false
		}
	}
	return true
}

func main() {
	p := palindrome{1, 2, 3}
	fmt.Println(ispalindrome(p))
	p = palindrome{1, 2, 1}
	fmt.Println(ispalindrome(p))
}
