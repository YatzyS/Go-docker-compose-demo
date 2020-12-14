package handlers

import (
	"fmt"
	"log"
	"net/http"

	"data_add/models"
	"data_add/utils"
)

type Add struct {
	l *log.Logger
}

// NewAdd - creates and return pointer struct Add
func NewAdd(l *log.Logger) *Add {
	return &Add{l}
}

// ServerHTTP - handles HTTP request to /add
func (a *Add) ServerHTTP(rw http.ResponseWriter, r *http.Request) {
	a.l.Println("Adding data from given URL to JSON file")
	productData := &models.ProductData{}
	err := productData.FromJSON(r.Body)
	if err != nil {
		a.l.Println(err)
		http.Error(rw, "Unabled to parse input data", http.StatusBadRequest)
	}
	a.l.Printf("ProductData :- %#v", productData)
	err = utils.AddToJSON(productData)
	if err != nil {
		a.l.Println(err)
		http.Error(rw, "Server Side Error", http.StatusInternalServerError)
	}
	a.l.Printf("Data Inserted")
	rw.Write([]byte(fmt.Sprintf("Data inserted")))
}
