package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("hello, readers!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

/*
The io package specifies the io.Reader interface, which represents the read end of a stream of data

Read populates the given byte slice with data and returns the number of bytes populated and an error value.
It returns an io.EOF error when the stream ends.