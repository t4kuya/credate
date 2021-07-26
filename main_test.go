package main

func Example_help() {
	goMain([]string{"/some/path/of/credate", "-h"})
	// Output:
	// credate [OPTIONS] [FILEs...]
	//   OPTIONS
	//     -b, --byte        Prints the number of bytes in each input file.
	//     -c, --character   Prints the number of characters in each input file.
	//     -l, --line        Prints the number of lines in each input file.
	//     -w, --word        Prints the number of words in each input file.
	//     -h, --help        Prints this message.
	//
	//     ModifyTime in the text file is displayed when no command option is given.
	//     Or, it will be displayed with all options except help.
}

func Example_credate() {
	goMain([]string{"credate", "./testdata/twincle_twincle_little_star.txt", "-w"})
	// Output:
	// ./testdata/twincle_twincle_little_star.txt
	// word: 32
}

func Example_credate2() {
	goMain([]string{"credate", "./testdata/twincle_twincle_little_star.txt", "-l"})
	// Output:
	// ./testdata/twincle_twincle_little_star.txt
	// line: 5
}
