package main

import (
	"fmt"
	"os"
	"syscall"
)

func Create(filename string) {
	for try := 0; try < 2; try++ {
		file, err := os.Create(filename)
		if err == nil {
			fmt.Println(file)
			break
		}

		if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
			deleteTempFiles()
			continue
		}

		return
	}
}

func deleteTempFiles() {}

//
//func main() {
//	Create("./test.txt")
//}
