# Oppgave 1


## Fyll ut manglende tall i tabell:
  <table style="width: 687px;" border="1">
  <tbody>
  <tr>
<td style="width: 250px;">
<p>Binære tall (mest signifikant bit til venstre</p>
</td>
<td style="width: 250px;">
<p>Hexadesimaltall</p>
</td>
<td style="width: 250px;">
<p>Desimaltall</p>
</td>
</tr>
<tr>
<td style="width: 200px; text-align: center;">
<p>1101</p>
</td>
<td style="width: 200px; text-align: center;">
<p>0xD</p>
</td>
<td style="width: 200px; text-align: center;">
<p>13</p>
</td>
</tr>
<tr>
<td style="width: 265px; text-align: center;">
<p>110111101010</p>
</td>
<td style="width: 204px; text-align: center;">
<p>0xDEA</p>
</td>
<td style="width: 195px; text-align: center;">
<p>3562</p>
</td>
</tr>
<tr>
<td style="width: 265px; text-align: center;">
<p>1010111100110100</p>
</td>
<td style="width: 204px; text-align: center;">
<p>0xAF34</p>
</td>
<td style="width: 195px; text-align: center;">
<p>44852</p>
</td>
</tr>
<tr>
<td style="width: 265px; text-align: center;">
<p>1111111111111111</p>
</td>
<td style="width: 204px; text-align: center;">
<p>0xFFFF</p>
</td>
<td style="width: 195px; text-align: center;">
<p>65535</p>
</td>
</tr>
<tr>
<td style="width: 265px; text-align: center;">
<p>10001011110001010</p>
</td>
<td style="width: 204px; text-align: center;">
<p>0x1178A</p>
</td>
<td style="width: 195px; text-align: center;">
<p>71562</p>
</td>
</tr>
</tbody>
</table>


## Oppgave A

Før vi går inn på metodene vi bruker for å konvertere binær-, hexadesimal- og desimaltall må vi se på hvilke tall som ligger til grunn for hvert enkelt tallsystem. Binærtallsystemet har 2 som grunntall. det vil si at systemet har kun to siffer før det går videre til neste tallposisjon. Binære tall har 10 som grunntall, og hexadesimaltall har 16 som grunntall. Desimalene for de seks siste desimalene vi ikke har siffer for går fra A til F. For å gjøre det klart hvilken talltype vi oppererer med, kan vi notere grunntallet etter et tall. Eksempel: 1101<sub>2</sub>, 13<sub>10</sub>, D<sub>16</sub>

Vi vil nå gå gjennom metoder for konvertering mellom talltypene:


#### Desimal til binær:
Del tallet på to. Dersom det blir en rest, det vil si at svaret ender med .5, noterer vi 1. Derson du får et heltall noterer vi 0. Fortsett så denne prosessen med svartallet fra delingen til du har 0. Det første som ble notert er det siste sifferet i sifferrekken som danner det binære tallet.

#### Hexadesimal til binær:
Her tar vi utgangspunkt i hvert enkelt siffer i det hexadesimale systemet. Hvert enkelt siffer blir konvertert til sin binære motpart. Dersom binærtallet har færre enn 4 siffer, legges 0 til i front helt til hvert binærtall består av 4 siffer. For eksempel noteres 3<sub>10</sub> som 0011<sub>2</sub>. Når du har funnet binærtallet til hvert enkelt hexadesimalsiffer, konkateneres binærtallene i samme sifferrekkefølge som det orginale hexadesimaltallet.

#### Desimal til hexadesimal:
Metoden er lik Desimal til binær, men vi deler på 16. Resttallet oversettes til den hexadesimale representasjonen av tallet. Det siste hexadesimale sifferet som blir notert er det første sifferet i det endelige hexadesimale tallet.

#### Binær til hexadesimal:
Metoden er lik som hexadesimal til binær, bare motsatt. Det binære tallet deles opp i grupper på 4 siffer. Disse gruppene kan alltid oversettes til ett enkelt hexadesimalsiffer. Dersom det siste binære tallet ikke har 4 siffer, legges det igjen til 0-er til tallet består av fire siffer. De oppdelte binærtallene oversettes til hexadesimalsiffre og settes sammen til hesadesimaltall i samme rekkefølge som det ble delt opp.

#### Binær til desimal:
Her tar vi hvert enkelt siffer i binærsifferrekken og bruker følgende formel, hvor n=posisjon i binærsifferrekken: siffer*2<sup>n</sup>
Merk at den bakerste posisjonen er definert som posisjon 0. Dette vil resultere i at 0-ene ikke returnerer noen verdi, og 1-erene returnerer en verdi. Verdiene som blir returnert adderes, og vi får desimaltallet.

#### Hexadesimal til desimal:
Samme formel som Binær til desimal, men 2 byttes til 16: siffer*16<sup>n</sup>
Tallene adderes igjen til slutt for å finne desimaltallet.

### Kilder:
https://www.uio.no/studier/emner/matnat/ifi/INF1000/h15/undervisningsmateriale/andre-tallsystemer-(matematisk).pdf
