package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: %s s1 s2\n", os.Args[0])
		return
	}

	if judge(os.Args[1], os.Args[2]) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func judge(s1 string, s2 string) bool {
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		m1[s1[i]]++
	}
	for i := 0; i < len(s2); i++ {
		m2[s2[i]]++
	}

	if len(m1) != len(m2) {
		return false
	} else {
		for k, v := range m1 {
			if v != m2[k] {
				return false
			}
		}
	}
	return true
}
