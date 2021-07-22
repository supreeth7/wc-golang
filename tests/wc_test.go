package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"testing"

	"github.com/supreeth7/wc-golang/cmd"
)

func TestReadWords(t *testing.T) {
	file, err := ioutil.TempFile(".", "test")
	if err != nil {
		log.Fatal(err)
	}

	text := []byte("hello\ngo\nlang")

	ioutil.WriteFile(file.Name(), text, fs.ModeAppend)

	actual := cmd.ReadWords(file.Name())
	expected := 3

	if actual != expected {
		t.Error("Actual: %d \t Expected: %d", actual, expected)
	}
	if err != nil {
		t.Error("Failed to read csv data")
	}

}
