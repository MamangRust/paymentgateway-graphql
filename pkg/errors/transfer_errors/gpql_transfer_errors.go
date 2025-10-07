package transfer_errors

import (
	"net/http"

	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/response"
)

var (
	ErrGraphqlTransferNotFound  = response.NewGraphqlError("transfer", "Transfer not found", int(http.StatusNotFound))
	ErrGraphqlTransferInvalidID = response.NewGraphqlError("transfer", "Invalid Transfer ID", int(http.StatusBadRequest))
	ErrGraphqlInvalidUserID     = response.NewGraphqlError("card_id", "Invalid user ID", int(http.StatusBadRequest))
	ErrGraphqlInvalidCardNumber = response.NewGraphqlError("card_id", "Invalid card number", int(http.StatusBadRequest))
	ErrGraphqlInvalidMonth      = response.NewGraphqlError("month", "Invalid month", int(http.StatusBadRequest))
	ErrGraphqlInvalidYear       = response.NewGraphqlError("year", "Invalid year", int(http.StatusBadRequest))

	ErrGraphqlValidateCreateTransferRequest = response.NewGraphqlError("transfer", "Invalid input for create transfer", int(http.StatusBadRequest))
	ErrGraphqlValidateUpdateTransferRequest = response.NewGraphqlError("transfer", "Invalid input for update transfer", int(http.StatusBadRequest))
)
