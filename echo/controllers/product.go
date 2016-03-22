package controllers

import (
    "net/http"
    "nmct/echo/models"
    "gopkg.in/mgo.v2/bson"
    "github.com/labstack/echo"
)

func APIProductsGet(c echo.Context) error {
    products := models.Products{}
    err := products.GetAll();

    if err != nil {
        return Return500(c, err.Error())
    }

    return c.JSON(http.StatusOK, products.Products)
}

func APIProductsPost(c echo.Context) error {
    product, err := models.ParseProduct(c)
    if err != nil {
        return Return500(c, err.Error())
    }

    products := models.Products{}
    err = products.Add(product)
    if err != nil {
        return Return500(c, err.Error())
    }

    return Return201(c)
}

func APIProductGet(c echo.Context) error {
    id := bson.ObjectIdHex(c.Param("id"))

    products := models.Products{}
    product, err := products.Get(id);

    if err != nil {
        return Return404(c, err.Error())
    }

    return c.JSON(http.StatusOK, product)
}

func APIProductPut(c echo.Context) error {
    product, err := models.ParseProduct(c)

    if err != nil {
        return Return500(c, err.Error())
    }

    products := models.Products{}
    product.Id = bson.ObjectIdHex(c.Param("id"));
    err = products.Update(product)
    if err != nil {
        return Return500(c, err.Error())
    }

    return Return201(c)
}

func APIProductDelete(c echo.Context) error {
    products := models.Products{}
    id := bson.ObjectIdHex(c.Param("id"));
    err := products.Delete(id)
    if err != nil {
        return Return500(c, err.Error())
    }

    return Return200(c)
}
