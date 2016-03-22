package models

import (
    "errors"
    "strconv"
    "github.com/labstack/echo"
)

func ParseProduct(c echo.Context) (*Product, error) {
    var err error
    product := new(Product)

    product.Name = c.FormValue("name")
    if product.Name == "" {
        return nil, errors.New("Name can't be empty")
    }

    product.Price, err = strconv.ParseFloat(c.FormValue("price"), 64)
    if err != nil {
        return nil, errors.New("Can't parse price to float")
    }

    product.Stock, err = strconv.Atoi(c.FormValue("stock"))
    if err != nil {
        return nil, errors.New("Can't parse stock to int")
    }

    return product, nil
}

