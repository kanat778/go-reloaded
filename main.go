package main

import (
	"fmt"
	"os"
	"reloaded/functions"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("ERROR: Invalid number of arguments!")
		return
	}

	if args[0] == args[1] {
		fmt.Println("ERROR: Invalid filenames!")
		return
	}

	fmt.Println(args)

	if !strings.Contains(args[0], ".txt") || !strings.Contains(args[1], ".txt") {
		fmt.Println("ERROR: Invalid datatype of arguments.Please use .txt notation")
		return
	}

	_, err := os.Stat(args[0])
	if err != nil {
		os.Create(args[0])
	}

	bytes, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Println("ERROR: Not found the file!")
		err.Error()
	}

	fmt.Println(bytes)

	strarr := reloaded.BytesToStrArray(bytes)
	fmt.Println(strarr)
	strarr = reloaded.CommandFix(strarr)
	fmt.Println(strarr)

}
