package middleware

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

func GetUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	val := ctx.Value(UserID)
	str, ok := val.(string)
	if !ok {
		return uuid.Nil, errors.New("userID not found or invalid")
	}
	return uuid.Parse(str)
}
