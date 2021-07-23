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
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

const helpText = `
 wcg wc --help
 Usage: wcg wc [OPTION]... [FILE]...
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
                       	NUL-terminated names in file F;
                       	If F is - then read names from standard input
  -L, --max-line-length  print the maximum display width
  -w, --words        	print the word counts
  	--help 	display this help and exit
  	--version  output version information and exit

 GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
 Full documentation <https://www.gnu.org/software/coreutils/wc>
 or available locally via: info '(coreutils) wc invocation'
 `

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wcg",
	Short: "A clone of the famous linux wc command",
	Long:  "Prints newline, word, and byte counts for each FILE, and a total line if more than one FILE is specified",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},

	Version: "1.0.0",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {fmt.Println("Hello!")},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.SetHelpTemplate(helpText)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".wc-golang" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".wc-golang")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
