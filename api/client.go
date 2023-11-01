package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Client struct {
	client *http.Client
}

func NewClient() (*Client, error) {
	return &Client{
		client: http.DefaultClient,
	}, nil
}

func (c Client) GetAsset(quantity string, apiKey string) ([]Asset, error) {
	url := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?limit=%s&%s", quantity, apiKey)
	resp, err := c.client.Get(url)

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
	err = json.Unmarshal(body, &catImages)
	if err != nil {
		log.Fatal(err)
	}

	return catImages, nil
}
