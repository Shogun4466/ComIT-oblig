I filen addup.go implementerer vi følgende feilhåndtering på input:

	fmt.Println("Enter number to add to the first number: ")
	fmt.Scan(&number2)
	intNumber2, err := strconv.Atoi(number2)
	if err != nil{
		fmt.Println("Please input a number. Shutting down.")
		os.Exit(0)
	}

Denne koden sørger for at input fra bruker er tall som kan legges sammen
av funksjonen i koden. Programmet avslutter dersom input ikke er tall.


I filen addtofile.go ligger feilhåndteringen i hovedsak i følgende kode i tillegg til tilsvarende
feilhåndtering som i addup.go:

	file, err := os.Create("result.txt")
	if err != nil {
		log.Fatal("Failed to create file", err)
	}
	defer file.Close()

	f, err := os.OpenFile("result.txt",os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open file", err)
	}
	if _, err := fmt.Fprintf(f, "%d\n%d", intNumber1, intNumber2); err != nil {
		log.Fatal("Failed to write to file", err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
  
  Først lages filen result.txt på nytt ved hver kjøring av programmet, og vi har alltid en blank fil. 
  Dersom programmet ikke kan lage en fil vil det rapportere "Failed to create file". Programmet åpner
  så den nye filen. Dersom det oppstår en feil her, vil programmet rapportere "Failed to open file".
  Her spesifiserer vi os.O_WRONLY slik at vi er sikre på at den nye og åpne filen kan skrives i.
  Deretter skriver programmet verdiene bruker har definert til filen. Ved feil rapporteres "Failed to
  write to file". Til slutt lukkes filen slik at den er klar for å bli skrevet i av sumfromfile.go.
  Programmet har også en feilhåndteringsfunksjon:
  
  func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
 
 Denne funksjonen er implementert i koden som leser sluttresultatet fra result.txt, og lar kjøringen
 programmet gå til den innebygde panikkfunksjonen i golang.
 
  
  I sumfromfile.go ser feilhåndteringen til filhåndteringen slik ut:
  
  	file, err := os.OpenFile("result.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open file", err)
	}
	if _, err := fmt.Fprintf(file,"\n%d\n", result); err != nil {
		log.Fatal("Failed to write to file", err)
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
  
  Feilhåndteringen er så å si lik, men her blir filen åpnet med os.O_APPEND i tillegg til O_WRONLY.
  Filen er ikke lengre tom, og vi ønsker å beholde informasjonen i den for å printe input som en del
  av resultatet. Derfor må vi bruke append for å legge til den nye informasjonen etter den eksisterende.
