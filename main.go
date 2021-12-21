package main

import "echo-rest/routes"

// main akan mendefinisikan servernya
func main() {
	e := routes.Init()

	e.Logger.Fatal(e.Start("localhost:8888"))
}