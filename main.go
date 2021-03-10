package main

import (
	"go-gin-api-rest/controllers"
)

func main() {
	r := controllers.HandleRequests()

	r.Run()
}
