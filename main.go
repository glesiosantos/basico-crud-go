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

	//Produto
	// Descricao, Preco, Quantidade 


	sql := `
		INSERT INTO produtos (descricao, preco, quantidade)
		VALUE ($1,$2,$3)
	`

	_, err := db.Exec(
		context.Background(),
		sql,
		"Cadeira de Vermelha Diretora com rodinhas e Apoio",
		675.90,
		10
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Produto registrado com sucesso!!!" )
}