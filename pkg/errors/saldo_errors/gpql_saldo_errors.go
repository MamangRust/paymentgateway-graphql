package saldo_errors

import (
	"net/http"

	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/response"
)

var (
	ErrGraphqlSaldoNotFound          = response.NewGraphqlError("saldo", "Saldo not found", int(http.StatusNotFound))
	ErrGraphqlSaldoInvalidID         = response.NewGraphqlError("saldo", "Invalid Saldo ID", int(http.StatusBadRequest))
	ErrGraphqlSaldoInvalidCardNumber = response.NewGraphqlError("saldo", "Invalid Saldo Card Number", int(http.StatusBadRequest))
	ErrGraphqlSaldoInvalidMonth      = response.NewGraphqlError("saldo", "Invalid Saldo Month", int(http.StatusBadRequest))
	ErrGraphqlSaldoInvalidYear       = response.NewGraphqlError("saldo", "Invalid Saldo Year", int(http.StatusBadRequest))

	ErrGraphqlValidateCreateSaldo = response.NewGraphqlError("saldo", "Invalid input for create saldo", int(http.StatusBadRequest))
	ErrGraphqlValidateUpdateSaldo = response.NewGraphqlError("saldo", "Invalid input for update saldo", int(http.StatusBadRequest))
)
