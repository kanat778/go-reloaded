package reloaded

import "regexp"

func IsCommand(s string) bool {
	up := regexp.MustCompile(`\(([^)]*(up)[^)]*)\)`)
	low := regexp.MustCompile(`\(([^)]*(low)[^)]*)\)`)
	cap1 := regexp.MustCompile(`\(([^)]*(cap)[^)]*)\)`)
	bin := regexp.MustCompile(`\(([^)]*(bin)[^)]*)\)`)
	hex := regexp.MustCompile(`\(([^)]*(hex)[^)]*)\)`)

	if up.Match([]byte(s)) {
		return true
	}

	if low.Match([]byte(s)) {
		return true
	}

	if cap1.Match([]byte(s)) {
		return true
	}

	if bin.Match([]byte(s)) {
		return true
	}

	if hex.Match([]byte(s)) {
		return true
	}

	return false
}
