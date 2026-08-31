package main

import (
	"context"
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"basico-crud-go/internal/modules/categoria"

)

func main() {
	url := "postgres://postgres:102030@localhost:5432/basicodb"
	db, err := pgxpool.New(context.Background(), url)

	if err != nil {
		log.Fatal("Erro de conexão", err)
	}

	defer db.Close()

	repository := categoria.NewRepository(db)
	service := categoria.NewService(repository)
	handler := categoria.NewHandler(service)

	router := chi.NewRouter()

	router.Get("/categorias", handler.ListarCategorias,)
	router.Post("/categorias", handler.CadastrarCategorias,)

	log.Println(
        "Servidor executando em http://localhost:8081",
    )

	err = http.ListenAndServe(
		":8081",
		router,
	)

	if err != nil {
		log.Fatal(err)
	}
}