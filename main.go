package main

import "github.com/Lynczera/goS/app"

func main() {
	err := app.Setup()
	if err != nil {
		panic(err)
	}

}
