# wc command with golang
A clone of the famous wc Unix command developed with [go](https://golang.org/) & [cobra](https://github.com/spf13/cobra).

### Syntax
```
wcg [OPTION]... [FILE]...
```

With no FILE, when FILE is -, the standard input is read.

### Options

**1. -l or --lines** <br>
This option prints the number of lines present in a file. With this option wc command displays two-columnar output, 1st column shows number of lines present in a file and 2nd itself represent the file name.

**2. -w or --words**<br>
This option prints the number of words present in a file. With this option wc command displays two-columnar output, 1st column shows number of words present in a file and 2nd is the file name.

**3. -c or --bytes** <br>
This option displays count of bytes present in a file. With this option it display two-columnar output, 1st column shows number of bytes present in a file and 2nd is the file name.

**4. -m or --chars** <br>
Using -m option ‘wc’ command displays count of characters from a file.

**5. -L or --max-line-length** <br>
The ‘wc’ command allow an argument -L, it can be used to print out the length of longest (number of characters) line in a file.

**6. --version** <br>
This option is used to display the version of wc which is currently running on your system.

**7. –h or --help** <br>
This option is used to display the help message.

### Example

```
$ wcg -m apple.txt
  34 apple.txt
```
### Compilation/Testing steps

- Run ```go build``` to build the command which leave the binary result in the current working directory.
- Run ```go install``` which builds then installs the package in your $GOPATH/pkg directory.
- Run ```go test``` to run the tests | ```go test -v``` to run the tests with a verbose output.