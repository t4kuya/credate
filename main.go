package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func helpMessage(originalProgramName string) string {
	name := filepath.Base(originalProgramName)
	return fmt.Sprintf(`%s [OPTIONS] <VECTORS...>
OPTIONS
    -b,   Prints the number of bytes in each input file.

    -m,   Prints the number of characters in each input file.
    
    -l,   Prints the number of lines in each input file.

    -w,   Prints the number of words in each input file.

    -d,   Prints the creation date of file.
    
    -h, --help                     prints this message.
VECTORS
    the source of vectors for calculation.`, name)
}

func goMain(args []string) int {
	fmt.Println(helpMessage(args[0]))
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
