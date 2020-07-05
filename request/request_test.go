package request

import (
	"flag"
	"path/filepath"
	"testing"
)

var update = flag.Bool("update", false, "updated golden files")

const datadir = "testdata"

func TestExecute(t *testing.T) {
	r, err := Open(filepath.Join(datadir, "reg-get.yaml"))
	if err != nil {
		t.Errorf("failed to open, %v", err)
	}

}
