package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const input = "1234 5678 1234567890 1234567901234567890"
	scanner := bufio.NewScanner(strings.NewReader(input))

	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}

	scanner.Split(split)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}
