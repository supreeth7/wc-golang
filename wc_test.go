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

func TestWc(t *testing.T) {
	filename := createTempFile()
	defer os.Remove(filename)

	assertCorrectMessage := func(t testing.TB, actual, expected int) {
		t.Helper()
		if actual != expected {
			t.Errorf("Actual:%d Expected:%d", actual, expected)
		}
	}

	t.Run("counting words in a file", func(t *testing.T) {
		actual := cmd.GetWordCount(filename)
		expected := 3
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("counting lines in a file", func(t *testing.T) {
		actual := cmd.GetLineCount(filename)
		expected := 3
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("counting characters in a file", func(t *testing.T) {
		actual := cmd.GetCharacterCount(filename)
		expected := 13
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("counting bytes in a file", func(t *testing.T) {
		actual := cmd.GetByteCount(filename)
		expected := 13
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("counting the max line length in a file", func(t *testing.T) {
		actual := cmd.GetMaxLineLength(filename)
		expected := 5
		assertCorrectMessage(t, actual, expected)
	})

}
