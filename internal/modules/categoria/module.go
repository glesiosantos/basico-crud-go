package categoria

import "github.com/jackc/pgx/v5/pgxpool"

func NewModule(db *pgxpool) *Handler {
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	return handler
}