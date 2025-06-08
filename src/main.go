package main

import (
	"net/http"

	"aplicacaoweb/src/routes"
)

func main() {
	routes.CarregarRotas()
	http.ListenAndServe(":8000", nil)
}
