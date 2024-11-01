package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, _ := http.Get("https://pokeapi.co/api/v2/pokedex/kanto/")

	responseDta, _ := io.ReadAll(response.Body)
	fmt.Println(string(responseDta))
}
