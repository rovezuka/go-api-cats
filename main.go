package main

import (
	"fmt"
	"github.com/rovezuka/go-api-cats/api"
	"log"
)

func main() {
	apiClient, err := api.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	catImages, err := apiClient.GetAsset("10")
	if err != nil {
		log.Fatal(err)
	}

	for _, catImage := range catImages {
		fmt.Println(catImage.Info())
	}
}
