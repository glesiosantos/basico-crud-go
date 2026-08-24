package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	url := "postgres://postgres:102030@localhost:5432/basicodb"
	db, err := pgxpool.New(context.Background(), url)

	if err != nil {
		log.Fatal("Erro de conexão", err)
	}

	defer db.Close()

	sql := `
		INSERT INTO produtos (descricao, preco, quantidade)
		VALUES ($1,$2,$3)
	`
	_, err = db.Exec(
		context.Background(),
		sql,
		"Projeto Epson de 1800 Lumines",
		2995.9,
		1,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Produto registrado com sucesso!!!" )
}