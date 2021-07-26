[![build](https://github.com/t4kuya/credate/actions/workflows/build.yml/badge.svg)](https://github.com/t4kuya/credate/actions/workflows/build.yml)
[![Coverage Status](https://coveralls.io/repos/github/t4kuya/credate/badge.svg?branch=main)](https://coveralls.io/github/t4kuya/credate?branch=main)
[![codebeat badge](https://codebeat.co/badges/17166784-c833-4edf-89e9-4a43d306ad8d)](https://codebeat.co/projects/github-com-t4kuya-credate-main)
[![Go Report Card](https://goreportcard.com/badge/github.com/t4kuya/credate)](https://goreportcard.com/report/github.com/t4kuya/credate)
[![Dockerfile](https://img.shields.io/badge/Docker-ghcr.io%2Ft4kuya%2Fcredate%3A1.0.0-green?logo=docker)](https://hub.docker.com/repository/docker/credate/alpine)
[![DOI](https://zenodo.org/badge/370366767.svg)](https://zenodo.org/badge/latestdoi/370366767)
# credate
credate is a improved version of wc.

# Table of Contents
- [Description](#Description)
- [Usage](#Usage)
- [Icon](#Icon)

# Description
The traditional "word count" command counts the number of bytes, lines, words, and characters in a text file.
In addition to the "word count" command, "credate" tells you the modification date of the text file.

The original plan was to display ModifyTime with the option "-d", but it is now displayed only when no options are specified or when all options except "-h" and "--help" are specified.
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
# Icon
![km206143521605125200317](https://user-images.githubusercontent.com/84721993/119422504-392efc00-bd3c-11eb-8752-0f3b7403f648.png)
<br>https://www.iconpon.com/
