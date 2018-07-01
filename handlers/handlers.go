package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	. "github.com/wh-iterabb-it/catfacts-api/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Handler struct {
	Router *mux.Router
	DB     *sql.DB
}

func (h *Handler) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)

	var err error
	h.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	h.Router = mux.NewRouter()
	h.initializeRoutes()
}

func (h *Handler) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, h.Router))
}

func (h *Handler) initializeRoutes() {
	h.Router.HandleFunc("/catfacts", h.getCatfacts).Methods("GET")
	h.Router.HandleFunc("/catfact", h.createCatfact).Methods("POST")
	h.Router.HandleFunc("/catfact/{id:[0-9]+}", h.getCatfact).Methods("GET")
	h.Router.HandleFunc("/catfact/{id:[0-9]+}", h.updateCatfact).Methods("PUT")
	h.Router.HandleFunc("/catfact/{id:[0-9]+}", h.deleteCatfact).Methods("DELETE")
}

// GetCatfacts - gets a single catfact by ID
// HTTP Request Type: GET
// Router URI: /catfact/{id}
func (h *Handler) getCatfact(w http.ResponseWriter, r *http.Request) {
	// parsing vars
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid catfact ID")
		return
	}

	c := CatFact{ID: id}
	if err := c.getCatFact(h.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "catfact not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, c)
}

// GetCatfacts - gets a collection of catfacts
// HTTP Request Type: GET
// Router URI: /catfacts
func (h *Handler) getCatfacts(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	products, err := getCatFacts(h.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

// updateCatfact - updates a single catfact by ID
// HTTP Request Type: PUT
// Router URI: /catfact/{id}
func (h *Handler) updateCatfact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid catfact ID")
		return
	}

	var c CatFact
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	c.ID = id

	if err := c.updateCatFact(h.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, c)
}

// deleteCatfact - deletes a single catfact by ID
// HTTP Request Type: DELETE
// Router URI: /catfact/{id}
func (h *Handler) deleteCatfact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Catfact ID")
		return
	}

	c := CatFact{ID: id}
	if err := c.deleteCatFact(h.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// createCatfact - Create
// HTTP Request Type: POST
// Router URI: /catfact
func (h *Handler) createCatfact(w http.ResponseWriter, r *http.Request) {
	var c CatFact
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := c.createCatFact(h.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, c)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
