package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	client *http.Client
}

func NewClient() (*Client, error) {
	return &Client{
		client: http.DefaultClient, // Создание нового клиента с использованием стандартного HTTP-клиента
	}, nil
}

func (c Client) PostAsset(file, url, apiKey string) (string, error) {
	// Создайте клиент Resty
	client := resty.New()

	// Установите заголовок x-api-key
	client.SetHeader("x-api-key", apiKey)

	// Загружаем изображение и другие параметры
	resp, err := client.R().
		SetFile("file", file). // Устанавливаем файл для загрузки
		SetFormData(map[string]string{
			"sub_id": "my-user-1", // Устанавливаем форм-данные
		}).
		Post(url) // Выполняем HTTP POST-запрос

	if err != nil {
		return "", err
	}

	// Вывод ответа на экран
	return resp.String(), nil
}

func (c Client) GetAsset(quantity string, apiKey string) ([]Asset, error) {
	url := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?limit=%s&%s", quantity, apiKey)

	resp, err := c.client.Get(url) // Выполняем HTTP GET-запрос

	if err != nil {
		return []Asset{}, err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Printf("Response status: %d\n", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return []Asset{}, err
	}

	var catImages []Asset
	err = json.Unmarshal(body, &catImages) // Распаковываем JSON-ответ в структуру
	if err != nil {
		log.Fatal(err)
	}

	return catImages, nil
}
