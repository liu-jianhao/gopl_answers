package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: %s str num\n")
		return
	}
	b := []byte(os.Args[1])
	k, _ := strconv.Atoi(os.Args[2])
	b = rotate(b, k)
	fmt.Printf("%s\n", b)
}

func rotate(s []byte, k int) []byte {
	s1 := s[:k]
	s2 := s[k:]
	s2 = append(s2, s1...)
	return s2
}
