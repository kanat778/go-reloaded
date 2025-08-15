package reloaded

import (
	"regexp"
	"strconv"
)

func CommandInfo(s string) (string, int) {

	out_str := ""

	if s[:3] == "(up" {
		out_str = "up"
	}
	if s[:4] == "(hex" {
		out_str = "hex"
	}
	if s[:4] == "(bin" {
		out_str = "bin"
	}
	if s[:4] == "(low" {
		out_str = "low"
	}
	if s[:4] == "(cap" {
		out_str = "cap"
	}

	dig := regexp.MustCompile(`,\d+\)`)
	out_int := 0

	if s == "(up)" || s == "(hex)" || s == "(bin)" || s == "(low)" || s == "(cap)" {
		out_int = 1
	} else {
		temp := dig.Find([]byte(s))
		if temp != nil {
			out_int, _ = strconv.Atoi(string(temp[1 : len(temp)-1]))
		}
	}

	return out_str, out_int
}
