package reloaded

func CommandFix(input []string) []string {
	start_ind := 0
	end_ind := 0

	// ["it" "(" "cap", "\n" ")" "was" "the" "best" "of" "times" "."]

	// ["it" "(" "cap", "," "1" ")" "was" "the" "best" "of" "times" "."]
	for i := 0; i < len(input); i++ {
		if input[i] == "(" && (input[i+1] == "up" || input[i+1] == "low" || input[i+1] == "hex" || input[i+1] == "bin" || input[i+1] == "cap") {
			start_ind = i
			for j := i + 1; j < len(input); {
				if input[j] == "\n" {
					j++
					continue
				}
				if input[j] != ")" {
					input[i] = input[i] + input[j]
					j++
				} else {
					input[i] = input[i] + input[j]
					end_ind = j
					break
				}
			}
			input = append(input[:start_ind+1], input[end_ind+1:]...)

		}
	}

	return input
}

//input[i] = "(" + "cap"
