package categoria

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/go-chi/chi/v5"
)

func NewRegisterModule(db *pgxpool.Pool, router chi.Router){
	repository := newRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	router.Get("/categorias", handler.ListarCategorias,)
	router.Post("/categorias", handler.CadastrarCategorias,)
}