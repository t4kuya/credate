---
title: "Usage"
---

# Usage
```
credate [OPTION] [FILEs...]
 OPTIONS
   -b, --byte        Prints the number of bytes in each input file.
   -c, --character   Prints the number of characters in each input file.
   -l, --line        Prints the number of lines in each input file.
   -w, --word        Prints the number of words in each input file.
   -h, --help        Prints this message.

   ModifyTime in the text file is displayed when no command option is given.
   Or, it will be displayed with all options except help.
```

# Example Output
```
$ credate testdata/twincle_twincle_little_star.txt
testdata/twincle_twincle_little_star.txt
line: 5   byte: 170   char: 170   word: 32   ModifyTime: 2021-07-22 16:02:47.264973348 +0900 JST

$ credate -l -b testdata/twincle_twincle_little_star.txt
testdata/twincle_twincle_little_star.txt
line: 5   byte: 170 

$ credate -h
credate [OPTIONS] [FILEs...]
  OPTIONS
    -b, --byte        Prints the number of bytes in each input file.
    -c, --character   Prints the number of characters in each input file.
    -l, --line        Prints the number of lines in each input file.
    -w, --word        Prints the number of words in each input file.
    -h, --help        Prints this message.

    ModifyTime in the text file is displayed when no command option is given.
    Or, it will be displayed with all options except help.
```