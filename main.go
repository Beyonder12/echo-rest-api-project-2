package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("MY_APP_PORT")

	if port == "" {
		port = "8181"
	}

	e := echo.New()
	products := []map[int]string{{1: "mobiles"}, {2: "tv"}, {3: "laptops"}}
	e.GET("/products/:id", func(c echo.Context) error {
		var product map[int]string
		for _, p := range products {
			for k := range p {
				pID, err := strconv.Atoi(c.Param("id"))
				if err != nil {
					return err
				}
				if pID == k {
					product = p
				}
			}
		}
		if product == nil {
			return c.JSON(http.StatusOK, "Product not found")
		}
		return c.JSON(http.StatusOK, product)
	})
	e.Logger.Print(fmt.Sprintf("Listening on port %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
