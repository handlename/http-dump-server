package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		logger.Println("--- http request received ---")

		dump, err := httputil.DumpRequest(req, true)
		if err != nil {
			logger.Println("error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		logger.Println(string(dump))

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok")
	})

	logger.Println("server running at http://localhost:8080")
	logger.Fatal(http.ListenAndServe(":8080", mux))
}
