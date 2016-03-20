package controllers

import (
    "net/http"
    "path"
    "fmt"
    "strconv"
    "encoding/json"
    "github.com/julienschmidt/httprouter"
    "nmct/soc/models"
    "log"
)

func APIHelp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    http.ServeFile(w, r, path.Join("data", "help.json"))
}

func APIProductsGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    http.ServeFile(w, r, path.Join("data", "products.json"))
}

func APIProductsPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    products := models.GetProducts()
    product, err := models.ParseProduct(products.GetLastId()+1, r)
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
}

func APIProductGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id, err := strconv.Atoi(ps.ByName("id"))
    if err != nil {
        Return404(w, r, "ID is not numeric")
        return
    }

    products := models.GetProducts()
    product := products.Get(id);

    if product == nil {
        Return404(w, r, fmt.Sprintf("Product with ID '%v' could not be found", id))
        return
    }

    b, err := json.Marshal(product)
    if err != nil {
        Return404(w, r, fmt.Sprintf("Marshall of product failed: %v", err))
        return
    }

    w.Write(b)

}

func APIProductPut(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id, err := strconv.Atoi(ps.ByName("id"))
    if err != nil {
        Return404(w, r, "ID is not numeric")
        return
    }

    products := models.GetProducts()
    product, err := models.ParseProduct(id, r)
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
}

func APIProductDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id, err := strconv.Atoi(ps.ByName("id"))
    if err != nil {
        Return404(w, r, "ID is not numeric")
        return
    }

    products := models.GetProducts()
    err = products.Delete(id)

    if err == nil {
        w.Write([]byte(`{"success": true}`))
    } else {
        w.Write([]byte(fmt.Sprintf(`{"success": false, "error": "Error deleting product: %v"}`, err)))
    }
}

func Return404(w http.ResponseWriter, r *http.Request, err string) {
    log.Println(err)
    Serve404(w, r)
}

func Serve404(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")
    w.Write([]byte(`{"error": "404 page not found"}`))
}
