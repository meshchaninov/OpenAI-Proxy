package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	apiURL := "https://api.openai.com"
	apiKey := os.Getenv("OPENAI_API_KEY")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		req, err := http.NewRequest(r.Method, apiURL+r.URL.Path, r.Body)
		if err != nil {
			http.Error(w, "Failed to create request", http.StatusInternalServerError)
			return
		}

		req.Header = r.Header
		req.Header.Set("Authorization", "Bearer "+apiKey)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Failed to reach OpenAI API", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		for key, value := range resp.Header {
			w.Header()[key] = value
		}

		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	})

	port := "8080"
	log.Printf("Starting proxy server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
