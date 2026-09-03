package middleware

import (
	"net/http"

	"basico-crud-go/infra/auth"
)

func Auth(
	keycloak *auth.Keycloak,
) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {

				// Recupera Authorization
				authHeader := r.Header.Get(
					"Authorization",
				)

				// Extrai Bearer Token
				rawToken, err := ExtractBearerToken(
					authHeader,
				)

				if err != nil {
					http.Error(
						w,
						err.Error(),
						http.StatusUnauthorized,
					)
					return
				}

				// Valida JWT
				token, err := keycloak.Verifier.Verify(
					r.Context(),
					rawToken,
				)

				if err != nil {
					http.Error(
						w,
						"token inválido ou expirado",
						http.StatusUnauthorized,
					)
					return
				}

				// Extrai informações do usuário
				var claims UserClaims

				if err := token.Claims(&claims); err != nil {
					http.Error(
						w,
						"erro ao ler informações do token",
						http.StatusUnauthorized,
					)
					return
				}

				// Coloca usuário no Context
				ctx := WithUser(
					r.Context(),
					claims,
				)

				// Continua para o Handler
				next.ServeHTTP(
					w,
					r.WithContext(ctx),
				)
			},
		)
	}
}