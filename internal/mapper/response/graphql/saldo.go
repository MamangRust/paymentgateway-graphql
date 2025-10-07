package graphql

import (
	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/response"
	"github.com/MamangRust/paymentgatewaygraphql/internal/graph/model"
)

type saldoResponse struct {
}

func NewSaldoResponseMapper() *saldoResponse {
	return &saldoResponse{}
}

func (s *saldoResponse) ToGraphqlResponseSaldo(status, message string, saldo *response.SaldoResponse) *model.APIResponseSaldoResponse {
	return &model.APIResponseSaldoResponse{
		Status:  status,
		Message: message,
		Data:    s.mapResponseSaldo(saldo),
	}
}

func (s *saldoResponse) ToGraphqlResponsesSaldo(status, message string, saldo []*response.SaldoResponse) *model.APIResponsesSaldo {
	return &model.APIResponsesSaldo{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesSaldo(saldo),
	}
}

func (s *saldoResponse) ToGraphqlResponseSaldoDeleteAt(status, message string, saldo *response.SaldoResponseDeleteAt) *model.APIResponseSaldoResponseDeleteAt {
	return &model.APIResponseSaldoResponseDeleteAt{
		Status:  status,
		Message: message,
		Data:    s.mapResponseSaldoDeleteAt(saldo),
	}
}

