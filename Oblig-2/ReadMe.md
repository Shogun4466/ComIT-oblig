# Deltagere
Jørgen Christian Arentz Rostrup<br/>
Benjamin Sarpong<br/>
Ørjan Skjerpingstad Mortensen<br/>
Vuong Vo Nguyen<br/>
Tor Ivar Martin Vik

## Oppgave 1
I denne oppgaven skal vi skrive et golang program som skal returnere definert informasjon om en fil, fileinfo. 

Visse krav for oppgaven og hva den skal returnere er om det er/ikke er <br/>
filens infonavn<br/>
filens størrelse ut ifra bytes KB, MB og GB<br/>
om det er en katelog<br/>
om det en regelmessig fil<br/>
om den har Unix tillatelser i bits - lese, skrive funksjoner<br/>
om den kan tilføye noe<br/>
om det er en enhetsfil <br/>
og om den inneholder symbolsk lenke<br/>
I tillegg skal vi bygge en build-program som kan utføres fra kommandolinja og ha et argument for et filnavn

#### Forklaring: Legge inn en forklaring her 


## Oppgave 2 
I oppgaven skal vi lage en go-fil med filnavn filecount som skal lese en tekst-fil, skrive ut totalt antall linjer og de fem runene som forekommer mest i fila i medfølgende fil text.txt.<br/> 
Filen skal også bygges om til et kjørbar program og kunen utføres i terminalen.<br/>
Resultatet skal være i format med info om fila, antall linjer i fila, hvilke runer som er fremkalt mest og antallet på dem. 

#### Forklaring: Legge inn en forklaring her

## Oppgave 3
Oppgave 3 er delt inn i a, b, c, d og e.

I oppgave A skal vi lage en go-fil med navn "addup" og skal inneholde to funksjoner som kommuniserer med hverandre gjennom channels.<br/> 
Funksjon A skal lese inn to tall fra terminalen og sende disse to verdiene til funksjon b.<br/>
Funksjon B skal skal legge sammen de to tallene og sende resultatet tilbake til funksjon A som skriver ut resultatet til terminal.

#### Forklaring: Legge inn en forklaring her

I oppgave B skal vi lage to programmer med navnene addtofile og sumfromfile. Disse programmene skal kommuniserer gjennom en fil.<br/>
Programmet "addtofile" skal lese to tall og skriver disse til en fil.<br/>
Programmet "sumfromfile" skal lese tall fra fil laget av program A og adderer de to tallene. Summen skal dereetter skrives tilbake til samme fil.<br/>
Program A leser deretter resultatet fra fil og skriver ut summen til stdout.

#### Forklaring: Legge inn en forklaring her

I oppgave C skal vi beskrive og implementere en feilhåndtering på alle I/O, både i oppgave A og oppgave B.

#### Forklaring for implementering ligger under egen [Feilhåndtering.md](http://leggetillenke.com) i samme mappe. 

I oppgave D skal vi implementerer en håndtering av SIGINT i både oppgave a og oppgave b. Programmene skal også skrive ut en avslutningsmelding dersom de mottar SIGINT før de fullføres naturlig.

#### Forklaring: Legge inn en forklaring her

I oppgave E skal vi bygge om filene i oppgave 1 og 2 og gjøre dem om til kjørebare filer på operativsystemet (.exe) og legge dem til i en egen bin-mappe.

#### Forklaring: Legge inn en forklaring her
