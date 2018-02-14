package main

import "testing"
import "fmt"

// Her konverterer vi heksadesimaler til string for å kunne putte verdiene inn i løkka. 

func ExtendedASCIIText() string {
	a:= []byte("\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8A\x8B\x8C\x8D\x8E\x8F\x90\x91\x92\x93\x94\x95\x96\x97\x98" +
		"\x99\x9A\x9B\x9C\x9D\x9E\x9F\xA0\xA1\xA2\xA3\xA4\xA5\xA6\xA7\xA8\xA9\xAA\xAB\xAC\xAD\xAE\xAF\xB0\xB1\xB2\xB3" +
			"\xB4\xB5\xB6\xB7\xB8\xB9\xBA\xBB\xBC\xBD\xBE\xBF\xC0\xC1\xC2\xC3\xC4\xC5\xC6\xC7\xC8\xC9\xCA\xCB\xCC\xCD" +
				"\xCE\xCF\xD0\xD1\xD2\xD3\xD4\xD5\xD6\xD7\xD8\xD9\xDA\xDB\xDC\xDD\xDE\xDF\xE0\xE1\xE2\xE3\xE4\xE5\xE6" +
					"\xE7\xE8\xE9\xEA\xEB\xEC\xED\xEE\xEF\xF0\xF1\xF2\xF3\xF4\xF5\xF6\xF7\xF8\xF9\xFA\xFB\xFC\xFD\xFE")
	return string(a)
}

// Her forstår vi ikke syntaxen i function TestExtendedASCII
func TestExtendedASCII(t *testing.T) {
	if !(isExtendedASCII(ExtendedASCIIText())){
		t.Fail()
	}
}

/* Her forsøkte vi å definerer parametre for at c skal kun være innenfor rekkevide 128 og 255, 
men outputen blir ikke riktig og dermed greier vi ikke å kjøre det riktig i terminalen*/

func isExtendedASCII(s string) bool {
	for c := range s {
		fmt.Println(c)
		if c < 128, c > 256 {
			return false
		}
		return true
	}

}
