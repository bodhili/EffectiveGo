package main

import (
	"fmt"
	"sync"
	"time"
)

func ft(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		i := i
		go func() {
			defer wg.Done()
			ft(i)
		}()
	}

	wg.Wait()
	fmt.Println("ending...")
}
