package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func helpMessage(originalProgramName string) string {
	name := filepath.Base(originalProgramName)
	return fmt.Sprintf(`%s [OPTIONS] [FILEs...]
	OPTIONS
    -b, --byte        Prints the number of bytes in each input file.

    -c, --character   Prints the number of characters in each input file.

    -l, --line        Prints the number of lines in each input file.

    -w, word          Prints the number of words in each input file.

    -d, --date        Prints the creation date of file.

    -h, --help        Prints this message.
`, name)
}

func goMain(args []string) int {
	fmt.Println(helpMessage(args[0]))
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
