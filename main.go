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

	catImages, err := apiClient.GetAsset("10", "live_c1vDY98eMovuABHvsqTjkgnKRs1qm5vnDatRDMKE7crDh94FMmWopnxgq9Jd8wvd")
	if err != nil {
		log.Fatal(err)
	}

	for _, catImage := range catImages {
		fmt.Println(catImage.Info())
	}
}
