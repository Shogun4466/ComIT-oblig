## Oppgave 1 

I oppgave 1 skal vi lage en filnavn som kjører en webserver med portnr 8080. Oppgave1.go lager en lokal server som ser etter en hvilken som helst input. Når noen kontakter serveren via port 8080 vil den svare "Hello, client".

Dette er testet både med localhost:8080 og ved å bruke to forskjellige maskiner ved å bruke lokale IP-adresser.

## Oppgave 2 

I oppgave 2 skulle vi presentere data fra ulike API-er som inneholt JSON data. Dataen skulle bli omformulert til et format som var leslig for folk flest og bli gitt tilgang til via /1,/2,/3,/4,/5. I tillegg skulle disse stiene renderes vie go template. 

Det første som blir gjort i koden er importen av en ekstern router for å håndtere "pathinga" til koden. Så setter vi opp fem hånderinger som refererer til senere funksjoner. Etter dette sier vi at programmet skal lytte til port 8080. Vi lager så fem structs som håndterer JSON dataene fra de fem forskjellige url-ene vi definerer senere. Vi har så fem variabler som inneholder hver struct. I funksjonene sier vi hvilke url-er JSON dataene skal hentes fra, "unmarshal-er" JSON dataen og refererer til noe1.html-noe5.html som er ansvarlig for oppsettet av dataene.  

De fem html-filene er ansvarlig for oppsettet og presentasjonen av den "unmarshal-ede" JSON dataene. 

## Oppgave 3 

I oppgave 3

