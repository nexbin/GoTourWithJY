package main

import (
	"log"
	"timeTool/root"
)

func main() {
	err := root.Execute()
	if err != nil {
		log.Fatal(err)
		return
	}
}
