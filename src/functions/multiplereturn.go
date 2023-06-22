package functions

import (
	"io"
	"os"
)

func R0() int {
	i := 10
	y := i + 10

	return y
}

func R1() (int, int) {
	i := 10
	y := i + 10

	return y, i
}

func R2() (x int, v int) {
	i := 10
	y := i + 10

	return i, y
}

func R3() (x int, y int) {
	i := 10
	y = i + 10

	return i, y
}

// Contents returns the file's contents as a string.
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close() // f.Close will run when we're finished.

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...) // append is discussed later.
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err // f will be closed if we return here.
		}
	}
	return string(result), nil // f will be closed if we return here.
}
