package models

import (
	"database/sql"
	"fmt"
)

type (
	// CatFact represents a singular cat fun* fact
	// * not all facts are fun
	CatFact struct {
		ID          int    `json:"id"`
		Fact        string `json:"fact,omitempty"`
		Source_Name string `json:"source_name,omitempty"`
		Source_Url  string `json:"source_url,omitempty"`
	}
)

func (c *CatFact) getCatFact(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT id, fact, source_name, source_url FROM catfact WHERE id=%d", c.ID)
	return db.QueryRow(statement).Scan(&c.ID, &c.Fact, &c.Source_Name, &c.Source_Url)
}

func (c *CatFact) updateCatFact(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE catfact SET fact='%s', source_name='%s', source_url='%s' WHERE id=%d", c.Fact, c.Source_Name, c.Source_Url, c.ID)
	_, err := db.Exec(statement)
	return err
}
func (c *CatFact) deleteCatFact(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM catfact WHERE id=%d", c.ID)
	_, err := db.Exec(statement)
	return err
}
func (c *CatFact) createCatFact(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO catfact(fact, source_name, source_url) VALUES('%s','%s','%s')", c.Fact, c.Source_Name, c.Source_Url)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&c.ID)
	if err != nil {
		return err
	}
	return nil
}

func getCatFacts(db *sql.DB, start int, count int) ([]CatFact, error) {
	statement := fmt.Sprintf("SELECT id, fact, source_name, source_url FROM catfact LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	catfacts := []CatFact{}
	for rows.Next() {
		var c CatFact
		if err := rows.Scan(&c.ID, &c.Fact, &c.Source_Name, &c.Source_Url); err != nil {
			return nil, err
		}
		catfacts = append(catfacts, c)
	}
	return catfacts, nil
}
