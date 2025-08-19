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

	strarr := reloaded.BytesToStrArray(bytes)
	//fmt.Println("string array = ", strarr)
	strarr = reloaded.CommandFix(strarr)
	strarr = reloaded.ArticleFix(strarr)
	//fmt.Println("ArticleFix = ", articleFix)
	fmt.Println("command Fix = ", strarr)

	for ind, str := range strarr {
		if reloaded.IsCommand(str) {
			if ind == 0 {
				fmt.Println("Command can be applied only to word before it")
				return
			}
			command, index := reloaded.CommandInfo(str)

			if command == "bin" && index != 1 {
				fmt.Println("Error: bin command can not handle any parameters")
				return
			}
			if command == "hex" && index != 1 {
				fmt.Println("Error: hex command can not handle any parameters")
				return
			}
			if index == 0 {
				fmt.Println("Error: invalid command parameter. Please use a positive number as the parameter")
				return
			} else if index > ind {
				fmt.Println("Error: used invalid command parameter. Please check if parameter is bigger than words quantity before command")
				return
			} else {
				for i := ind; i > 0 && index > 0; i-- {
					if !reloaded.IsPunc(strarr[i-1]) {
						if command == "up" {
							strarr[i-1] = strings.ToUpper(strarr[i-1])
							index--
						}
						if command == "low" {
							strarr[i-1] = strings.ToLower(strarr[i-1])
							index--
						}

						if command == "cap" {
							strarr[i-1] = strings.ToLower(strarr[i-1])
							strarr[i-1] = strings.Title(strarr[i-1])
							index--
						}
					}
				}
			}
		}
	}
	fmt.Println(strarr)

	resHex := reloaded.AtoiBase("4F", "0123456789ABCDEF")
	fmt.Println("hex =", resHex)

	resBin := reloaded.AtoiBase("10", "01")
	fmt.Println("bin =", resBin)
}
