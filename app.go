package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/hysem/top-word-service/topword"
)

const N = 10

func main() {
	topwordUsecase := topword.NewUsecase()
	topwordHandler := topword.NewHandler(topwordUsecase)

	mux := http.NewServeMux()
	mux.HandleFunc("/top-words", topwordHandler.FindTopWords)

	address := getEnv("ADDRESS", "0.0.0.0:8080")
	log.Printf("listening@%s\n", address)

	if err := http.ListenAndServe(address, mux); err != nil {
		log.Fatalln("failed to start server")
	}
}

func getEnv(key string, def string) string {
	v := os.Getenv(strings.ToUpper(key))
	if v == "" {
		return def
	}
	return v
}
