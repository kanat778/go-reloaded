package reloaded

func AtoiBase(s string, base string) int {
	res := 0
	index := 0
	lent := len(s)
	if len(base) < 2 {
		return 0
	}
	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return 0
			}
		}
	}
	for _, symbol := range base {
		if symbol == '-' || symbol == '+' {
			return 0
		}
	}
	for _, letter := range s {
		for i := 0; i < len(base); i++ {
			if rune(base[i]) == letter {
				index = i
				break
			}
		}
		res = res + index*IterativePower(len(base), lent-1)
		lent--
	}
	return res

}

func IterativePower(nb int, power int) int {
	if power == 0 {
		return 1
	} else if power < 0 {
		return 0
	} else {
		res := 1
		for i := 0; i < power; i++ {
			res = res * nb
		}
		return res
	}
}
