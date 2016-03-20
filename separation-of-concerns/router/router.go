package router

import (
    "github.com/julienschmidt/httprouter"
    "nmct/soc/controllers"
    "net/http"
)

func NewRouter() *httprouter.Router {
    router := httprouter.New()

    for _, route := range routes {
        router.Handle(route.Method, route.Pattern, route.Handle.Handle)
    }

    // we have to wrap this in a http.HandlerFunc or it won't work
    router.NotFound = http.HandlerFunc(controllers.Serve404);

    return router
}
