package main

import (
	"echo-rest/db"
	"echo-rest/routes"
)

// main akan mendefinisikan servernya
func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start("localhost:8888"))
}

