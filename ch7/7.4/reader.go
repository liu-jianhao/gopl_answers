package main

import (
	"fmt"
	"io"
)

type StringReader struct {
	data string
	n    int
}

func (sr *StringReader) Read(b []byte) (int, error) {
	data := []byte(sr.data)
	if sr.n >= len(data) {
		return 0, io.EOF
	}

	data = data[sr.n:]
	n := 0
	n = copy(b, data)
	sr.n = sr.n + n
	return n, nil
}

func NewReader(in string) *StringReader {
	sr := new(StringReader)
	sr.data = in
	return sr
}

func main() {
	str := "Hello World!"
	sr := NewReader(str)
	data := make([]byte, 5)
	n, err := sr.Read(data)
	for err == nil {
		fmt.Println(n, string(data))
		n, err = sr.Read(data)
	}
}
