package main

import (
	"fmt"
	"log"
	"os"
	"searchipgo/app"
)

func main() {
	fmt.Println("SearchipGO")

	application := app.Generate()
	if error := application.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
