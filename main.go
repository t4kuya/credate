package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
	"unicode/utf8"

	flag "github.com/spf13/pflag"
)

type options struct {
	bytet     bool
	character bool
	word      bool
	line      bool
	help      bool
	args      []string
}

type cal struct {
	bytes int
	chars int
	words int
	lines int
}

func helpMessage(originalProgramName string) string {
	name := filepath.Base(originalProgramName)
	return fmt.Sprintf(`%s [OPTIONS] [FILEs...]
  OPTIONS
    -b, --byte        Prints the number of bytes in each input file.
    -c, --character   Prints the number of characters in each input file.
    -l, --line        Prints the number of lines in each input file.
    -w, --word        Prints the number of words in each input file.
    -d, --date        Prints the creation date of file.
    -h, --help        Prints this message.`, name)
}

func calculate(o *options, filenames []string) {
	var c cal
	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			continue
		}
		a, _ := ioutil.ReadAll(file)
		c.lines = bytes.Count(a, []byte{'\n'})
		c.bytes = len(a)
		c.chars = utf8.RuneCountInString(string(a))
		c.words = len(bytes.Fields(a))

		printer(o, c, filename)

	}
}

func parseArgs(args []string) (*options, error) {
	opts := &options{}
	flags := flag.NewFlagSet("credate", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args[0])) }
	flags.BoolVarP(&opts.bytet, "byte", "b", false, "number of bytes")
	flags.BoolVarP(&opts.character, "character", "c", false, "number of character")
	flags.BoolVarP(&opts.word, "word", "w", false, "number of words")
	flags.BoolVarP(&opts.line, "line", "l", false, "number of lines")
	flags.BoolVarP(&opts.help, "help", "h", false, "prints help")

	if err := flags.Parse(args); err != nil {
		return nil, err
	}

	if !opts.bytet && !opts.character && !opts.word && !opts.line {
		opts.bytet = true
		opts.character = true
		opts.word = true
		opts.line = true
	}
	opts.args = flags.Args()[1:]

	return opts, nil
}

//get modify time of files
func statTimes(name string) (mtime time.Time, err error) {
	fi, err := os.Stat(name)
	if err != nil {
		return
	}
	mtime = fi.ModTime()
	return
}

func printer(o *options, c cal, filename string) {
	name := filename
	mtime, err := statTimes(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(filename)
	if o.line {
		fmt.Printf("line: %d   ", c.lines)
	}
	if o.bytet {
		fmt.Printf("byte: %d   ", c.bytes)
	}
	if o.character {
		fmt.Printf("char: %d   ", c.chars)
	}
	if o.word {
		fmt.Printf("word: %d   ", c.words)
	}
	if o.line && o.bytet && o.character && o.word {
		fmt.Println("ModifyTime:", mtime)
	}

}

func goMain(args []string) int {
	opts, err := parseArgs(args)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	if opts.help {
		fmt.Println(helpMessage(args[0]))
		return 1
	}
	calculate(opts, opts.args)

	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
