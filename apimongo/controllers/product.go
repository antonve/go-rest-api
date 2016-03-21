package controllers

import (
    "net/http"
    "path"
    "fmt"
    "encoding/json"
    "github.com/julienschmidt/httprouter"
    "nmct/apimongo/models"
    "gopkg.in/mgo.v2/bson"
)

func APIHelp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    http.ServeFile(w, r, path.Join("data", "help.json"))
}

func APIProductsGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    products := models.Products{}
    err := products.GetAll();

    if err != nil {
        Return500(w, r, err.Error())
        return
    }

    b, err := json.Marshal(products.Products)
    if err != nil {
        Return500(w, r, fmt.Sprintf("Marshall of product failed: %v", err))
        return
    }

    w.Write(b)
}

func APIProductsPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    product, err := models.ParseProduct(r)
    if err != nil {
        Return500(w, r, err.Error())
        return
    }

    products := models.Products{}
    err = products.Add(product)
    if err != nil {
        Return500(w, r, err.Error())
        return
    }

    Return201(w, r)
}

func APIProductGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id := bson.ObjectIdHex(ps.ByName("id"))

    products := models.Products{}
    product, err := products.Get(id);

    if err != nil {
        Return404(w, r, err.Error())
        return
    }

    b, err := json.Marshal(product)
    if err != nil {
        Return500(w, r, fmt.Sprintf("Marshall of product failed: %v", err))
        return
    }

    w.Write(b)
}

func APIProductPut(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    product, err := models.ParseProduct(r)

    if err != nil {
        Return500(w, r, err.Error())
        return
    }

    products := models.Products{}
    product.Id = bson.ObjectIdHex(ps.ByName("id"));
    err = products.Update(product)
    if err != nil {
        Return500(w, r, err.Error())
        return
    }

    Return201(w, r)
}

func APIProductDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    products := models.Products{}
    id := bson.ObjectIdHex(ps.ByName("id"));
    err := products.Delete(id)
    if err != nil {
        Return500(w, r, err.Error())
        return
    }

    Return200(w, r)
}
