package server

import (
	"fmt"
	"net/http"
)

func StartWebServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "✅ Bot is alive!")
	})

	port := ":8080"
	http.ListenAndServe(port, nil)
}
