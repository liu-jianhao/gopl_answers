package main

import (
	"fmt"
	"io"
	"os"
)

type CountWriter struct {
	w io.Writer
	c int64
}

func (cw *CountWriter) Write(p []byte) (n int, err error) {
	n, err = cw.w.Write(p)
	if err != nil {
		return
	}
	cw.c += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CountWriter{
		w: w,
	}
	return &cw, &(cw.c)
}

func main() {
	cw, c := CountingWriter(os.Stdout)
	fmt.Fprintf(cw, "%s", "hello world\n")
	fmt.Println(*c)
	cw.Write([]byte("append something\n"))
	fmt.Println(*c)
}
