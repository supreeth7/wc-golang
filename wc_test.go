package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/supreeth7/wcg/cmd"
)

func createTempFile() *os.File {
	file, err := ioutil.TempFile(".", "test")

	if err != nil {
		log.Fatal(err)
	}

	text := []byte("hello\ngo\nlang")

	ioutil.WriteFile(file.Name(), text, fs.ModeAppend)

	return file
}

func TestWc(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, actual, expected int) {
		t.Helper()
		if actual != expected {
			t.Errorf("Actual:%d Expected:%d", actual, expected)
		}
	}

	checkError := func(err error) {
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	file := createTempFile()
	defer os.Remove(file.Name())
	data, err := cmd.ConvertFileToString(file.Name())

	checkError(err)

	t.Run("counting words in a file", func(t *testing.T) {
		actual := cmd.GetWordCount(data)
		expected := 3
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("counting lines in a file", func(t *testing.T) {
		actual := cmd.GetLineCount(data)
		expected := 3
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("counting characters in a file", func(t *testing.T) {
		actual := cmd.GetCharacterCount(data)
		expected := 13
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("counting bytes in a file", func(t *testing.T) {
		actual := cmd.GetByteCount(data)
		expected := 13
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("counting the max line length in a file", func(t *testing.T) {
		actual := cmd.GetMaxLineLength(data)
		expected := 5
		assertCorrectMessage(t, actual, expected)
	})
}
