package controllers

import (
	"net/http"
	"log"
)

func Return500(w http.ResponseWriter, r *http.Request, err string) {
	log.Println(err)
	Serve500(w, r)
}

func Serve500(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)

	w.Write([]byte(`{"success": false, "error": "500 internal server error"}`))
}

func Return404(w http.ResponseWriter, r *http.Request, err string) {
	log.Println(err)
	Serve404(w, r)
}

func Serve404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	w.Write([]byte(`{"success": false, "error": "404 page not found"}`))
}

func Serve405(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)

	w.Write([]byte(`{"success": false, "error": "405 method not allowed"}`))
}

func Return201(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte(`{"success": true}`))
}

func Return200(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(`{"success": true}`))
}
