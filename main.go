package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// a fun* fact about cats
// * = many facts are not fun
type CatFact struct {
	ID     string  `json:"id,omitempty"`
	Fact   string  `json:"fact,omitempty"`
	Source *Source `json:"source,omitempty"`
}

// where the hell did this fact come from?
type Source struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

// catfacts is a collection of CatFact
var catfacts []CatFact

// GetCatFactsEndpoint is used for getting a collection of catfacts
func GetCatFactEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	// looping catfacts
	for _, item := range catfacts {
		// if same id
		if item.ID == params["id"] {
			// return the catfact
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// empty catfact
	// json.NewEncode(w).Encode(&CatFact{})
}

// GetCatFactsEndpoint is used for getting a collection of catfacts
func GetCatFactsEndpoint(w http.ResponseWriter, req *http.Request) {
	// encodes and returns array of catfacts
	json.NewEncoder(w).Encode(catfacts)
}

// CreateCatFactEndpoint is used for creating a new catfact
// and adding it to results
func CreateCatFactEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var catfact CatFact
	_ = jsonNewDecoder(req.Body).Decode(&catfact)
	catfacts = append(catfacts, catfact)
	json.NewEncode(w).Encode(catfacts)
}

// DeleteCatFactEndpoint deletes a catfact
func DeleteCatFactEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	// looping catfacts
	for index, item := range catfacts {
		// if same id
		if item.ID == params["id"] {
			// delete the catfact from array (append everything but the catfact index)
			catfacts = append(catfacts[:index], catfacts[:index+1:]...)
      break
		}
	}
  json.NewEncode(w).Encode(catfacts)
}

func main() {
	router := mux.NewRouter()
	catfacts = append(catfacts, CatFact{ID: "1", Fact: "In 1987 cats overtook dogs as the number one pet in America.", Source: &Source{Name: "catoddities", Url: "http://www.catoddities.com/Trivia.html"}})
	catfacts = append(catfacts, CatFact{ID: "2", Fact: "About 37 percent of American homes today have at least one cat.", Source: &Source{Name: "catoddities", Url: "http://www.catoddities.com/Trivia.html"}})

	// returns several catfacts at once
	router.HandleFunc("/catfacts", GetCatFactsEndpoint).Methods("GET")
	// single return for catfact by ID
	router.HandleFunc("/catfact/{id}", GetCatFactEndpoint).Methods("GET")
	// creates a catfact
	router.HandleFunc("/catfact", CreateCatFactEndpoint).Methods("POST")
	// deletes a catfact
	router.HandleFunc("/catfact", DeleteCatFactEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":1337", router))
}
