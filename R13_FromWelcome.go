package main

import (
	//	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(bs []byte) (int, error) {

	num_bytes_read_from_r, err := r13.r.Read(bs)

	for i := 0; i < num_bytes_read_from_r; i++ {

		var sb byte
		if bs[i] > 'A' && bs[i] < 'Z' {
			sb = 'A'
		} else if bs[i] > 'a' && bs[i] < 'z' {
			sb = 'a'
		}

		var db byte
		if sb == 0 {
			db = bs[i]
		} else {
			db = sb + (bs[i]-sb+13)%26
		}
		bs[i] = db
	}

	return num_bytes_read_from_r, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
