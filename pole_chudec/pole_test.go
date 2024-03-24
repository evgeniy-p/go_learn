package main

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

func Logic(test_input string) (tempfile *os.File, result error) {
	content := []byte(test_input)
	tmpfile, err := ioutil.TempFile("", "mytempfile")
	if err != nil {
		return nil, errors.New("can't open file")
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		return nil, errors.New("can't write to file")
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		return nil, errors.New("can't read file")
	}
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin
	return tmpfile, nil
}

func TestGood(t *testing.T) {
	var err error
	os.Stdin, err = Logic("h")
	if err != nil {
		t.Errorf("%s", err)
	}
	got := get_input_char()
	want := "h"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestBad(t *testing.T) {
	var err error
	sout := os.Stdout
	os.Stdin, err = Logic("ad")
	if err != nil {
		t.Errorf("%s", err)
	}
	os.Stdout, _ = os.Open(os.DevNull)
	got := get_input_char()
	os.Stdout = sout
	want := "You have found a bug! Try again \n"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestEND(t *testing.T) {
	var err error
	sout := os.Stdout
	os.Stdin, err = Logic("END")
	if err != nil {
		t.Errorf("%s", err)
	}

	if os.Getenv("BE_CRASHER") == "1" {
		os.Stdout, _ = os.Open(os.DevNull)
		get_input_char()
		os.Stdout = sout
		return
	}
	cmd := exec.Command(os.Args[0], "-logic.run=TestEND")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	newerr := cmd.Run()
	if e, ok := newerr.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}
