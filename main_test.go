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
	//     -d, --date        Prints the creation date of file.
	//     -h, --help        Prints this message.
}

func Example_edkd() {
	goMain([]string{"credate", "./testdata/twincle_twincle_little_star.txt"})
	// Output:
	// testdata/twincle_twincle_little_star.txt
	//line: 5   byte: 170   char: 170   word: 32   ModifyTime: 2021-07-22 16:02:47.264973348 +0900 JST
}
