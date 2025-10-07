package userrole_errors

import (
	"net/http"

	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/response"
)

var (
	ErrFailedAssignRoleToUser = response.NewErrorResponse("Failed to assign role to user", http.StatusInternalServerError)
	ErrFailedRemoveRole       = response.NewErrorResponse("Failed to remove role from user", http.StatusInternalServerError)
)
