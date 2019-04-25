package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	elements := make(map[string]int)
	visit(elements, doc)
	for elem, counts := range elements {
		fmt.Printf("%s %d\n", elem, counts)
	}
}

func visit(e map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		e[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(e, c)
	}
}
