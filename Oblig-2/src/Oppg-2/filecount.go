package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"io/ioutil"
	"strings"
)

var count []int
var runecount []string

func main() {

	filename := os.Args[1]

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	lines, err := countLinesInFile(filename)
	if err != nil {
		os.Exit(2)
	}

	fmt.Println("\nInformation about",filename,":")
	fmt.Println("Number of lines in file:",lines)
	fmt.Println("\nMost common runes:")
	str := string(b)
	for i := 32; i <= 255; i++ {
		iterate(i, str)
	}

	var runesymbol string
	var tall= 0
	var n int = 0
	var a int = 0

	for i := 0; i < 5; i++ {
		tall = 0
		for l := 0; l < len(count); l++ {
			n = count[l]
			if n > tall {
				tall = n
				a = l
			}
		}
		runesymbol = runecount[a]
		count[a] = 0
		fmt.Println(i+1,"Rune:", runesymbol, ",", "Counts:", tall)
	}
}

func iterate(tall int, tekst string) {
	st := string(tall)
	symbol := fmt.Sprintf("%s", st)
	countstring := strings.Count(tekst, symbol)
	count = append(count, countstring)
	runecount = append(runecount, symbol)
}


func countLinesInFile(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	buf := make([]byte, 1024)
	lines := 0

	for {
		readBytes, err := file.Read(buf)
		if err != nil {
			if readBytes == 0 && err == io.EOF {
				err = nil
			}
			return lines, err
		}
		lines += bytes.Count(buf[:readBytes], []byte{'\n'})
	}
	return lines, nil
}
