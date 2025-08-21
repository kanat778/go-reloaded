package reloaded

func QuotesPair(arr []string) bool {
	counter := 0

	for _, str := range arr {
		if str == "'" {
			counter++
		}
	}
	if counter%2 == 0 {
		return true
	}
	return false
}
