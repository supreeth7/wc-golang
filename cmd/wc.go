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
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// wcCmd represents the wc command
var wcCmd = &cobra.Command{
	Use:   "wc",
	Short: "A clone of the famous linux wc command",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Filename is missing")
			os.Exit(1)
		}

		isBytes, _ := cmd.Flags().GetBool("bytes")
		isChars, _ := cmd.Flags().GetBool("chars")
		if isBytes {
			readBytes(args[0])
		} else if isChars {
			readCharacters(args[0])
		} else {
			readFile(args[0])
		}

	},
}

func init() {
	rootCmd.AddCommand(wcCmd)

	// Here you will define your flags and configuration settings.
	wcCmd.Flags().BoolP("bytes", "c", false, "prints the byte counts")
	wcCmd.Flags().BoolP("chars", "m", false, "prints the character counts")

}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(fileName string) {
	lines, characters, words := 0, 0, 0

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Err ", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines++

		line := scanner.Text()
		characters += len(line)

		splitLines := strings.Split(line, " ")
		words += len(splitLines)
	}

	fmt.Printf("Filname: %v\nTotal lines: %v\nTotal words: %v\nTotal characters: %v\n", fileName, lines, words, characters)

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
