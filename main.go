package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type errorBody struct {
	Error string `json:"error"`
}

type Response struct {
	Status string `json:"status"`
}

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")

	mux := http.NewServeMux()

	corsMux := corsMiddleware(mux)

	mux.HandleFunc("/v1/readiness", readinessHandlerFunc)
	mux.HandleFunc("/v1/err", errorHandlerFunc)

	server := &http.Server{
		Addr:    port,
		Handler: corsMux,
	}
	log.Fatal(server.ListenAndServe())
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func readinessHandlerFunc(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Status: "ok",
	}
	respondWithJSON(w, 200, resp)
}

func errorHandlerFunc(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 200, "Internal Server Error")
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if payload != nil {
		response, err := json.Marshal(payload)
		if err != nil {
			log.Println("error marshalling", err)
			w.WriteHeader(500)
			response, _ := json.Marshal(errorBody{
				Error: "error marshalling",
			})
			w.Write(response)
			return
		}
		w.WriteHeader(code)
		w.Write(response)
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, errorBody{
		Error: msg,
	})
}
