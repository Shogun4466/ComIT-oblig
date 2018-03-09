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

#### Forklaring: Funksjonen "main" krever et argument fra kommandolinja, slik at den kan kjøre alle de nødvendige mappene og filene. Etter dette implementerer vi lesinga av fila for både .os og info. Vi oppretter så variabler for bytestørrelser og bruker bytes som en grunnverdi int64. "Prinln" kjører vi hver en løkke for å se om fila er enten en directory. Vi bruker os-pakka og modus for å se om fila har Unix permission, append, device file og symbolic link.

## Oppgave 2 
I oppgaven skal vi lage en go-fil med filnavn filecount som skal lese en tekst-fil, skrive ut totalt antall linjer og de fem runene som forekommer mest i fila i medfølgende fil text.txt.<br/> 
Filen skal også bygges om til et kjørbar program og kunen utføres i terminalen.<br/>
Resultatet skal være i format med info om fila, antall linjer i fila, hvilke runer som er fremkalt mest og antallet på dem. 

#### Forklaring: Vi oppretter først to variabler, antall og tegn i form av int og string. Funksjonen inneholder argument for filnavnet og skal ha error hvis den ikke greier å lese eller telle linjene i fila. Vi setter så en utskrift for hva filnavnet heter, hvor mange linjer og de mest forekommende runene i teksten. Vi iterater gjennom alle ASCII-tegnene mellom 32 og 255 (som inneholder runer) ved å lage en løkke. Vi lager så 4 variabler, runesymbol indikerer hvilken hvilke rune som forekommer mest ut ifra løkka, mens tall teller i løkka hver gang hovedløkka går gjennom alle runene. Dette er en bubblesort. 

## Oppgave 3
Oppgave 3 er delt inn i a, b, c, d og e.

I oppgave A skal vi lage en go-fil med navn "addup" og skal inneholde to funksjoner som kommuniserer med hverandre gjennom channels.<br/> 
Funksjon A skal lese inn to tall fra terminalen og sende disse to verdiene til funksjon b.<br/>
Funksjon B skal skal legge sammen de to tallene og sende resultatet tilbake til funksjon A som skriver ut resultatet til terminal.

#### Forklaring: Første funksjon inneholder SIGINT og terminerer etter gitt sekunder. Videre i funksjonen lager den en kanal for x antall sekunder som gir deg tid til å skrive inn tall, der den termineres automatisk etter 10 sekunder. I funksjonen readinput lager vi 2 forskjellige variabler, nemlig de variablene vi velger å skrive inn som input. Den skanner hver verdi og legger inn i channel. Mot slutten sendes kanalen inn som et resultat og printer i form av det ene tallet adderes med det andre = resultat. Siste funksjonen i gofila forklarer sammenhengen mellom nummer, kanalene og resultatet.  

I oppgave B skal vi lage to programmer med navnene addtofile og sumfromfile. Disse programmene skal kommuniserer gjennom en fil.<br/>
Programmet "addtofile" skal lese to tall og skriver disse til en fil.<br/>
Programmet "sumfromfile" skal lese tall fra fil laget av program A og adderer de to tallene. Summen skal dereetter skrives tilbake til samme fil.<br/>
Program A leser deretter resultatet fra fil og skriver ut summen til stdout.

#### Forklaring: Koden er basert fra to programmer. Det første skal da putte inn to tall som da sendes videre til neste fil. Neste fil legger sammen og inneholder et resultat som sendes tilbake til første fil. Første fil printer da ut verdien ut til terminalen. 

I oppgave C skal vi beskrive og implementere en feilhåndtering på alle I/O, både i oppgave A og oppgave B.

#### Forklaring for implementering ligger under egen [Feilhåndtering.md](https://github.com/Shogun4466/ComIT-oblig/blob/master/Oblig-2/src/Oppg-3/feilh%C3%A5ndtering.md) i samme mappe. 

I oppgave D skal vi implementerer en håndtering av SIGINT i både oppgave a og oppgave b. Programmene skal også skrive ut en avslutningsmelding dersom de mottar SIGINT før de fullføres naturlig.

#### Forklaring: Trenger ikke noe spesifisert forklaring, SIGINT kjøres når man trykker Control C. 

I oppgave E skal vi bygge om filene i oppgave 1 og 2 og gjøre dem om til kjørebare filer på operativsystemet (.exe) og legge dem til i en egen bin-mappe.

#### Forklaring: Det ligger kjørbarfiler i BIN, bare å ha nødvendig filer og kjøre via terminalen. 
