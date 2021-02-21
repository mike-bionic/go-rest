package main

import "os"

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("GOAPP_DB_USERNAME")
		os.Getenv("GOAPP_DB_PASSWORD")
		os.Getenv("GOAPP_DB_NAME")
	)

	a.Run(":8010")
}