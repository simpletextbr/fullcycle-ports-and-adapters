package handler

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/simpletextbr/fullcycle-ports-and-adapters/adapters/dto"
	"github.com/simpletextbr/fullcycle-ports-and-adapters/application"
)

func MakeProductHandler(r *mux.Router, n *negroni.Negroni, service application.IProductService) {

	r.Handle("/products/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/product", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")

}

func getProduct(service application.IProductService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]

		product, err := service.GetByID(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	})
}

func createProduct(service application.IProductService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var productDto dto.Product

		err := json.NewDecoder(r.Body).Decode(&productDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(JsonError(err.Error()))
			return
		}

		prodcut, err := service.Create(productDto.Name, productDto.Price)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(JsonError(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(prodcut)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(JsonError(err.Error()))
			return
		}
	})
}
