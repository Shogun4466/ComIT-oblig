package main

import "fmt"

//func main printer strengen fra func ExtendedASCIIText

func main() {
	fmt.Printf (ExtendedASCIIText())
}

//func ExtendedASCIIText genererer strengen "a" fra byte-verdiene som er input og
//returnerer verdiene som en streng.

func ExtendedASCIIText () string {
	a := []byte("\xE2\x82\xAC \xC3\xB7 \xC2\xBE \x64\x6F\x6C\x6C\x61\x72")
	return string(a)
}
