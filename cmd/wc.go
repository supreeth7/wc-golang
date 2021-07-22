/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// wcCmd represents the wc command
var wcCmd = &cobra.Command{
	Use:   "wc",
	Short: "A clone of the famous linux wc command",

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("filename missing")
		}

		if len(args) >= 1 {
			return nil
		}

		return fmt.Errorf("invalid command specified: %s", args[0])
	},

	Run: func(cmd *cobra.Command, args []string) {

		isBytes, _ := cmd.Flags().GetBool("bytes")
		isChars, _ := cmd.Flags().GetBool("chars")
		isLines, _ := cmd.Flags().GetBool("lines")
		isWords, _ := cmd.Flags().GetBool("words")

		if isBytes {
			readBytes(args[0])
		} else if isChars {
			readCharacters(args[0])
		} else if isLines {
			readLines(args[0])
		} else if isWords {
			readWords(args[0])
		} else {
			fmt.Println("Invalid command")
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(wcCmd)

	// Here you will define your flags and configuration settings.
	wcCmd.Flags().BoolP("bytes", "c", false, "prints the byte count")
	wcCmd.Flags().BoolP("chars", "m", false, "prints the character count")
	wcCmd.Flags().BoolP("lines", "l", false, "prints the line count")
	wcCmd.Flags().BoolP("words", "w", false, "prints the word count")
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

//This function returns the total bytes from a given file
func readBytes(fileName string) {
	file, err := os.Open(fileName)
	checkError(err)
	b := make([]byte, 100)
	bytes, err := file.Read(b)
	checkError(err)
	fmt.Printf("%d %s\n", bytes, fileName)
}

//This function returns the total characters from a given file
func readCharacters(fileName string) {
	file, err := os.Open(fileName)
	checkError(err)
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanRunes)

	characters := 0

	for scanner.Scan() {
		characters++
	}

	fmt.Printf("%d %s\n", characters, fileName)
}

//This function returns the total lines from a given file
func readLines(fileName string) {
	file, err := os.Open(fileName)
	checkError(err)
	scanner := bufio.NewScanner(file)
	lines := 0

	for scanner.Scan() {
		lines++
	}

	fmt.Printf("%d %s\n", lines, fileName)
}

//This function returns the total word count from a given file
func readWords(fileName string) {
	file, err := os.Open(fileName)
	checkError(err)
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	words := 0

	for scanner.Scan() {
		words++
	}

	fmt.Printf("%d %s\n", words, fileName)
}
