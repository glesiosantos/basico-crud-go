package middleware

import (
	"errors"
	"strings"
)

func ExtractBearerToken(authHeader string) (string, error) {

	if authHeader == "" {
		return "", errors.New("token não informado")
	}

	parts := strings.SplitN(
		authHeader,
		" ",
		2,
	)

	if len(parts) != 2 {
		return "", errors.New("formato do token inválido")
	}

	if !strings.EqualFold(parts[0], "Bearer") {
		return "", errors.New("tipo de autenticação inválido")
	}

	token := strings.TrimSpace(parts[1])

	if token == "" {
		return "", errors.New("token não informado")
	}

	return token, nil
}