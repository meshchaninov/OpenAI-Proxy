package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	apiURL := "https://api.openai.com"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)

		req, err := http.NewRequest(r.Method, apiURL+r.URL.Path, r.Body)
		if err != nil {
			log.Printf("Error creating request: %v", err)
			http.Error(w, "Failed to create request", http.StatusInternalServerError)
			return
		}

		req.Header = r.Header

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Error reaching OpenAI API: %v", err)
			http.Error(w, "Failed to reach OpenAI API", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		for key, value := range resp.Header {
			w.Header()[key] = value
		}

		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)

		log.Printf("Responded with status code: %d", resp.StatusCode)
	})

	port := "8080"
	log.Printf("Starting proxy server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
