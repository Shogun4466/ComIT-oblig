# Deltagere
Jørgen Christian Arentz Rostrup, Benjamin Sarpong, Ørjan Skjerpingstad Mortensen, Vuong Vo Nguyen, Tor Ivar Martin Vik

# Oppgave 1
I denne oppgaven skal vi skrive et golang program som skal returnere definert informasjon om en fil, fileinfo. 

Visse krav for oppgaven og hva den skal returnere er om det er/er ikke: 
filens infonavn
filens størrelse ut ifra bytes KB, MB og GB
om det er en katelog
om det en regelmessig fil
om den har Unix tillatelser i bits - lese, skrive funksjoner. 
om den kan tilføye noe
om det er en enhetsfil 
og om den inneholder symbolsk lenke
I tillegg skal vi bygge en build-program som kan utføres fra kommandolinja og ha et argument for et filnavn

Forklaring: 


# Oppgave 2 
I oppgaven skal vi lage en go-fil med filnavn filecount som skal lese en tekst-fil, skrive ut totalt antall linjer og de fem runene som forekommer mest i fila i medfølgende fil text.txt. Filen skal også bygges om til et kjørbar program og kunen utføres i terminalen.
Resultatet skal være i format med info om fila, antall linjer i fila, hvilke runer som er fremkalt mest og antallet på dem. 

Forklaring: 

# Oppgave 3
Oppgave 3 er delt inn i a, b, c, d og e.

I oppgave A skal vi lage en go-fil med navn "addup" og skal inneholde to funksjoner som kommuniserer med hverandre gjennom channels. 
Funksjon A skal lese inn to tall fra terminalen og sende disse to verdiene til funksjon b.
Funksjon B skal skal legge sammen de to tallene og sende resultatet tilbake til funksjon A som skriver ut resultatet til terminal.
 
**Forklaring: *Legg inn tekst* ** 

I oppgave B skal vi lage to programmer med navnene addtofile og sumfromfile. Disse programmene skal kommuniserer gjennom en fil.
Programmet "addtofile" skal lese to tall og skriver disse til en fil.
Programmet "sumfromfile" skal lese tall fra fil laget av program A og adderer de to tallene. Summen skal dereetter skrives tilbake til samme fil. Program A leser deretter resultatet fra fil og skriver ut summen til stdout.

I oppgave C skal vi beskrive og implementere en feilhåndtering på alle I/O, både i oppgave A og oppgave B
**Forklaring: Forklaring for implementering ligger under egen [Feilhåndteringsfila.md](http://leggetillenke.com) i samme mappe.** 

I oppgave D skal vi implementerer en håndtering av SIGINT i både oppgave a og oppgave b. Programmene skal også skrive ut en avslutningsmelding dersom de mottar SIGINT før de fullføres naturlig.
**Forklaring: *Legg til tekst* ** 

e) Bygg filene som i oppgave 1 og 2 til kjørbare filer på ditt operativsystem og legg dem ved i /bin mappen.
