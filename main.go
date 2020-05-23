package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("hoge")
	fmt.Fprintf(writer, "Hello World, %s!!", request.URL.Path[1:])
}

func main() {
	//http.HandleFunc("/", handler)
	//http.ListenAndServe(":8888", nil)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8888"))
}
