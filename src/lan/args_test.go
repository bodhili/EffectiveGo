package lan

import (
	"container/list"
	"fmt"
	"os"
	"strings"
	"testing"
)

type CmdArgs struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func TestCmdArgs(t *testing.T) {
	args := list.New()
	for i := 1; i < len(os.Args); i++ {
		//for _, arg := range os.Args[1:] {
		arg := os.Args[i]
		if strings.Contains(arg, "=") {
			before, after, _ := strings.Cut(arg, "=")
			cmdArgs := CmdArgs{
				Key:   before,
				Value: after,
			}
			args.PushBack(cmdArgs)
		}
	}

	if args.Len() == 0 {
		fmt.Println("no args")
	}

	for e := args.Front(); e != nil; e = e.Next() {
		c := e.Value.(CmdArgs)
		fmt.Println(c.Key + "=" + c.Value)
	}
}
