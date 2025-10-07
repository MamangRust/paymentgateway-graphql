package withdraw_errors

import (
	"net/http"

	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/response"
)

var (
	ErrGraphqlWithdrawNotFound  = response.NewGraphqlError("withdraw", "Withdraw not found", int(http.StatusNotFound))
	ErrGraphqlWithdrawInvalidID = response.NewGraphqlError("withdraw", "Invalid Withdraw ID", int(http.StatusBadRequest))
	ErrGraphqlInvalidUserID     = response.NewGraphqlError("card_id", "Invalid user ID", int(http.StatusBadRequest))
	ErrGraphqlInvalidCardNumber = response.NewGraphqlError("card_id", "Invalid card number", int(http.StatusBadRequest))
	ErrGraphqlInvalidMonth      = response.NewGraphqlError("month", "Invalid month", int(http.StatusBadRequest))
	ErrGraphqlInvalidYear       = response.NewGraphqlError("year", "Invalid year", int(http.StatusBadRequest))

	ErrGraphqlValidateCreateWithdrawRequest = response.NewGraphqlError("withdraw", "Invalid input for create withdraw", int(http.StatusBadRequest))
	ErrGraphqlValidateUpdateWithdrawRequest = response.NewGraphqlError("withdraw", "Invalid input for update withdraw", int(http.StatusBadRequest))
)
