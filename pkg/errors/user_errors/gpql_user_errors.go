package user_errors

import (
	"net/http"

	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/response"
)

var (
	ErrGraphqlUserNotFound  = response.NewGraphqlError("error", "User not found", int(http.StatusNotFound))
	ErrGraphqlUserInvalidId = response.NewGraphqlError("error", "Invalid User ID", int(http.StatusNotFound))
)
