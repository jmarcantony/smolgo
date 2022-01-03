package main

func RemoveNewLine(f []byte) ([]byte, map[int]bool) {
    var (
        prev byte
        m = map[int]bool{}
    )
    for i, v := range f {
        if v == ';' {
            m[i] = true
        }
        if (v == '\n' && v != ';') && (prev != ';' || (prev != '"' && !IsAlphabet(prev))) {
            f[i] = ';'
        }
        prev = v
    }
    return f, m
}
