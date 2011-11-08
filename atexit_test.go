package atexit

import (
	"exec"
	"testing"
	"io/ioutil"
	"time"
	"os"
)


func TestRegister(t *testing.T) {
	current := len(handlers)
	Register(func() {})
	if len(handlers) != current + 1 {
		t.Fatalf("can't add handler")
	}
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func TestHandler(t *testing.T) {
	filename := "/tmp/testprog.out"
	arg := time.UTC().String()

	os.Remove(filename)
	if exists(filename) {
		t.Fatalf("can't delete %s", filename)
	}

	if err := exec.Command("6g", "testprog.go").Run(); err != nil {
		t.Fatalf("can't compile")
	}

	if err := exec.Command("6l", "testprog.6").Run(); err != nil {
		t.Fatalf("can't link\n")
	}

	err := exec.Command("./6.out", filename, arg).Run()
	if err.String() != "exit status 1" {
		t.Fatalf("bad exit status (%s), should be 1", err)
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("can't read output file %s", filename)
	}

	if string(data) != arg {
		t.Fatalf("bad data")
	}
}
