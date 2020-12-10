package utils

import(
	"../models"
	"fmt"
	"strings"
	"github.com/gocolly/colly"
)

func GetProductData(urlData *models.URLData) (*models.ProductData, error) {
	var name, price, imageURL, totalReviews, description string
	err := error(nil)
	fmt.Println("here")
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
