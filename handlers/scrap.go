package handlers

import (
	"log"
	"net/http"
	"io/ioutil"

	"../models"
	"../utils"
	"bytes"
)

type Scrap struct {
	l *log.Logger
}

func NewScrap(l *log.Logger) *Scrap {
	return &Scrap{l}
}

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
	res, err := http.Post("http://localhost:9090/add", "application/json", productByteBuffer)
	if err != nil {
		s.l.Fatal(err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		s.l.Fatal(err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
	}
	rw.Write(bodyBytes)
}
