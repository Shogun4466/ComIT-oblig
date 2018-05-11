<h1>Gruppemedlemmer:</h1>
Benjamin Sarpong
<br>Ørjan Mortensen
<br>Vuong Nguyen
<br>Tor Ivar Martin Vik
<br>Jørgen Chr. A. Rostrup

<h1>Systemspesifikasjoner:</h1>

Programmet vårt er lagd for å vise innbyggerne i Stavanger hvor de kan finne Miljøstasjoner. Dette vises på nett i enten tabell form eller via Google maps. Man går frem ved å starte programmet for så å skrive localhost:8080/index i nettleseren, herfra kan man velge hvordan man vil se informasjonen. Brukere kan selv velge om de vil se en oversikt over miljøstasjonene i en tabell, eller om de vil se hvor stasjonene er i google maps. Vi vil hjelpe Stavangers innbyggere med å bli mer miljøvennelige!


<h1>Systemarkitekrtur</h1> <a href="https://imgur.com/6s6B2DN"><img src="https://i.imgur.com/6s6B2DN.png" title="source: imgur.com" /></a>

Vårt program henter sanntids-informasjon fra to APIer. Informasjonen hentet fra API-ene blir brukt til å fremvise informasjon om miljøstasjoner i Stavanger. Når programmet tas i bruk vil API-dataen bli vist på en av to mulige måter. Visningsmåten er basert på valget til brukeren. Brukeren velger visning med en av to knapper. Den første visningen direkterer brukeren til miljøstasjoner.html. Denne siden setter all API dataen fra miljostasjoner API-et utenom latitude og longitude inn i en tabell. Den andre visningen sender brukeren til kart.html og bruker googlemaps API-et for å gi en kartvisningsfunksjonalitet. Den bruker også verdiene latitude og longitude til å markere alle miljøstasjonene på kartet.
