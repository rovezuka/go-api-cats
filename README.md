# Проект API-клиента для работы с котами

Этот проект представляет собой простой API-клиент, разработанный для взаимодействия с API, предоставляющим изображения котов. Он включает в себя функциональность для выполнения GET- и POST-запросов к API.

## Установка

Для использования этого API-клиента вам потребуется Go (версия 1.11 и выше).

1. Склонируйте репозиторий:

 ```bash
git clone https://github.com/rovezuka/go-api-cats.git
```

2. Перейдите в директорию проекта:
```bash
cd go-api-cats
```

3. Замените YOUR_API_KEY на ваш ключ API в файле main.go в соответствии с комментариями в коде.

4. Запустите приложение:
```go
go run main.go
```

## Использование
API-клиент предоставляет два основных метода: GetAsset и PostAsset, которые позволяют выполнять GET- и POST-запросы к API для получения изображений котов и отправки файлов.

### Пример использования
```go
package main

import (
	"fmt"
	"log"

	"github.com/rovezuka/go-api-cats/api"
)

const API_KEY = "YOUR_API_KEY"

func main() {
	apiClient, err := api.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	catImages, err := apiClient.GetAsset("10", API_KEY)
	if err != nil {
		log.Fatal(err)
	}

	for _, catImage := range catImages {
		fmt.Println(catImage.Info())
	}

	apiClient.PostAsset("FILE_PATH", "POST_URL", API_KEY)
}
```

В приведенном выше примере API_KEY должен быть заменен на ваш ключ API, а также указаны пути к файлам и URL для POST-запроса.

## Структура проекта
- `api/client.go`: Код, реализующий клиент API.
- `api/responses.go`: Определение структуры Asset для представления информации об активе.

- `main.go`: Пример использования API-клиента.

## Зависимости
Этот проект использует следующие зависимости:

- `github.com/go-resty/resty/v2`: Библиотека для выполнения HTTP-запросов.
