package role_errors

import (
	"net/http"

	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/response"
)

var (
	ErrGraphqlRoleNotFound      = response.NewGraphqlError("error", "Role not found", int(http.StatusNotFound))
	ErrGraphqlRoleInvalidId     = response.NewGraphqlError("error", "Invalid Role ID", int(http.StatusNotFound))
	ErrGraphqlRoleInvalidUserId = response.NewGraphqlError("error", "Invalid Role User ID", int(http.StatusNotFound))

	ErrGraphqlValidateCreateRole = response.NewGraphqlError("error", "validation failed: invalid create Role request", int(http.StatusBadRequest))
	ErrGraphqlValidateUpdateRole = response.NewGraphqlError("error", "validation failed: invalid update Role request", int(http.StatusBadRequest))
)
