package merchant_errors

import (
	"net/http"

	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/response"
)

var (
	ErrGraphqlMerchantNotFound       = response.NewGraphqlError("merchant", "Merchant not found", int(http.StatusNotFound))
	ErrGraphqlMerchantInvalidID      = response.NewGraphqlError("merchant", "Invalid Merchant ID", int(http.StatusBadRequest))
	ErrGraphqlMerchantInvalidUserID  = response.NewGraphqlError("merchant", "Invalid Merchant User ID", int(http.StatusBadRequest))
	ErrGraphqlMerchantInvalidApiKey  = response.NewGraphqlError("merchant", "Invalid Merchant Api Key", int(http.StatusBadRequest))
	ErrGraphqlMerchantInvalidMonth   = response.NewGraphqlError("month", "Invalid Merchant Month", int(http.StatusBadRequest))
	ErrGraphqlMerchantInvalidYear    = response.NewGraphqlError("year", "Invalid Merchant Year", int(http.StatusBadRequest))
	ErrGraphqlValidateCreateMerchant = response.NewGraphqlError("merchant", "Invalid input for create merchant", int(http.StatusBadRequest))
	ErrGraphqlValidateUpdateMerchant = response.NewGraphqlError("merchant", "Invalid input for update merchant", int(http.StatusBadRequest))
)
