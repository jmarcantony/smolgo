package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        log.Fatalln("Pass a file path to minify")
    }
    filename := args[0]
    if !strings.HasSuffix(filename, ".go") {
        log.Fatalln("Not a go file!")
    }
    file, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }
    file, c := RemoveNewLine(file)
    file = RemoveSpaces(file)
    file = RemoveSemiColon(file, c)
    var f []byte
    for _, v := range file {
        if v != 0 {
            f = append(f, v)
        }
    }
    for _, b := range f {
        fmt.Printf("%c", b)
    }
}
