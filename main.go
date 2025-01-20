package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Задаем адрес OpenAI API
	apiURL := "https://api.openai.com"

	// Функция-обработчик запросов
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Создаем новый запрос, направленный на OpenAI API
		req, err := http.NewRequest(r.Method, apiURL+r.URL.Path, r.Body)
		if err != nil {
			http.Error(w, "Failed to create request", http.StatusInternalServerError)
			return
		}

		// Копируем все заголовки из оригинального запроса
		req.Header = r.Header

		// Выполняем запрос к OpenAI API
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Failed to reach OpenAI API", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Копируем заголовки ответа
		for key, value := range resp.Header {
			w.Header()[key] = value
		}

		// Устанавливаем статус ответа и копируем тело ответа
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	})

	// Запуск сервера
	port := "8080"
	log.Printf("Starting proxy server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
