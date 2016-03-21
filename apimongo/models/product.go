package models

import (
    "errors"
    "gopkg.in/mgo.v2/bson"
    "fmt"
)

type Products struct {
    Products []Product `json:"products"`
}

type Product struct {
    Id          bson.ObjectId   `json:"id" bson:"_id,omitempty"`
    Name        string          `json:"name" bson:"name"`
    Price       float64         `json:"price" bson:"price"`
    Stock       int             `json:"stock" bson:"stock"`
}

func (products *Products) Length() int {
    return len(products.Products);
}

func (products *Products) GetAll() error {
    ds := NewDataStore()
    defer ds.Close()

    c := ds.C("products")
    err := c.Find(nil).All(&products.Products)

    return err
}

func (products *Products) Get(id bson.ObjectId) (*Product, error) {
    for _, product := range products.Products {
        if (product.Id == id) {
            return &product, nil
        }
    }

    ds := NewDataStore()
    defer ds.Close()

    c := ds.C("products")
    product := Product{}
    err := c.Find(bson.M{"_id": id}).One(&product)

    if err != nil {
        return nil, errors.New(fmt.Sprintf("Couldn't find product with ID '%v'", id))
    }

    products.Products = append(products.Products, product)

    return &product, nil
}

func (products *Products) Update(product *Product) error {
    ds := NewDataStore()
    defer ds.Close()

    c := ds.C("products")
    err := c.Update(bson.M{"_id": product.Id}, product)

    if err != nil {
        return err
    }

    for index, product := range products.Products {
        if (product.Id == products.Products[index].Id) {
            products.Products[index] = product
        }
    }

    return nil
}

func (products *Products) Delete(id bson.ObjectId) error {

    ds := NewDataStore()
    defer ds.Close()

    c := ds.C("products")
    err := c.RemoveId(id)

    if err != nil {
        return err
    }

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
    }

    return nil
}

func (products *Products) Add(product *Product) error {
    ds := NewDataStore()
    defer ds.Close()

    c := ds.C("products")

    err := c.Insert(&product)
    if err != nil {
        return errors.New("Cound't add product to database")
    }

    products.Products = append(products.Products, *product)

    return nil
}
