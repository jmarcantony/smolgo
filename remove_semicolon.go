package main

/*
package main;;import(;"fmt";"math";);;;func squareNum(i int){;fmt.Println(math.Pow(float64(i),2));};func main(){;for i:=0;i<11;i++{;squareNum(i);};};
*/

func checkBetween(b byte, i int, f []byte) bool {
    if (i-1 < 0) || (i + 1 >= len(f)) {
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
    return f
}
