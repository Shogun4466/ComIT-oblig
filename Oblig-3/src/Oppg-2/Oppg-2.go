package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)
//Creates a router of type mux from gorillatoolkits.org. The router handles url requests and matches them with phats I predetermine.
func main() {

	r := mux.NewRouter()
	http.Handle("/", r)
	r.HandleFunc("/1", response1)
	r.HandleFunc("/2", response2)
	r.HandleFunc("/3", response3)
	r.HandleFunc("/4", response4)

	http.ListenAndServe(":8080", nil);
}
// Stores the jason data in set structs
type Noe1 struct {

	Entries []struct {

		Name   string `json:"navn"`
		Number string `json:"nummer"`

	} `json:"entries"`
}
type Noe2 struct {
	Entries []struct {
		Tvangsavvikling string `json:"tvangsavvikling"`
		Regnskap string `json:"regnskap"`
		Forradrpostnr string `json:"forradrpostnr"`
		Antall_ansatte string `json:"ansatte_antall"`
		Postadresse string `json:"postadresse"`

	} `json:"entries"`
}
type Noe3 struct {

	Entries []struct {

		Latitude string `json:"latitude"`
		Name string `json:"navn"`
		Plastic string `json:"plast"`
		GlasMetal string `json:"glass_metall"`
		Shoe string `json:"tekstil_sko,omitempty"`
		Longitude string `json:"longitude"`

	}`json:"entries"`
}
type Noe4 struct {

	Entries []struct {

		Name string `json:"name"`
		Shortname string `json:"shortNavn"`
		Groupable string `json:"groupable"`
		Searchable string `json:"searchable"`
		IndexPrimaryKey string `json:"indexPrimaryKey"`
		Description string `json:"description"`
		Definition string `json:"definition"`

	}`json:"entries"`
}
var noe1 Noe1
var noe2 Noe2
var noe3 Noe3
var noe4 Noe4

func response1(w http.ResponseWriter, r *http.Request) {

	url := "https://hotell.difi.no/api/json/difi/geo/fylke"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	json.Unmarshal(body, &noe1)

	t, err := template.ParseFiles("noe1.html")

	err = t.Execute(w, noe1)
	if err != nil {
		log.Fatal(err)
	}
}
func response2(w http.ResponseWriter, r *http.Request) {

	url := "https://hotell.difi.no/api/json/brreg/enhetsregisteret?page=8"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	json.Unmarshal(body, &noe2)

	t, err := template.ParseFiles("noe2.html")

	err = t.Execute(w, noe2)
	if err != nil {
		log.Fatal(err)
	}
}
	func response3(w http.ResponseWriter, r *http.Request) {

	url := "https://hotell.difi.no/api/json/stavanger/miljostasjoner"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	json.Unmarshal(body, &noe3)

	t, err := template.ParseFiles("noe3.html")

	err = t.Execute(w, noe3)
	if err != nil {
	log.Fatal(err)
	}
}
func response4(w http.ResponseWriter, r *http.Request) {

	url := "https://hotell.difi.no/api/json/difi/geo/fylke/fields"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	json.Unmarshal(body, &noe4)

	t, err := template.ParseFiles("noe4.html")

	err = t.Execute(w, noe4)
	if err != nil {
		log.Fatal(err)
	}
}
