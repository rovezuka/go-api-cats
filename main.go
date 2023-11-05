package main

import (
	"fmt"
	"log"

	"github.com/rovezuka/go-api-cats/api"
)

const API_KEY = "YOUR_API_KEY" // Константа для хранения ключа API

func main() {
	// Создание нового клиента API
	apiClient, err := api.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	// Выполнение GET-запроса для получения изображений
	catImages, err := apiClient.GetAsset("10", API_KEY)
	if err != nil {
		log.Fatal(err)
	}

	// Вывод информации об изображениях
	for _, catImage := range catImages {
		fmt.Println(catImage.Info())
	}

	// Выполнение POST-запроса для отправки файла
	apiClient.PostAsset("FILE_PATH", "POST_URL", API_KEY)
}
