package methods

import "fmt"

type ByteSlice []byte

func (p *ByteSlice) Write(data []byte) (n int, err error) {
	// Again as above.
	*p = data
	var b ByteSlice
	_, _ = fmt.Fprintf(&b, "This hour has %d days\n", 7)
	return len(data), nil
}
