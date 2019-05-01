package main

import (
	"fmt"
	"strings"
)

func main() {
	f := func(s string) string {
		if s == "foo" {
			return "bar"
		} else {
			return "other"
		}
	}
	s := "$foo$foo$foo$foo$foo$foo$foo"
	ret := expand(s, f)
	fmt.Println(s)
	fmt.Println(ret)
}

func expand(s string, f func(string) string) string {
	ret := strings.Replace(s, "$foo", f("foo"), 1024)
	return ret
}
