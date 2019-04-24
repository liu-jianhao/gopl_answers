package main

import "fmt"

func main() {
	data := []string{"one", "one", "one", "two", "two", "three", "three"}
	data = removeDup(data)
	fmt.Printf("%q\n", data)
}

func removeDup(strings []string) []string {
	out := strings[:1]
	pre := strings[0]
	for i, s := range strings {
		if i != 0 && pre != s {
			out = append(out, s)
		}
		pre = s
	}
	return out
}
