package main

import (
	"fmt"
	"log"
	"net/http"
)

func serveFile(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path)
	}
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", serveFile("../static/pages/main.html"))
	router.HandleFunc("/fernandoPessoa", serveFile("../static/pages/fernando.html"))
	router.HandleFunc("/puschkin", serveFile("../static/pages/puschkin.html"))
	router.HandleFunc("/horacio", serveFile("../static/pages/horacio.html"))

	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static/style"))))

	fmt.Println("Servidor rodando na porta 8082")

	if err := http.ListenAndServe(":8082", router); err != nil {
		log.Fatal("erro ao inicializar servidor " + err.Error())
	}
}
