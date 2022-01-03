package main

func checkBetween(i int, f []byte) (r bool) {
	if (i-1 < 0) || (i+1 >= len(f)) {
		return
	}
	if f[i-1] == '"' || f[i+1] == '"' || f[i-1] == '`' || f[i+1] == '`' || f[i-1] == '\'' || f[i+1] == '\'' {
		return
	}
	if !IsAlphabet(f[i-1]) && !IsAlphabet(f[i+1]) {
		r = true
	}
	return
}

func RemoveSemiColon(f []byte, c map[int]bool) []byte {
	var prev byte
	for i, v := range f {
		if v == ';' {
			if _, ok := c[i]; ok {
				prev = f[i]
				continue
			}
			if prev == ';' || checkBetween(i, f) {
				f[i] = 0
			}
		}
		prev = f[i]
	}
	if f[len(f)-1] == ';' {
		f[len(f)-1] = 0
	}
	return f
}
