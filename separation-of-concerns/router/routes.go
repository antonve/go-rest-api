package router

import (
    "nmct/soc/controllers"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    Handle      *Action
}

type Routes []Route

var routes = Routes{
    Route{
        "API help",
        "GET",
        "/",
        &Action{controllers.APIHelp},
    },
    Route{
        "Get products",
        "GET",
        "/products",
        &Action{controllers.APIProductsGet},
    },
    Route{
        "Add new product",
        "POST",
        "/products",
        &Action{controllers.APIProductsPost},
    },
    Route{
        "Get product details",
        "GET",
        "/products/:id",
        &Action{controllers.APIProductGet},
    },
    Route{
        "Update product",
        "PUT",
        "/products/:id",
        &Action{controllers.APIProductPut},
    },
    Route{
        "Delete product",
        "DELETE",
        "/products/:id",
        &Action{controllers.APIProductDelete},
    },
}