package models

import (
	"io/ioutil"
	"path"
	"log"
	"encoding/json"
)

func GetProductsJson() Products {
	file, err := ioutil.ReadFile(path.Join("data", "products.json"))

	if err != nil {
		log.Println("Error reading JSON file: %v", err)
	}

	var products Products

	err = json.Unmarshal(file, &products.Products)
	if err != nil {
		log.Println("Error with JSON products: %v", err)
	}

	return products
}

func SaveProducts(products *Products) error {
	b, err := json.MarshalIndent(products.Products, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path.Join("data", "products.json"), b, 0644)
	if err != nil {
		return err
	}

	return nil
}
