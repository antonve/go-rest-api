package main

import (
    "encoding/json"
    "net/http"
    "log"
    "io/ioutil"
    "path"
    "strconv"
    "fmt"
    "errors"
)

// DATA TYPES

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

// CONTROLLERS

func ServeHelp(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")
    http.ServeFile(w, r, path.Join("data", "help.json"))
}

func ServeProducts(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")

    switch r.Method {
        case "GET":
            http.ServeFile(w, r, path.Join("data", "products.json"))
        case "POST":
            products := GetProductsJson()
            product, err := ParseProduct(products.GetLastId()+1, r)
            if err != nil {
                w.Write([]byte(fmt.Sprintf(`{"success": false, "error": "Error adding product: %v"}`, err)))
                return
            }

            err = products.Add(product)
            if err == nil {
                w.Write([]byte(`{"success": true}`))
            } else {
                w.Write([]byte(fmt.Sprintf(`{"success": false, "error": "Error adding product: %v"}`, err)))
            }
        default:
            Return404(w, r, "Unknown method")
            return
    }
}

func ServeProductsWithId(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")

    id, err := strconv.Atoi(path.Base(r.URL.Path))
    if err != nil {
        Return404(w, r, "ID is not numeric")
        return
    }

    switch r.Method {
        case "GET":
            products := GetProductsJson()
            product := products.Get(id);

            if product != nil {
                b, err := json.Marshal(product)
                if err != nil {
                    Return404(w, r, fmt.Sprintf("Marshall of product failed: %v", err))
                    return
                }

                w.Write(b)
            }

        case "PUT":
            products := GetProductsJson()
            product, err := ParseProduct(id, r)
            if err != nil {
                w.Write([]byte(fmt.Sprintf(`{"success": false, "error": "Error updating product: %v"}`, err)))
                return
            }

            err = products.Update(product)
            if err == nil {
                w.Write([]byte(`{"success": true}`))
            } else {
                w.Write([]byte(fmt.Sprintf(`{"success": false, "error": "Error updating product: %v"}`, err)))
            }

        case "DELETE":
            products := GetProductsJson()
            err = products.Delete(id)

            if err == nil {
                w.Write([]byte(`{"success": true}`))
            } else {
                w.Write([]byte(fmt.Sprintf(`{"success": false, "error": "Error deleting product: %v"}`, err)))
            }

        default:
            Return404(w, r, "Unknown method")
            return
    }
}

// HELPER FUNCTIONS

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

func Return404(w http.ResponseWriter, r *http.Request, err string) {
    fmt.Println(err);
    w.Write([]byte(`{"error": "404 page not found"}`))
}

func ReadJson(filename string) []byte {
    file, err := ioutil.ReadFile(filename)

    if err != nil {
        log.Println("Error reading JSON file: %v", err)
    }

    return file
}

func GetProductsJson() Products {
    file := ReadJson(path.Join("data", "products.json"))
    var products Products
    
    err := json.Unmarshal(file, &products.Products)
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

// MAIN ROUTER

func main() {
    http.HandleFunc("/", ServeHelp)
    http.HandleFunc("/products", ServeProducts)
    http.HandleFunc("/products/", ServeProductsWithId)
    log.Fatal(http.ListenAndServe("localhost:8090", nil))
}