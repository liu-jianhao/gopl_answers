package main

import "fmt"

func main() {
	fmt.Println(noreturn())
}

func noreturn() (r int) {
	defer func() {
		if p := recover(); p != nil {
			r = 1
		} else {
			r = 0
		}
	}()
	panic(1)
}
