package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

//func main håndterer alle de forskjellige URLene programmet vårt benytter seg av. Først opprettes en router som vi har
//hentet fra gorilla/mux. Deretter sender go-koden brukeren videre til andre funksjonskall avhengig av hvilke URLer
//som skal brukes.
func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	r.HandleFunc("/index", response1)
	r.HandleFunc("/Miljøstasjoner", response2)
	r.HandleFunc("/Kart", response3)
	http.ListenAndServe(":8080", nil);
}

//For tabellvisningen i funksjonen response2 prosesserer go-koden .jsondataene vi henter. Koden under inneholder
//informasjonen som trengs for at html-malen kan lage en tabell.
type Miljø struct {

	Entries []struct {

		Name string `json:"navn"`
		Plastic string `json:"plast"`
		GlasMetal string `json:"glass_metall"`
		Shoe string `json:"tekstil_sko,omitempty"`

	}`json:"entries"`
}

type Test struct{
	Posts int `json:"posts"`
}

var miljø Miljø
var test Test

//funksjonen response1 henter html-malen index.html
func response1(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Print(err)
	}

	err = t.Execute(w, miljø)
	if err != nil {
		log.Fatal(err)
	}
}

//funksjonen response2 henter API-et som inneholder informasjonen over alle miljøstasjoner i Stavanger.
//Deretter prosesseres informasjonen i .json-filen, og blir vist til brukeren gjennom malen miljøstasjoner.html.
func response2(w http.ResponseWriter, r *http.Request) {


	url := "https://hotell.difi.no/api/json/stavanger/miljostasjoner"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &miljø)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	t, err := template.ParseFiles("miljøstasjoner.html")
	if err != nil {
		log.Print(err)
	}

	err = t.Execute(w, miljø)
	if err != nil {
		log.Fatal(err)
	}
}

//funksjonen response3 henter malen kart.html hvor all koden til kart-visningen av informasjonen ligger.
func response3(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("kart.html")
	if err != nil {
		log.Print(err)
	}

	err = t.Execute(w, miljø)
	if err != nil {
		log.Fatal(err)
	}
}
