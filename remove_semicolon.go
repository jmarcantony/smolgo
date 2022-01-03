package main

func checkBetween(b byte, i int, f []byte) bool {
	if (i-1 < 0) || (i+1 >= len(f)) {
		return false
	}
	if f[i-1] == '"' || f[i+1] == '"' {
		return false
	}
	if !IsAlphabet(f[i-1]) && !IsAlphabet(f[i+1]) {
		return true
	}
	return false
}

func RemoveSemiColon(f []byte, c map[int]bool) []byte {
	var prev byte
	for i, v := range f {
		if _, ok := c[i]; ok {
			continue
		}
		if v == ';' {
			if prev == ';' || checkBetween(v, i, f) {
				f[i] = 0
			}
		}
		prev = f[i]
	}
	f[len(f)-1] = 0
	return f
}
