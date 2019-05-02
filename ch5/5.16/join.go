package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(join("hello", "world"))
	fmt.Println(join("database", "network", "linux"))
}

func join(s ...string) string {
	return strings.Join(s, "")
}
