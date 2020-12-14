package handlers

import (
	"io/ioutil"
	"log"
	"net/http"

	"bytes"

	"data_scrapper/models"
	"data_scrapper/utils"
)

// Scrap - struct to handle logging
type Scrap struct {
	l *log.Logger
}

// NewScrap - Creates new Scrap struct object
func NewScrap(l *log.Logger) *Scrap {
	return &Scrap{l}
}

// ServerHTTP - Handles HTTP request to /scrap
func (s *Scrap) ServerHTTP(rw http.ResponseWriter, r *http.Request) {
	s.l.Println("Scrapping data from given URL")
	urlData := &models.URLData{}
	err := urlData.FromJSON(r.Body)
	if err != nil {
		s.l.Println(err)
		http.Error(rw, "Unabled to parse input data", http.StatusBadRequest)
	}
	s.l.Printf("URL :- %#v", urlData)
	product, err := utils.GetProductData(urlData)
	if err != nil {
		http.Error(rw, "Unabled to get data from amazon", http.StatusNotFound)
	}
	s.l.Printf("Product :- %#v", product)
	productByteBuffer := new(bytes.Buffer)
	err = product.ToJSON(productByteBuffer)
	if err != nil {
		s.l.Fatal(err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
	}
	res, err := http.Post("http://add_data:9091/add", "application/json", productByteBuffer)
	if err != nil {
		s.l.Fatal(err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		s.l.Fatal(err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
	}
	s.l.Printf(string(bodyBytes))
	rw.Write(bodyBytes)
}
