package main

import (
	"fmt"
	"io"
	"os"
)

func contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			// do nothing
			fmt.Println(err)
		}
	}(f)

	var result []byte
	buf := make([]byte, 100)

	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
	}

	return string(result), nil
}

func main() {
	s, err := contents("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(s)
}
