package utils

import (
	"fmt"
	"strings"

	"data_scrapper/models"

	"github.com/gocolly/colly"
)

// GetProductData - used to scrap product data from amazon
func GetProductData(urlData *models.URLData) (*models.ProductData, error) {
	var name, price, imageURL, totalReviews, description string
	err := error(nil)
	c := colly.NewCollector()
	c.OnHTML("#productTitle", func(e *colly.HTMLElement) {
		name = strings.Trim(e.Text, "\n")
	})
	c.OnHTML("img#landingImage", func(e *colly.HTMLElement) {
		imageURL = e.Attr("data-old-hires")
	})
	c.OnHTML("#acrCustomerReviewText", func(e *colly.HTMLElement) {
		totalReviews = strings.Trim(e.Text, "\n")
	})
	c.OnHTML("#featurebullets_feature_div", func(e *colly.HTMLElement) {
		description = strings.Trim(e.Text, "\n")
	})
	c.OnHTML("span#priceblock_ourprice", func(e *colly.HTMLElement) {
		price = strings.Trim(e.Text, "\n")
	})
	c.OnHTML("span#priceblock_dealprice", func(e *colly.HTMLElement) {
		price = strings.Trim(e.Text, "\n")
	})
	c.OnError(func(r *colly.Response, err error) {
		err = fmt.Errorf("request url:", r.Request.URL, "failed with response:", r, "\nerror:", err)
	})
	c.Visit(urlData.URL)
	product := models.CreateProduct(urlData.URL, name, imageURL, description, price, totalReviews)
	return product, err
}
