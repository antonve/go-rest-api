package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// This is our middleware for the httprouter
// Here we can manipulate the request for all requests

type Action struct {
	Handler  httprouter.Handle
}

func (a *Action) Handle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-type", "application/json")

	a.Handler(w, r, ps)
}
