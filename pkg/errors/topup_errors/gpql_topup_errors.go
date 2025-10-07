package topup_errors

import (
	"net/http"

	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/response"
)

var (
	ErrGraphqlTopupNotFound     = response.NewGraphqlError("topup", "Topup not found", int(http.StatusNotFound))
	ErrGraphqlTopupInvalidID    = response.NewGraphqlError("topup", "Invalid Topup ID", int(http.StatusBadRequest))
	ErrGraphqlTopupInvalidMonth = response.NewGraphqlError("month", "Invalid Topup Month", int(http.StatusBadRequest))
	ErrGraphqlInvalidCardNumber = response.NewGraphqlError("card_id", "Invalid card number", int(http.StatusBadRequest))
	ErrGraphqlTopupInvalidYear  = response.NewGraphqlError("year", "Invalid Topup Year", int(http.StatusBadRequest))

	ErrGraphqlValidateCreateTopup = response.NewGraphqlError("topup", "Invalid input for create topup", int(http.StatusBadRequest))
	ErrGraphqlValidateUpdateTopup = response.NewGraphqlError("topup", "Invalid input for update topup", int(http.StatusBadRequest))
)
