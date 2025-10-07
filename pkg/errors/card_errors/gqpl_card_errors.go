package card_errors

import (
	"net/http"

	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/response"
)

var (
	ErrGraphqlInvalidCardID     = response.NewGraphqlError("card_id", "Invalid card ID", int(http.StatusBadRequest))
	ErrGraphqlInvalidUserID     = response.NewGraphqlError("card_id", "Invalid user ID", int(http.StatusBadRequest))
	ErrGraphqlInvalidCardNumber = response.NewGraphqlError("card_id", "Invalid card number", int(http.StatusBadRequest))
	ErrGraphqlInvalidMonth      = response.NewGraphqlError("month", "Invalid month", int(http.StatusBadRequest))
	ErrGraphqlInvalidYear       = response.NewGraphqlError("year", "Invalid year", int(http.StatusBadRequest))
)
