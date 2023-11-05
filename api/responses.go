package api

import "fmt"

// Asset представляет информацию об активе.
type Asset struct {
	Id     string `json:"id"`     // Уникальный идентификатор актива
	Url    string `json:"url"`    // URL актива
	Width  int    `json:"width"`  // Ширина актива
	Height int    `json:"height"` // Высота актива
}

// Info возвращает строковое представление информации об активе.
func (a Asset) Info() string {
	return fmt.Sprintf("ID: %s\nURL: %s\nWidth: %d\nHeight: %d\n\n",
		a.Id,
		a.Url,
		a.Width,
		a.Height)
}
