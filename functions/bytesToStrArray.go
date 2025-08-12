package reloaeded

func BytesToStrArray(bytes []byte) []string {
	slice := []string{}
	str := ""

	for _, v := range bytes {
		if v == '.' || v == ',' || v == '!' || v == '?' || v == ':' || v == ';' || v == '(' || v == ')' || v == '\'' || v == '"' {
			if len(str) != 0 {
				slice = append(slice, str)
				str = ""
			}
			slice = append(slice, string(v))
			continue
		}
		if v != ' ' && v != '\n' && v != '	' {
			str += string(v)
		} else {
			if len(str) != 0 {
				slice = append(slice, str)
				str = ""
			}
		}

		if v == '\n' {
			slice = append(slice, string(v))
		}
	}
	if len(str) != 0 {
		slice = append(slice, str)
		str = ""
	}

	return slice
}
