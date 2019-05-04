package main

import (
	"fmt"
	"io"
	"os"
)

type LimitedReader struct {
	reader io.Reader
	cur    int
	limit  int
}

func (lr *LimitedReader) Read(b []byte) (int, error) {
	if lr.cur >= lr.limit {
		return 0, io.EOF
	}

	if lr.cur+len(b) > lr.limit {
		b = b[:lr.limit-lr.cur]
	}
	n, err := lr.reader.Read(b)
	if err != nil {
		return n, err
	}
	lr.cur += n
	return n, err
}

func LimitReader(r io.Reader, n int) io.Reader {
	lr := LimitedReader{
		reader: r,
		limit:  n,
	}

	return &lr
}

func main() {
	file, err := os.Open("limit.txt") // 1234567890
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lr := LimitReader(file, 5)
	buf := make([]byte, 10)
	n, err := lr.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, buf) // 5 [49 50 51 52 53 0 0 0 0 0]
}
