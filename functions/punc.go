package reloaded

func Punc(arr []string) []string {
	counter := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] == "." || arr[i] == "," || arr[i] == "!" || arr[i] == "?" || arr[i] == ":" || arr[i] == ";" {
			for j := i - 1; ; {
				if arr[j] == "\n" || IsCommand(arr[j]) {
					counter++
					j--
				} else {
					break
				}
			}
			arr[i-counter-1] = arr[i-counter-1] + arr[i]
			arr = append(arr[:i], arr[i+1:]...)
			i--
		}
		counter = 0
	}
	return arr
}
