package consurrency

import (
	"bytes"
	"fmt"
	"time"
)

func Announce(message string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println(message)
	}() // Note the parentheses - must call the function.
}

var freeList = make(chan *bytes.Buffer, 100)
var serverChan = make(chan *bytes.Buffer)

// A leaky buffer
func client() {
	for {
		var b *bytes.Buffer
		// Grab a buffer if available; allocate if not.
		select {
		case b = <-freeList:
			// Got one; nothing more to do.
		default:
			// None free, so allocate a new one.
			b = new(bytes.Buffer)
		}
		load(b)         // Read next message from the net.
		serverChan <- b // Send to server.
	}
}

func load(b *bytes.Buffer) {

}

// A leaky buffer
func server() {
	for {
		b := <-serverChan // Wait for work.
		process(b)
		// Reuse buffer if there's room.
		select {
		case freeList <- b:
			// Buffer on free list; nothing more to do.
		default:
			// Free list full, just carry on.
		}
	}
}

func process(b *bytes.Buffer) {

}
