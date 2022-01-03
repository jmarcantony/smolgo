package main

/*
   65 - 90  | A - Z
   97 - 122 | a - z
*/
func IsAlphabet(b byte) (l bool) {
	if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') {
		l = true
	}
	return
}

func RemoveSpaces(f []byte) []byte {
	var prev byte
	for i, v := range f {
		if v == ' ' {
			if i+1 < len(f) {
				if !IsAlphabet(f[i+1]) {
					f[i] = 0
					continue
				}
			}
			if !IsAlphabet(prev) {
				f[i] = 0
			}
		}
		prev = v
	}
	return f
}
