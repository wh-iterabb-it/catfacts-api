// model.go
package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
)

type (
	Models interface {
		NewModels() *models

		getCatFact(db *sql.DB) error
	}

	CatFact struct {
		ID     string  `json:"id,omitempty"`
		Fact   string  `json:"fact,omitempty"`
		Source *Source `json:"source,omitempty"`
	}

	Source struct {
		Name string `json:"name,omitempty"`
		Url  string `json:"url,omitempty"`
	}
)

func NewModels() *models {

	// catfacts is a collection of CatFact
	var catfacts []CatFact

	m := new(models)
	return m
}

func (u *CatFact) getCatFact(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *CatFact) updateCatFact(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *CatFact) deleteCatFact(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *CatFact) createCatFact(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getCatFacts(db *sql.DB, start int, count int, w http.ResponseWriter) {
	// encodes and returns array of catfacts
	result, err := json.NewEncoder(w).Encode(catfacts)

	if err != nil {
		return err
	}

	return result
}
