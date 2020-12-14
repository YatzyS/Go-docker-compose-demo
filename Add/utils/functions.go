package utils

import (
	"bytes"
	"fmt"
	"os"

	"data_add/models"
)

// AddToJSON - used to add product data to a JSON file
func AddToJSON(product *models.ProductData) error {
	file, err := os.OpenFile(JSON_FILE_NAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Couldn't create or append to existing file ", err)
	}
	defer file.Close()
	productBytesBuffer := new(bytes.Buffer)
	err = product.ToJSON(productBytesBuffer)
	if err != nil {
		return fmt.Errorf("Couldn't Unmarshal ProductData struct ", err)
	}
	_, err = file.Write(productBytesBuffer.Bytes())
	if err != nil {
		return fmt.Errorf("Couldn't write JSON to file ", err)
	}
	return nil
}
