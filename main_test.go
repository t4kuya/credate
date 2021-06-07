package main

func Example_help() {
	goMain([]string{"/some/path/of/credate", "-h"})
	// Output:
	// ccat [OPTIONS]  [FILEs...]
	// OPTIONS
	// 	-h, --help              このメッセージを出力します．
	// ARGUMENTS
	// 	FILEs...                中身の確認または結合を行うファイル．
}
