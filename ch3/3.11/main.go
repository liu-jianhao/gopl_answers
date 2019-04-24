package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	start := 0
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		start = 1
	}

	index := strings.Index(s, ".")

	if index != -1 {
		transfer(s[start:index-1], &buf)
		buf.WriteString(".")
		i := index + 1
		for ; i+3 < len(s); i += 3 {
			buf.WriteString(s[i : i+3])
			buf.WriteString(",")
		}
		buf.WriteString(s[i:])
	} else {
		transfer(s[start:], &buf)
	}

	return buf.String()
}

func transfer(s string, buf *bytes.Buffer) {
	n := len(s)
	buf.WriteString(s[:n%3])
	for i := n % 3; i < n; i += 3 {
		if i != 0 {
			buf.WriteString(",")
		}
		buf.WriteString(s[i : i+3])
	}
}
