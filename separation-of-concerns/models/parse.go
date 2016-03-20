package models

import (
	"net/http"
	"errors"
	"strconv"
)

func ParseProduct(id int, r *http.Request) (*Product, error) {
	r.ParseForm()
	var err error
	product := new(Product)
	product.Id = id;

	product.Name = r.FormValue("name")
	if product.Name == "" {
		return nil, errors.New("Name can't be empty")
	}

	product.Price, err = strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		return nil, errors.New("Can't parse price to float")
	}

	product.Stock, err = strconv.Atoi(r.FormValue("stock"))
	if err != nil {
		return nil, errors.New("Can't parse stock to int")
	}

	return product, nil
}

