package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	buf.WriteString(s[:n%3])
	for i := n % 3; i < n; i += 3 {
		if i != 0 {
			buf.WriteString(",")
		}
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}
