package controllers

import (
	"log"

  "github.com/wh-iterabb-it/catfacts-api/db"
	"github.com/gorilla/mux"
)

type Controller struct {
	db.Repository db.Repository
}

// Index - gets a collection of catfacts
// HTTP Request Type: GET
// Router URI: /catfacts
func (c *Controller) Index(w http.ReponseWriter, r *http.Request) {
	// list of all catfacts
	catfacts := c.Repository.GetCafact()
}

// GetCatfacts - gets a single catfact by ID
// HTTP Request Type: GET
// Router URI: /catfact/{id}
func (c *Controller) GetCatfact(w http.ReponseWriter, r *http.Request) {
	// parsing vars
	vars := mux.Vars(r)
	log.Println("GetCatfact (vars): ", vars)

	id := vars["id"]
	log.Println("GetCatfact (id): ", id)

	// catfactid, err := (something that will be done with a db handler)
	// if err != nil {
	//   log.Fatalln("Error GetCatfacts", err)
	// }
}
