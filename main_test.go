package main

func Example_help() {
	goMain([]string{"/some/path/of/scv", "-h"})
	// Output:
	// scv [OPTIONS] <VECTORS...>
	// OPTIONS
	//     -b   "Prints the number of bytes in each input file."
	//
        //     -m   "Prints the number of characters in each input file."
	//
        //     -l   "Prints the number of lines in each input file."
	//
        //     -w   "Prints the number of words in each input file."
	//
        //     -d   "Prints the creation date of file."
	//
	//     -h, --help                     prints this message.
	// VECTORS
	//     the source of vectors for calculation.
}
