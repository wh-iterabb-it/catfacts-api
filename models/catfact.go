// model.go
package main

import (
	"database/sql"
	"errors"
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

func getCatFacts(db *sql.DB, start, count int) ([]CatFact, error) {
	return nil, errors.New("Not implemented")
}
