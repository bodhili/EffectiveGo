package names

import "net"

type Names struct {
	ID   int32
	Name string
	addr net.Addr
}

func (n *Names) Id() int32 {
	return n.ID
}

func (n *Names) SetName(name string) {
	if len(name) != 0 {
		n.Name = name
	}
}
