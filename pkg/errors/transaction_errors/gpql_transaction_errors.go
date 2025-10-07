package transaction_errors

import (
	"net/http"

	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/response"
)

var (
	ErrGraphqlTransactionNotFound          = response.NewGraphqlError("transaction", "Transaction not found", int(http.StatusNotFound))
	ErrGraphqlTransactionInvalidID         = response.NewGraphqlError("transaction", "Invalid Transaction ID", int(http.StatusBadRequest))
	ErrGraphqlTransactionInvalidMerchantID = response.NewGraphqlError("transaction", "Invalid Transaction Merchant ID", int(http.StatusBadRequest))
	ErrGraphqlInvalidCardNumber            = response.NewGraphqlError("card_id", "Invalid card number", int(http.StatusBadRequest))
	ErrGraphqlInvalidMonth                 = response.NewGraphqlError("month", "Invalid month", int(http.StatusBadRequest))
	ErrGraphqlInvalidYear                  = response.NewGraphqlError("year", "Invalid year", int(http.StatusBadRequest))

	ErrGraphqlValidateCreateTransactionRequest = response.NewGraphqlError("transaction", "Invalid input for create card", int(http.StatusBadRequest))
	ErrGraphqlValidateUpdateTransactionRequest = response.NewGraphqlError("transaction", "Invalid input for update card", int(http.StatusBadRequest))
)
