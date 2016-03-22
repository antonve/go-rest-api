package controllers

import (
	"net/http"
	"log"
	"github.com/labstack/echo"
)

func Return500(c echo.Context, err string) error {
	log.Println(err)
	return Serve500(c)
}

func Serve500(c echo.Context) error {
	return c.JSONBlob(http.StatusInternalServerError, []byte(`{"success": false, "error": "500 internal server error"}`))
}

func Return404(c echo.Context, err string) error {
	log.Println(err)
	return Serve404(c)
}

func Serve404(c echo.Context) error {
	return c.JSONBlob(http.StatusNotFound, []byte(`{"success": false, "error": "404 page not found"}`))
}

func Serve405(c echo.Context) error {
	return c.JSONBlob(http.StatusMethodNotAllowed, []byte(`{"success": false, "error": "405 method not allowed"}`))
}

func Return201(c echo.Context) error {
	return c.JSONBlob(http.StatusCreated, []byte(`{"success": true}`))
}

func Return200(c echo.Context) error {
	return c.JSONBlob(http.StatusOK, []byte(`{"success": true}`))
}
