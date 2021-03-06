package handlers

import (
	"net/http"

	"github.com/TamerB/go-microservices/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//   200: productsResponse
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
