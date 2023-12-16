package main

import "log"

func main() {
	err := FetchAndRegisterMcdonaldsMenu()
	if err != nil {
		log.Fatal(err)
	}
}
