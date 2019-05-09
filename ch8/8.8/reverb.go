package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	notidleCh := make(chan interface{})
	go func() {
		ticker := time.NewTicker(time.Second)
		counter := 0
		max := 10
		for {
			select {
			case <- ticker.C:
				counter++
				if counter == max {
					ticker.Stop()
					c.Close()
					return
				}
			case <- notidleCh:
				counter = 0
			}
		}
	}()

	input := bufio.NewScanner(c)
	for input.Scan() {
		notidleCh <- struct{}{}
		go echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
