package middleware

import "context"

type contextKey string

const userContextKey contextKey = "user"

func WithUser(
	ctx context.Context,
	user UserClaims,
) context.Context {

	return context.WithValue(
		ctx,
		userContextKey,
		user,
	)
}

func GetUser(
	ctx context.Context,
) (UserClaims, bool) {

	user, ok := ctx.Value(
		userContextKey,
	).(UserClaims)

	return user, ok
}