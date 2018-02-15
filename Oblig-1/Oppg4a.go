package main

import "fmt"

//func main henter output fra func makeText for så bruke denne stringen i løkken til
//func iterateOverExtendedASCIIStringLiteral for å generere svaret på oppgaven.

func main(){
	text := makeText()
	iterateOverExtendedASCIIStringLiteral(text)
}

/* makeText lager først en tom slice, så genereres verdiene fra 0x80 til 0xFF og settes
inn i slice "a". Konverterer så byteverdiene i slicen til en string-verdi.
*/

func makeText()string{
	var a []byte
	for i := 0x80; i <= 0xFF; i++{
		a = append(a, byte(i))
	}
	extAscii := string(a)
	return extAscii
}

/* iterateOverExtendedASCIIStringLiteral kjører en løkke hvor hver enkelt verdi i strengen
fra makeText blir skrevet ut som hex-verdi, som karakter og som binær-verdi.
 */

func iterateOverExtendedASCIIStringLiteral(a string){
	for i := 0; i < len(a); i++ {
		fmt.Printf("%X %c %b \n", a[i], a[i], a[i])
	}
}
