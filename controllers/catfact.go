package controllers

type Controller struct {
	Repository Repository
}

// Index - gets a collection of catfacts
// HTTP Request Type: GET
// Router URI: /catfacts
func (c *Controller) Index(w http.ReponseWriter, r *http.Request) {
	// list of all catfacts
	catfacts := c.Repository.GetCafacts()
}

// GetCatfacts - gets a single catfact by ID
// HTTP Request Type: GET
// Router URI: /catfact/{id}
func (c *Controller) GetCatfacts(w http.ReponseWriter, r *http.Request) {

}
