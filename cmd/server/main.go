package main

import (
	"fmt"
	"log"
	"net/http"

	"go-board-app/internal/middleware"
	"go-board-app/internal/router"
)

func main() {
	r := router.NewRouter()
	handler := middleware.EnableCORS(r)

	fmt.Println("✅ Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
