package pkg

import (
	"fmt"
	"testing"
)

func TestArrays(t *testing.T) {
	arr := [4][5]int{{1, 2, 3, 4, 0}, {5, 6, 7, 8, 10}, {1, 1, 1, 1, 1}, {2, 2, 2, 2, 2}}

	println(len(arr))
	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print(arr[i][j], " ")
		}

		fmt.Println()
	}
}
