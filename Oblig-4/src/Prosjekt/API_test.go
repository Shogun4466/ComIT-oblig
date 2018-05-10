package main

import (
	"net/http"
	"encoding/json"
	"log"
	"testing"
)

//Returnerer dataene i APIet
func ResponseTest(url string) Test {
	resp, err := http.Get("https://hotell.difi.no/api/json/stavanger/miljostasjoner")
	if err != nil {
		log.Fatal(err)
	}
  
	defer resp.Body.Close()
  
	var data Test
  
	err = json.NewDecoder(resp.Body).Decode(&data)
  
	return data
}

//Test for å se om responsen er som forventet (fungerer så lenge det ikke blir lagt til flere stasjoner)
func TestCorrectResponsese(t *testing.T) {
	r := ResponseTest("https://hotell.difi.no/api/json/stavanger/miljostasjoner")
	if r.Posts != 56 {
		t.Errorf("APIet returnerte posts: %d, forventet 56", r.Posts)
	}
}

//Test for å se om responsen er noe annet enn forventet
func TestWrongResponsese(t *testing.T) {
	r := ResponseTest("https://hotell.difi.no/api/json/stavanger/miljostasjoner")
	if r.Posts == 56 {
		t.Errorf("APIet returnerte posts: 56 som forventet")
	}
}
