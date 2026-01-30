package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Estructura para respuesta JSON
type Response struct {
	Message string `json:"message"`
	Status string `json:"status"`
}

// Handler para el endpoint de salud
func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Applicación funcionando correctamente",
		Status:  "healthy",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

// Handler para el endpoint principal
func homeHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "¡Bienvenido a mi aplicación DevOps!",
		Status:  "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func main(){
	// Crear router
	r := mux.NewRouter()

	// Definir rutas
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/health", healthHandler).Methods("GET")

	// Puerto desde variable de entorno o 8080 por defecto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Servidor iniciando en puerto %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}