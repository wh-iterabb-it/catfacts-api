package models

type (
	Models interface {
		NewCatFacts() *models
	}

	// CatFact represents a singular cat fun* fact
	// * not all facts are fun
	CatFact struct {
		ID     string  `json:"id,omitempty"`
		Fact   string  `json:"fact,omitempty"`
		Source *Source `json:"source,omitempty"`
	}

	// Source represents who or what made the CatFact item
	Source struct {
		Name string `json:"name,omitempty"`
		Url  string `json:"url,omitempty"`
	}

	// CatFacts is an array of CatFact objects
	CatFacts []CatFact
)

func NewCatFacts() *models {
	m := new(models)
	return m
}
