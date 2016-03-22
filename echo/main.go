package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
    "github.com/labstack/echo/middleware"

    "nmct/echo/controllers"
    "os"
    "io"
    "fmt"
    "log"
)

func main() {
    // Echo instance
    e := echo.New()

    // Middleware
    e.Use(middleware.Recover())
    defer setupErrorLogging(e)()

    // Routes
    e.Static("/", "data/help.json")

    g := e.Group("/products")
    g.Use(middleware.BasicAuth(func(username, password string) bool {
        if username == "foo" && password == "bar" {
            return true
        }
        return false
    }))

    g.Get("", echo.HandlerFunc(controllers.APIProductsGet))
    g.Post("", echo.HandlerFunc(controllers.APIProductsPost))
    g.Get("/:id", echo.HandlerFunc(controllers.APIProductGet))
    g.Put("/:id", echo.HandlerFunc(controllers.APIProductPut))
    g.Delete("/:id", echo.HandlerFunc(controllers.APIProductDelete))

    // Start server
    e.Run(standard.New(":1323"))
}

func setupErrorLogging(e *echo.Echo) func() {
    logFile, err := os.OpenFile("error.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
    if err != nil {
        log.Panicln(err)
    }

    e.Use(middleware.LoggerFromConfig(middleware.LoggerConfig{
        Format: "time=${time_rfc3339}, remote_ip=${remote_ip}, method=${method}, " +
        "path=${path}, status=${status}, took=${response_time}, sent=t=${response_size} bytes\n",
        Output: io.MultiWriter(os.Stderr, logFile),
    }))

    return func() {
        e := logFile.Close()
        if e != nil {
            fmt.Fprintf(os.Stderr, "Problem closing the log file: %s\n", e)
        }
    }
}
