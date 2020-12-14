package models

import(
	"fmt"
	"encoding/json"
	"io"
	"time"
)

// URLData - To hold URL data from request
type URLData struct {
	URL string `json: "url"`
}

// String - Used to get struct value in string format
func (u *URLData) String() string {
	return fmt.Sprintf("URL: %s", u.URL)
}

// FromJSON - used to decode JSON to struct
func (u *URLData) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

// ToJSON - used to encode struct to JSON
func (u *URLData) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

// ProductData - used to hold data realted to product and timestamp
type ProductData struct {
	URL string `json: "url"`
	Product Product `json: "product"`
	TimeStamp string `json: "timestamp"`
}

// Product - used to hold product
type Product struct {
	Name string `json: "name"`
	ImageURL string `json: "imageURL"`
	Description string `json: "description"`
	Price string `json: "price"`
	TotalReviews string `json: "totalReviews"`
}

// CreateProduct -Creates a new Product with timestamp
func CreateProduct(url, name, imageURL, description, price, totalReviews string) *ProductData {
	return &ProductData{
		URL:url,
		Product: Product{
			Name: name,
			ImageURL: imageURL,
			Description: description,
			Price: price,
			TotalReviews: totalReviews,
		},
		TimeStamp: time.Now().String(),
	}
} 

// String - Used to get struct value in string format
func (p *ProductData) String() string {
	return fmt.Sprintf("URL: %s \n Product Name:- %s \n ImageURL: %s \n Description %s \n Price %s \n TotalReviews: %s", p.URL, p.Product.Name, p.Product.ImageURL, p.Product.Description, p.Product.Price, p.Product.TotalReviews)
}

// FromJSON - used to decode JSON to struct
func (p *ProductData) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// ToJSON - used to encode struct to JSON
func (p *ProductData) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}