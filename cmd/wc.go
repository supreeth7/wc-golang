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

	"github.com/spf13/cobra"
)

// wcCmd represents the wc command
var wcCmd = &cobra.Command{
	Use:   "wc",
	Short: "A clone of the famous linux wc command",
	Args:  cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		isBytes, _ := cmd.Flags().GetBool("bytes")
		isChars, _ := cmd.Flags().GetBool("chars")
		isLines, _ := cmd.Flags().GetBool("lines")
		isWords, _ := cmd.Flags().GetBool("words")

		switch {
		case isBytes:
			result := ReadBytes(args[0])
			printResult(result, args[0])
		case isChars:
			result := ReadCharacters(args[0])
			printResult(result, args[0])
		case isLines:
			result := ReadLines(args[0])
			printResult(result, args[0])
		case isWords:
			result := ReadWords(args[0])
			printResult(result, args[0])
		default:
			lines := ReadLines(args[0])
			words := ReadWords(args[0])
			bytes := ReadBytes(args[0])
			chars := ReadCharacters(args[0])
			fmt.Printf("%d %d %d %d %s\n", lines, words, bytes, chars, args[0])
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

//Error handling
func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

//This function prints the result with the given arguments
func printResult(n int, file string) {
	fmt.Printf("%d %s\n", n, file)
}

//This function returns the total bytes from a given file
func ReadBytes(fileName string) int {
	file, err := os.Open(fileName)
	checkError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	bytes := 0

	for scanner.Scan() {
		bytes++
	}
	return bytes
}

//This function returns the total characters from a given file
func ReadCharacters(fileName string) int {
	file, err := os.Open(fileName)
	checkError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanRunes)

	characters := 0

	for scanner.Scan() {
		characters++
	}

	return characters
}

//This function returns the total lines from a given file
func ReadLines(fileName string) int {
	file, err := os.Open(fileName)
	checkError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := 0

	for scanner.Scan() {
		lines++
	}

	return lines
}

//This function returns the total word count from a given file
func ReadWords(fileName string) int {
	file, err := os.Open(fileName)
	checkError(err)
	scanner := bufio.NewScanner(file)
	defer file.Close()

	scanner.Split(bufio.ScanWords)

	words := 0

	for scanner.Scan() {
		words++
	}

	return words
}
