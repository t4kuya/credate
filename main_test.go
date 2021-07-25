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

func Example_credate3() {
	goMain([]string{"credate", "./testdata/twincle_twincle_little_star.txt", "-c"})
	// Output:
	// ./testdata/twincle_twincle_little_star.txt
	// char: 170
}
