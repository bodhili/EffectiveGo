package pkg

import (
	"io/fs"
	"os"
	"testing"
)

func TestFS(t *testing.T) {
	root := "/usr/local/go/bin"
	fileSystem := os.DirFS(root)

	err := fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			t.Error(err)
		} else {
			t.Log(path)
		}
		return nil
	})

	if err != nil {
		t.Error(err)
		return
	}
}
