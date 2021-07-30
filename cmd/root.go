/*
Copyright © 2021 SUPREETH BASABATTINI <sbasabat@redhat.com>

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
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/spf13/cobra"
)

const helpText = `
 Usage: wcg [OPTION]... [FILE]...
 Print newline, word, and byte counts for each FILE, and a total line if
 more than one FILE is specified.  A word is a non-zero-length sequence of
 characters delimited by white space.

 With no FILE, or when FILE is -, read standard input.

 The options below may be used to select which counts are printed, always in
 the following order: newline, word, character, byte, maximum line length.
  -c, --bytes        	print the byte counts
  -m, --chars        	print the character counts
  -l, --lines        	print the newline counts
  	--files0-from=F	read input from the files specified by
                       	NULL-terminated names in file F;
                       	If F is - then read names from standard input
  -L, --max-line-length  print the maximum display width
  -w, --words        	print the word counts
  	--help 	display this help and exit
  	--version  output version information and exit

 GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
 Full documentation at: <https://www.gnu.org/software/coreutils/wc>
 or available locally via: info '(coreutils) wc invocation'
 `

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "wcg",
	Short:   "A clone of the famous linux wc command",
	Long:    "Prints newline, word, and byte counts for each FILE, and a total line if more than one FILE is specified",
	Version: "1.0.0",

	RunE: func(cmd *cobra.Command, args []string) error {

		isBytes, _ := cmd.Flags().GetBool("bytes")
		isChars, _ := cmd.Flags().GetBool("chars")
		isLines, _ := cmd.Flags().GetBool("lines")
		isWords, _ := cmd.Flags().GetBool("words")
		isMaxLength, _ := cmd.Flags().GetBool("max-line-length")

		var (
			file     *os.File
			fileName string
			data     string
		)

		if args[0] == "-" {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter text: ")
			text, err := reader.ReadString('\n')

			checkError(err)

			if text == "" {
				return errors.New("empty text")
			}

			data = text
			fileName = ""
		} else {
			file = openFile(args[0])
			defer file.Close()
			data = ConvertFileToString(file)
			fileName = file.Name()
		}

		switch {
		case isBytes:
			result := GetByteCount(data)
			printResult(result, fileName)
		case isChars:
			result := GetCharacterCount(data)
			printResult(result, fileName)
		case isLines:
			result := GetLineCount(data)
			printResult(result, fileName)
		case isWords:
			result := GetWordCount(data)
			printResult(result, fileName)
		case isMaxLength:
			result := GetMaxLineLength(data)
			printResult(result, fileName)
		default:
			lines := GetLineCount(data)
			words := GetWordCount(data)
			bytes := GetByteCount(data)
			chars := GetCharacterCount(data)
			maxLinLength := GetMaxLineLength(data)
			fmt.Printf("%d %d %d %d %d%s\n", lines, words, bytes, chars, maxLinLength, fileName)
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.SetHelpTemplate(helpText)

	// flags
	rootCmd.Flags().BoolP("bytes", "c", false, "prints the byte count")
	rootCmd.Flags().BoolP("chars", "m", false, "prints the character count")
	rootCmd.Flags().BoolP("lines", "l", false, "prints the line count")
	rootCmd.Flags().BoolP("words", "w", false, "prints the word count")
	rootCmd.Flags().BoolP("max-line-length", "L", false, "prints the maximum line length count")
}

//handler for errors
func checkError(e error) error {
	if e != nil {
		return e
	}

	return nil
}

//This function prints the result with the given arguments
func printResult(n int, file string) {
	fmt.Printf("%d %s\n", n, file)
}

//Opens the given file and returns it
func openFile(fileName string) (fp *os.File) {
	file, err := os.Open(fileName)
	checkError(err)
	return file
}

//This function converts a file into a string
func ConvertFileToString(file *os.File) string {
	data, err := ioutil.ReadFile(file.Name())
	checkError(err)
	return string(data)
}

//This function returns the total bytes from a given file
func GetByteCount(data string) int {
	return len(data)
}

//This function returns the total characters from a given file
func GetCharacterCount(data string) int {
	return utf8.RuneCount([]byte(data))
}

//This function returns the total lines from a given file
func GetLineCount(data string) int {
	scanner := bufio.NewScanner(strings.NewReader(data))
	lines := 0

	for scanner.Scan() {
		lines++
	}

	return lines
}

//This function returns the total word count from a given file
func GetWordCount(data string) int {
	scanner := bufio.NewScanner(strings.NewReader(data))
	scanner.Split(bufio.ScanWords)

	words := 0

	for scanner.Scan() {
		words++
	}

	return words
}

//This function returns the maximum line length from a given file
func GetMaxLineLength(data string) int {
	scanner := bufio.NewScanner(strings.NewReader(data))

	longestLine := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > longestLine {
			longestLine = len(line)
		}
	}

	return longestLine
}
