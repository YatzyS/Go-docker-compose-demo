package utils

import(
	"../models"
	"fmt"
	"strings"
	"github.com/gocolly/colly"
	"os"
	"bytes"
)

func GetProductData(urlData *models.URLData) (*models.ProductData, error) {
	var name, price, imageURL, totalReviews, description string
	err := error(nil)
	c := colly.NewCollector()
	c.OnHTML("#productTitle", func(e *colly.HTMLElement) {
		name = strings.Trim(e.Text,"\n")
	})
	c.OnHTML("#landingImage", func(e *colly.HTMLElement) {
		imageURL = e.Attr("src")
	})
	c.OnHTML("#acrCustomerReviewText", func(e *colly.HTMLElement) {
		totalReviews = strings.Trim(e.Text,"\n")
	})
	c.OnHTML("#featurebullets_feature_div", func(e *colly.HTMLElement) {
		description = strings.Trim(e.Text,"\n")
	})
	c.OnHTML("div#newAccordionRow span.a-color-price", func(e *colly.HTMLElement) {
		price = strings.Trim(e.Text,"\n")
	})
	c.OnError(func(r *colly.Response, err error) {
		err = fmt.Errorf("request url:", r.Request.URL, "failed with response:", r, "\nerror:", err)
	})
	c.Visit(urlData.URL)
	product := models.CreateProduct(urlData.URL, name, imageURL, description, price, totalReviews)
	return product, err
}

func AddToJSON(product *models.ProductData) error {
	file, err := os.OpenFile(JSON_FILE_NAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Couldn't create or append to existing file ",err)
	}
	defer file.Close()
	productBytesBuffer := new(bytes.Buffer)
	err = product.ToJSON(productBytesBuffer)
	if err != nil {
		return fmt.Errorf("Couldn't Unmarshal ProductData struct ",err)
	}
	_,err = file.Write(productBytesBuffer.Bytes())
	if err != nil {
		return fmt.Errorf("Couldn't write JSON to file ",err)
	}
	return nil
}

