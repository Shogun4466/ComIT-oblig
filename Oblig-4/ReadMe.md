 <h1>Systemspesifikasjoner:</h1>

Programmet vårt er lagd for å vise innbyggerne i Stavanger hvor de kan finne Miljøstasjoner. Dette vises på nett i enten tabell form eller via Google maps. Man går frem ved å starte programmet for så å skrive localhost:8080/index i nettleseren, herfra kan man velge hvordan man vil se informasjonen. Vi vil hjelpe Stavangers innbyggere med å bli mer miljøvennelige! 



<h1>Systemarkitekrtur</h1> <a href="https://imgur.com/6s6B2DN"><img src="https://i.imgur.com/6s6B2DN.png" title="source: imgur.com" /></a>

Vårt programm henter nåtids informasjon fra to APIer liggende i https://hotell.difi.no/api/json/stavanger/miljostasjoner og https://maps.googleapis.com/maps/api/js?key=AIzaSyAyb_m6VmJ0foIEpxxmJ7ubQISfWV2TfX4&callback=initMap. Informasjonen hentet fra API-et blir brukt til å fremvise informasjon om miljøstasjoner i Stavanger området. Når programmet tas i bruk vill API dataen bli vist på en av to mulige måter. Visnings måten er basert på valget til brukeren mellom to knapper. Den første visningen direkterer brukeren til miljøstasjoner.html. Denne siden setter all API dataen fra miljostasjoner API-et utenom latitude og longitude inn i en tabell. Den andre visningen sneder brukeren til kart.html og bruker googlemaps API-et for å gi en kart visnings funksjonalitet. Den bruker også verdiene latitude og longitude til å markere alle miljøstasjonene på kartet.
