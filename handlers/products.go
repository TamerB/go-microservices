package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"../data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.getProducts(w, r)
		return
	case http.MethodPost:
		p.addProduct(w, r)
		return
	case http.MethodPut:
		rgx := regexp.MustCompile(`/([0-9]+)`)
		g := rgx.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			p.l.Println("Invalid URI, more than one id")
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid URI, more than one capture group")
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			p.l.Println("Invalid URI, unable to conver to number", idString)
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.l.Println("got id", id)
		p.updateProduct(id, w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}

	p.l.Printf("Product: %#v", prod)

	data.AddProduct(prod)
}

func (p *Products) updateProduct(id int, w http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}

	p.l.Printf("Product: %#v", prod)

	err1 := data.UpdateProduct(id, prod)

	if err1 == data.ErrorProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err1 != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}
