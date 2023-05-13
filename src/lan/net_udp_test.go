package lan

import (
	"fmt"
	"net"
	"testing"
	"time"
)

const (
	NetworkUdp = "udp"
	Address    = "127.0.0.1:8080"
)

func TestUDPServer(t *testing.T) {
	addr := resolveUDPAddr(t)
	conn, listenErr := net.ListenUDP(NetworkUdp, addr)
	if listenErr != nil {
		t.Fatalf("Network listen is error, and addr=%v \n", Address)
	}
	t.Log("Network server is started \n")
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			t.Log("Close server conn is failure")
		}
	}(conn)

	buf := make([]byte, 1024)
	for {
		n, rAddr, readErr := conn.ReadFromUDP(buf)
		if readErr != nil {
			t.Log(readErr)
			continue
		}
		t.Logf("Received %d bytes from %v: %s\n", n, addr, string(buf[:n]))
		time.Sleep(time.Second * 5)
		_, rErr := conn.WriteToUDP([]byte("I'm received message"), rAddr)
		if rErr != nil {
			t.Log(rErr)
			continue
		}
	}

}

func TestUDPClient(t *testing.T) {
	addr := resolveUDPAddr(t)
	conn, dErr := net.DialUDP(NetworkUdp, nil, addr)
	if dErr != nil {
		t.Fatal(dErr)
	}
	defer func(conn *net.UDPConn) {
		cErr := conn.Close()
		if cErr != nil {
			t.Log("Close client conn is failure")
		}
	}(conn)

	message := []byte("I'm send message")
	_, wErr := conn.Write(message)
	if wErr != nil {
		t.Log(err)
	}

	// 接收数据
	buf := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Received response:", string(buf[:n]))
}

func resolveUDPAddr(t *testing.T) *net.UDPAddr {
	addr, err := net.ResolveUDPAddr(NetworkUdp, Address)
	if err != nil {
		t.Fatalf("Network resolve is error, and addr=%v \n", Address)
	}
	return addr
}
