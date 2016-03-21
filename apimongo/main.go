package main

import (
    "net/http"
    "log"
    "os"
    "io"
    "fmt"
    "nmct/apimongo/router"
)

func main() {
    router := router.NewRouter()

    defer setupErrorLogging()()

    log.Fatal(http.ListenAndServe(":8090", router))
}

func setupErrorLogging() func() {
    logFile, err := os.OpenFile("error.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
    if err != nil {
        log.Panicln(err)
    }

    log.SetOutput(io.MultiWriter(os.Stderr, logFile))

    return func() {
        e := logFile.Close()
        if e != nil {
            fmt.Fprintf(os.Stderr, "Problem closing the log file: %s\n", e)
        }
    }
}
