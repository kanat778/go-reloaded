package reloaded

import "regexp"

func IsPunc(s string) bool {
	com := regexp.MustCompile(`[?:;,.!\n()'{}\[\]"]`)
	return com.Match([]byte(s))
}
