package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "hello world"

	fmt.Println(len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x \n", s[i])
	}

	fmt.Println(utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println(s[1:2]) // charAt(1)
}
