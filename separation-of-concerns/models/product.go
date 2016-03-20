package models

import (
    "errors"
    "fmt"
    "net/http"
    "strconv"
    "path"
    "encoding/json"
    "log"
    "io/ioutil"
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


// Products datastorage helpers

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
