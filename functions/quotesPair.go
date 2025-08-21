package reloaded

func NewlineAndCommandIndforw(arr []string, i int) int {
	j := i + 1
	for ; j < len(arr); j++ {
		if arr[j] == "'" {
			return -1
		}
		if arr[j] != "\n" && !IsCommand(arr[j]) {
			break
		}
	}
	return j
}

func NewlineAndCommandIndback(arr []string, i int) int {
	j := 0
	if i > 0 {
		j = i - 1
	}
	for ; j >= 0; j-- {
		if arr[j] != "\n" && !IsCommand(arr[j]) {
			break
		}
	}
	return j
}

func QuotesFix(arr []string) []string {
	counter := 0
	j := 0
	k := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == "'" {
			counter++
			if counter%2 == 1 {
				j = NewlineAndCommandIndforw(arr, i)
				if j == -1 {
					continue
				} else {
					arr[j] = "'" + arr[j]
					arr = append(arr[:i], arr[i+1:]...)
					i--
				}

			} else {
				k = NewlineAndCommandIndback(arr, i)
				arr[k] = arr[k] + arr[i]
				arr = append(arr[:i], arr[i+1:]...)
				i--
			}
		}
	}
	return arr
}