func (s *saldoResponse) ToGraphqlResponsePaginationSaldo(status, message string, saldo []*response.SaldoResponse, pagination *response.PaginationMeta) *model.APIResponsePaginationSaldo {
	return &model.APIResponsePaginationSaldo{
		Status:     status,
		Message:    message,
		Data:       s.mapResponsesSaldo(saldo),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (s *saldoResponse) ToGraphqlResponsePaginationSaldoDeleteAt(status, message string, saldo []*response.SaldoResponseDeleteAt, pagination *response.PaginationMeta) *model.APIResponsePaginationSaldoDeleteAt {
	return &model.APIResponsePaginationSaldoDeleteAt{
		Status:     status,
		Message:    message,
		Data:       s.mapResponsesSaldoDeleteAt(saldo),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (s *saldoResponse) ToGraphqlResponseDelete(status, message string) *model.APIResponseSaldoDelete {
	return &model.APIResponseSaldoDelete{
		Status:  status,
		Message: message,
	}
}
func (s *saldoResponse) ToGraphqlResponseAll(status, message string) *model.APIResponseSaldoAll {
	return &model.APIResponseSaldoAll{
		Status:  status,
		Message: message,
	}
}

func (s *saldoResponse) ToGraphqlResponseMonthTotalSaldo(status, message string, response []*response.SaldoMonthTotalBalanceResponse) *model.APIResponseMonthTotalSaldo {
	return &model.APIResponseMonthTotalSaldo{
		Status:  status,
		Message: message,
		Data:    s.mapSaldoMonthTotalBalanceResponses(response),
	}
}

func (s *saldoResponse) ToGraphqlResponseYearTotalSaldo(status, message string, response []*response.SaldoYearTotalBalanceResponse) *model.APIResponseYearTotalSaldo {
	return &model.APIResponseYearTotalSaldo{
		Status:  status,
		Message: message,
		Data:    s.mapSaldoYearTotalBalanceResponses(response),
	}
}

func (s *saldoResponse) ToGraphqlResponseMonthSaldoBalances(status, message string, response []*response.SaldoMonthBalanceResponse) *model.APIResponseMonthSaldoBalances {
	return &model.APIResponseMonthSaldoBalances{
		Status:  status,
		Message: message,
		Data:    s.mapSaldoMonthBalanceResponses(response),
	}
}

func (s *saldoResponse) ToGraphqlResponseYearBalance(status, message string, response []*response.SaldoYearBalanceResponse) *model.APIResponseYearSaldoBalances {
	return &model.APIResponseYearSaldoBalances{
		Status:  status,
		Message: message,
		Data:    s.mapSaldoYearBalanceResponses(response),
	}
}

func (s *saldoResponse) mapResponseSaldo(saldo *response.SaldoResponse) *model.SaldoResponse {
	withdrawAmount := int32(saldo.WithdrawAmount)

	return &model.SaldoResponse{
		ID:             int32(saldo.ID),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int32(saldo.TotalBalance),
		WithdrawTime:   &saldo.WithdrawTime,
		WithdrawAmount: &withdrawAmount,
		CreatedAt:      saldo.CreatedAt,
		UpdatedAt:      saldo.UpdatedAt,
	}
}

func (s *saldoResponse) mapResponsesSaldo(saldos []*response.SaldoResponse) []*model.SaldoResponse {
	var responseSaldos []*model.SaldoResponse

	for _, saldo := range saldos {
		responseSaldos = append(responseSaldos, s.mapResponseSaldo(saldo))
	}

	return responseSaldos
}

func (s *saldoResponse) mapResponseSaldoDeleteAt(saldo *response.SaldoResponseDeleteAt) *model.SaldoResponseDeleteAt {
	withdrawAmount := int32(saldo.WithdrawAmount)

	return &model.SaldoResponseDeleteAt{
		ID:             int32(saldo.ID),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int32(saldo.TotalBalance),
		WithdrawTime:   &saldo.WithdrawTime,
		WithdrawAmount: &withdrawAmount,
		CreatedAt:      saldo.CreatedAt,
		UpdatedAt:      saldo.UpdatedAt,
	}
}

func (s *saldoResponse) mapResponsesSaldoDeleteAt(saldos []*response.SaldoResponseDeleteAt) []*model.SaldoResponseDeleteAt {
	var responseSaldos []*model.SaldoResponseDeleteAt

	for _, saldo := range saldos {
		responseSaldos = append(responseSaldos, s.mapResponseSaldoDeleteAt(saldo))
	}

	return responseSaldos
}

func (s *saldoResponse) mapSaldoMonthTotalBalanceResponse(saldo *response.SaldoMonthTotalBalanceResponse) *model.SaldoMonthTotalBalanceResponse {
	totalBalance := 0

	if saldo.TotalBalance != 0 {
		totalBalance = saldo.TotalBalance
	}

	return &model.SaldoMonthTotalBalanceResponse{
		Month:        saldo.Month,
		Year:         saldo.Year,
		TotalBalance: int32(totalBalance),
	}
}

func (s *saldoResponse) mapSaldoMonthTotalBalanceResponses(saldos []*response.SaldoMonthTotalBalanceResponse) []*model.SaldoMonthTotalBalanceResponse {
	var responsesSaldo []*model.SaldoMonthTotalBalanceResponse

	for _, saldo := range saldos {
		responsesSaldo = append(responsesSaldo, s.mapSaldoMonthTotalBalanceResponse(saldo))
	}

	return responsesSaldo
}

func (s *saldoResponse) mapSaldoYearTotalBalanceResponse(saldo *response.SaldoYearTotalBalanceResponse) *model.SaldoYearTotalBalanceResponse {
	totalBalance := 0

	if saldo.TotalBalance != 0 {
		totalBalance = saldo.TotalBalance
	}

	return &model.SaldoYearTotalBalanceResponse{
		Year:         saldo.Year,
		TotalBalance: int32(totalBalance),
	}
}

func (s *saldoResponse) mapSaldoYearTotalBalanceResponses(saldos []*response.SaldoYearTotalBalanceResponse) []*model.SaldoYearTotalBalanceResponse {
	var responsesSaldo []*model.SaldoYearTotalBalanceResponse

	for _, saldo := range saldos {
		responsesSaldo = append(responsesSaldo, s.mapSaldoYearTotalBalanceResponse(saldo))
	}

	return responsesSaldo
}

func (s *saldoResponse) mapSaldoMonthBalanceResponse(saldo *response.SaldoMonthBalanceResponse) *model.SaldoMonthBalanceResponse {
	totalBalance := 0

	if saldo.TotalBalance != 0 {
		totalBalance = saldo.TotalBalance
	}

	return &model.SaldoMonthBalanceResponse{
		Month:        saldo.Month,
		TotalBalance: int32(totalBalance),
	}
}

func (s *saldoResponse) mapSaldoMonthBalanceResponses(saldos []*response.SaldoMonthBalanceResponse) []*model.SaldoMonthBalanceResponse {
	var responsesSaldo []*model.SaldoMonthBalanceResponse

	for _, saldo := range saldos {
		responsesSaldo = append(responsesSaldo, s.mapSaldoMonthBalanceResponse(saldo))
	}

	return responsesSaldo
}

func (s *saldoResponse) mapSaldoYearBalanceResponse(saldo *response.SaldoYearBalanceResponse) *model.SaldoYearBalanceResponse {
	totalBalance := 0

	if saldo.TotalBalance != 0 {
		totalBalance = saldo.TotalBalance
	}

	return &model.SaldoYearBalanceResponse{
		Year:         saldo.Year,
		TotalBalance: int32(totalBalance),
	}
}

func (s *saldoResponse) mapSaldoYearBalanceResponses(saldos []*response.SaldoYearBalanceResponse) []*model.SaldoYearBalanceResponse {
	var responsesSaldo []*model.SaldoYearBalanceResponse

	for _, saldo := range saldos {
		responsesSaldo = append(responsesSaldo, s.mapSaldoYearBalanceResponse(saldo))
	}

	return responsesSaldo
}
