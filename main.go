package main

import (
	"context"
	"fmt"
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
	p "basico-crud-go/produtos"
)

func main() {
	url := "postgres://postgres:102030@localhost:5432/basicodb"
	db, err := pgxpool.New(context.Background(), url)

	if err != nil {
		log.Fatal("Erro de conexão", err)
	}

	defer db.Close()

	produto := p.Produto{
		Descricao: "Monitor de 22 polegadas",
		Preco: 1900.0,
		Quantidade: 10,
		Categoria: p.Categoria {
			Id: 2,
		},
	}

	err = p.AddProduto(db, produto)

	if err != nil {
		log.Fatal("Error.: ", err)
	}

	produtos, err := p.ListarProdutos(db)

	if err != nil {
		log.Fatal("Error.: ", err)
	}

	for _, produto := range produtos {
		fmt.Printf("%d - %s - %s\n", 
			produto.Id, 
			produto.Descricao, 
			produto.Categoria.Nome,
		)
	}

}