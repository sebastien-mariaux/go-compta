package main

import "go-compta/api"

func main() {
	api.SetupRouter().Run("localhost:8080")
}
