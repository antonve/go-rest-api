package models

import (
    "errors"
    "fmt"
)

type Products struct {
    Products []Product `json:"products"`
}

type Product struct {
    Id          int         `json:"id"`
    Name        string      `json:"name"`
    Price       float64     `json:"price"`
    Stock       int         `json:"stock"`
}

func (products *Products) Length() int {
    return len(products.Products);
}

func (products *Products) GetLastId() int {
    return products.Products[products.Length() - 1].Id;
}

func (products *Products) Get(id int) *Product {
    for _, product := range products.Products {
        if (product.Id == id) {
            return &product
        }
    }

    return nil
}

func (products *Products) Update(product *Product) error {
    for index, el := range products.Products {
        if (el.Id == product.Id) {
            products.Products[index] = *product
            err := SaveProducts(products)
            return err
        }
    }

    return errors.New("Couldn't find product to update")
}

func (products *Products) Delete(id int) error {
    indexToDelete := 0;
    for index, el := range products.Products {
        if (el.Id == id) {
            indexToDelete = index
            break;
        }
    }
    
    if indexToDelete != 0 {
        // used to delete an item in a slice reference: https://github.com/golang/go/wiki/SliceTricks
        products.Products = products.Products[:indexToDelete+copy(products.Products[indexToDelete:], products.Products[indexToDelete+1:])]
        err := SaveProducts(products)

        return err
    }

    return errors.New("Couldn't find product to delete")
}

func (products *Products) Add(product *Product) error {
    for _, el := range products.Products {
        if (el.Id == product.Id) {
            return errors.New(fmt.Sprintf("Product with id %v already exists", product.Id))
        }
    }

    products.Products = append(products.Products, *product)
    return SaveProducts(products)
}
