package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/supreeth7/wc-golang/cmd"
)

func createTempFile() string {
	file, err := ioutil.TempFile(".", "test")
	if err != nil {
		log.Fatal(err)
	}

	text := []byte("hello\ngo\nlang")

	ioutil.WriteFile(file.Name(), text, fs.ModeAppend)
	return file.Name()
}

func TestReadWords(t *testing.T) {
	filename := createTempFile()
	defer os.Remove(filename)

	actual := cmd.ReadWords(filename)
	expected := 3

	if actual != expected {
		t.Errorf("Actual: %q \t Expected: %q", actual, expected)
	}

}
