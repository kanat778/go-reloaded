package reloaded

import "fmt"

func CommandFix(input []string) []string {
	//start_ind := 0
	//end_ind := 0

	for i := 0; i < len(input); i++ {
		if input[i] == "(" && (input[i+1] == "up" || input[i+1] == "low" || input[i+1] == "hex" || input[i+1] == "bin" || input[i+1] == "cap") {
			fmt.Println("AAAAAAAAAAAAAAAAA")
		}
	}

	return input
}
