package api

import "fmt"

type Asset struct {
	Id     string `json:"id"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func (a Asset) Info() string {
	return fmt.Sprintf("ID: %s\nURL: %s\nWidth: %d\nHeight: %d\n\n",
		a.Id,
		a.Url,
		a.Width,
		a.Height)
}
