package reloaded

func BracketsFix(arr []string) []string {
	counter := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] == "(" || arr[i] == "[" || arr[i] == "{" {
			for j := i + 1; ; {
				if arr[j] == "\n" || IsCommand(arr[j]) {
					counter++
					j++
				} else {
					break
				}
			}
			arr[i] = arr[i] + arr[i+counter+1]
			arr = append(arr[:i+1], arr[i+counter+2:]...)
			i--
		}

		counter = 0
	}

	for k := 0; k < len(arr); k++ {
		if arr[k] == ")" || arr[k] == "]" || arr[k] == "}" {
			for v := k - 1; ; {
				if arr[v] == "\n" || IsCommand(arr[v]) {
					counter++
					v--
				} else {
					break
				}
			}
			arr[k-counter-1] = arr[k-counter-1] + arr[k]
			arr = append(arr[:k], arr[k+1:]...)
			k--
		}
		counter = 0
	}
	return arr
}
