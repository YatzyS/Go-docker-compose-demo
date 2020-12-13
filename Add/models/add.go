package models

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type ProductData struct {
	URL       string  `json: "url"`
	Product   Product `json: "product"`
	TimeStamp string  `json: "timestamp"`
}

type Product struct {
	Name         string `json: "name"`
	ImageURL     string `json: "imageURL"`
	Description  string `json: "description"`
	Price        string `json: "price"`
	TotalReviews string `json: "totalReviews"`
}

func CreateProduct(url, name, imageURL, description, price, totalReviews string) *ProductData {
	return &ProductData{
		URL: url,
		Product: Product{
			Name:         name,
			ImageURL:     imageURL,
			Description:  description,
			Price:        price,
			TotalReviews: totalReviews,
		},
		TimeStamp: time.Now().String(),
	}
}

func (p *ProductData) String() string {
	return fmt.Sprintf("URL: %s \n Product Name:- %s \n ImageURL: %s \n Description %s \n Price %s \n TotalReviews: %s", p.URL, p.Product.Name, p.Product.ImageURL, p.Product.Description, p.Product.Price, p.Product.TotalReviews)
}

func (p *ProductData) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *ProductData) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
