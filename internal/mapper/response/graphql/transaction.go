package graphql

import (
	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/response"
	"github.com/MamangRust/paymentgatewaygraphql/internal/graph/model"
)

type transactionResponseMapper struct {
}

func NewTransactionResponseMapper() *transactionResponseMapper {
	return &transactionResponseMapper{}
}

func (t *transactionResponseMapper) ToGraphqlResponseTransactionMonthStatusSuccess(status, message string, data []*response.TransactionResponseMonthStatusSuccess) *model.APIResponseTransactionMonthStatusSuccess {
	return &model.APIResponseTransactionMonthStatusSuccess{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTransactionMonthStatusSuccess(data),
	}
}

func (t *transactionResponseMapper) ToGraphqlResponseTransactionYearStatusSuccess(status, message string, data []*response.TransactionResponseYearStatusSuccess) *model.APIResponseTransactionYearStatusSuccess {
	return &model.APIResponseTransactionYearStatusSuccess{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTransactionYearStatusSuccess(data),
	}
}

func (t *transactionResponseMapper) ToGraphqlResponseTransactionMonthStatusFailed(status, message string, data []*response.TransactionResponseMonthStatusFailed) *model.APIResponseTransactionMonthStatusFailed {
	return &model.APIResponseTransactionMonthStatusFailed{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTransactionMonthStatusFailed(data),
	}
}

func (t *transactionResponseMapper) ToGraphqlResponseTransactionYearStatusFailed(status, message string, data []*response.TransactionResponseYearStatusFailed) *model.APIResponseTransactionYearStatusFailed {
	return &model.APIResponseTransactionYearStatusFailed{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTransactionYearStatusFailed(data),
	}
}

func (t *transactionResponseMapper) ToGraphqlResponseTransactionMonthMethod(status, message string, data []*response.TransactionMonthMethodResponse) *model.APIResponseTransactionMonthMethod {
	return &model.APIResponseTransactionMonthMethod{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTransactionMonthlyMethod(data),
	}
}

func (t *transactionResponseMapper) ToGraphqlResponseTransactionYearMethod(status, message string, data []*response.TransactionYearMethodResponse) *model.APIResponseTransactionYearMethod {
	return &model.APIResponseTransactionYearMethod{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTransactionYearlyMethod(data),
	}
}

func (t *transactionResponseMapper) ToGraphqlResponseTransactionMonthAmount(status, message string, data []*response.TransactionMonthAmountResponse) *model.APIResponseTransactionMonthAmount {
	return &model.APIResponseTransactionMonthAmount{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTransactionMonthlyAmount(data),
	}
}

func (t *transactionResponseMapper) ToGraphqlResponseTransactionYearAmount(status, message string, data []*response.TransactionYearlyAmountResponse) *model.APIResponseTransactionYearAmount {
	return &model.APIResponseTransactionYearAmount{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTransactionYearlyAmount(data),
	}
}

func (t *transactionResponseMapper) ToGraphqlTransactionAll(status, message string) *model.APIResponseTransactionAll {
	return &model.APIResponseTransactionAll{
		Status:  status,
		Message: message,
	}
}

func (t *transactionResponseMapper) ToGraphqlTransactionDelete(status, message string) *model.APIResponseTransactionDelete {
	return &model.APIResponseTransactionDelete{
		Status:  status,
		Message: message,
	}
}

func (t *transactionResponseMapper) ToGraphqlResponseTransaction(status, message string, data *response.TransactionResponse) *model.APIResponseTransaction {
	return &model.APIResponseTransaction{
		Status:  status,
		Message: message,
		Data:    t.mapResponseTransaction(data),
	}
}

func (t *transactionResponseMapper) ToGraphqlResponseTransactions(status, message string, data []*response.TransactionResponse) *model.APIResponseTransactions {
	return &model.APIResponseTransactions{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTransaction(data),
	}
}

func (t *transactionResponseMapper) ToGraphqlResponseTransactionDeleteAt(status, message string, data *response.TransactionResponseDeleteAt) *model.APIResponseTransactionDeleteAt {
	return &model.APIResponseTransactionDeleteAt{
		Status:  status,
		Message: message,
		Data:    t.mapResponseTransactionDeleteAt(data),
	}
}

func (t *transactionResponseMapper) ToGraphqlPaginationTransaction(status, message string, data []*response.TransactionResponse, pagination *response.PaginationMeta) *model.APIResponsePaginationTransaction {
	return &model.APIResponsePaginationTransaction{
		Status:     status,
		Message:    message,
		Data:       t.mapResponsesTransaction(data),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (t *transactionResponseMapper) ToGraphqlPaginationTransactionDeleteAt(status, message string, data []*response.TransactionResponseDeleteAt, pagination *response.PaginationMeta) *model.APIResponsePaginationTransactionDeleteAt {
	return &model.APIResponsePaginationTransactionDeleteAt{
		Status:     status,
		Message:    message,
		Data:       t.mapResponsesTransactionDeleteAt(data),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (t *transactionResponseMapper) mapResponseTransaction(transaction *response.TransactionResponse) *model.TransactionResponse {
	return &model.TransactionResponse{
		ID:              int32(transaction.ID),
		TransactionNo:   transaction.TransactionNo,
		CardNumber:      transaction.CardNumber,
		Amount:          int32(transaction.Amount),
		PaymentMethod:   transaction.PaymentMethod,
		TransactionTime: transaction.TransactionTime,
		MerchantID:      int32(transaction.MerchantID),
		CreatedAt:       transaction.CreatedAt,
		UpdatedAt:       transaction.UpdatedAt,
	}
}

func (t *transactionResponseMapper) mapResponsesTransaction(transactions []*response.TransactionResponse) []*model.TransactionResponse {
	var result []*model.TransactionResponse

	for _, transaction := range transactions {
		result = append(result, t.mapResponseTransaction(transaction))
	}

	return result
}

func (t *transactionResponseMapper) mapResponseTransactionDeleteAt(transaction *response.TransactionResponseDeleteAt) *model.TransactionResponseDeleteAt {
	return &model.TransactionResponseDeleteAt{
		ID:              int32(transaction.ID),
		TransactionNo:   transaction.TransactionNo,
		CardNumber:      transaction.CardNumber,
		Amount:          int32(transaction.Amount),
		PaymentMethod:   transaction.PaymentMethod,
		TransactionTime: transaction.TransactionTime,
		MerchantID:      int32(transaction.MerchantID),
		CreatedAt:       transaction.CreatedAt,
		UpdatedAt:       transaction.UpdatedAt,
		DeletedAt:       transaction.DeletedAt,
	}
}

func (t *transactionResponseMapper) mapResponsesTransactionDeleteAt(transactions []*response.TransactionResponseDeleteAt) []*model.TransactionResponseDeleteAt {
	var result []*model.TransactionResponseDeleteAt

	for _, transaction := range transactions {
		result = append(result, t.mapResponseTransactionDeleteAt(transaction))
	}

	return result
}

func (t *transactionResponseMapper) mapResponseTransactionMonthStatusSuccess(s *response.TransactionResponseMonthStatusSuccess) *model.TransactionMonthStatusSuccessResponse {
	return &model.TransactionMonthStatusSuccessResponse{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

func (t *transactionResponseMapper) mapResponsesTransactionMonthStatusSuccess(Transactions []*response.TransactionResponseMonthStatusSuccess) []*model.TransactionMonthStatusSuccessResponse {
	var responses []*model.TransactionMonthStatusSuccessResponse

	for _, Transaction := range Transactions {
		responses = append(responses, t.mapResponseTransactionMonthStatusSuccess(Transaction))
	}

	return responses
}

func (t *transactionResponseMapper) mapResponseTransactionMonthStatusFailed(s *response.TransactionResponseMonthStatusFailed) *model.TransactionMonthStatusFailedResponse {
	return &model.TransactionMonthStatusFailedResponse{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

func (t *transactionResponseMapper) mapResponsesTransactionMonthStatusFailed(Transactions []*response.TransactionResponseMonthStatusFailed) []*model.TransactionMonthStatusFailedResponse {
	var responses []*model.TransactionMonthStatusFailedResponse

	for _, Transaction := range Transactions {
		responses = append(responses, t.mapResponseTransactionMonthStatusFailed(Transaction))
	}

	return responses
}

func (t *transactionResponseMapper) mapResponseTransactionYearStatusSuccess(s *response.TransactionResponseYearStatusSuccess) *model.TransactionYearStatusSuccessResponse {
	return &model.TransactionYearStatusSuccessResponse{
		Year:         s.Year,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

func (t *transactionResponseMapper) mapResponsesTransactionYearStatusSuccess(Transactions []*response.TransactionResponseYearStatusSuccess) []*model.TransactionYearStatusSuccessResponse {
	var responses []*model.TransactionYearStatusSuccessResponse

	for _, Transaction := range Transactions {
		responses = append(responses, t.mapResponseTransactionYearStatusSuccess(Transaction))
	}

	return responses
}

func (t *transactionResponseMapper) mapResponseTransactionYearStatusFailed(s *response.TransactionResponseYearStatusFailed) *model.TransactionYearStatusFailedResponse {
	return &model.TransactionYearStatusFailedResponse{
		Year:        s.Year,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

func (t *transactionResponseMapper) mapResponsesTransactionYearStatusFailed(Transactions []*response.TransactionResponseYearStatusFailed) []*model.TransactionYearStatusFailedResponse {
	var responses []*model.TransactionYearStatusFailedResponse

	for _, Transaction := range Transactions {
		responses = append(responses, t.mapResponseTransactionYearStatusFailed(Transaction))
	}

	return responses
}

func (t *transactionResponseMapper) mapResponseTransactionMonthlyMethod(s *response.TransactionMonthMethodResponse) *model.TransactionMonthMethodResponse {
	return &model.TransactionMonthMethodResponse{
		Month:             s.Month,
		PaymentMethod:     s.PaymentMethod,
		TotalTransactions: int32(s.TotalTransactions),
		TotalAmount:       int32(s.TotalAmount),
	}
}

func (s *transactionResponseMapper) mapResponsesTransactionMonthlyMethod(Transactions []*response.TransactionMonthMethodResponse) []*model.TransactionMonthMethodResponse {
	var responses []*model.TransactionMonthMethodResponse

	for _, Transaction := range Transactions {
		responses = append(responses, s.mapResponseTransactionMonthlyMethod(Transaction))
	}

	return responses
}

func (t *transactionResponseMapper) mapResponseTransactionYearlyMethod(s *response.TransactionYearMethodResponse) *model.TransactionYearMethodResponse {
	return &model.TransactionYearMethodResponse{
		Year:              s.Year,
		PaymentMethod:     s.PaymentMethod,
		TotalTransactions: int32(s.TotalTransactions),
		TotalAmount:       int32(s.TotalAmount),
	}
}

func (s *transactionResponseMapper) mapResponsesTransactionYearlyMethod(Transactions []*response.TransactionYearMethodResponse) []*model.TransactionYearMethodResponse {
	var responses []*model.TransactionYearMethodResponse

	for _, Transaction := range Transactions {
		responses = append(responses, s.mapResponseTransactionYearlyMethod(Transaction))
	}

	return responses
}

func (t *transactionResponseMapper) mapResponseTransactionMonthlyAmount(s *response.TransactionMonthAmountResponse) *model.TransactionMonthAmountResponse {
	return &model.TransactionMonthAmountResponse{
		Month:       s.Month,
		TotalAmount: int32(s.TotalAmount),
	}
}

func (s *transactionResponseMapper) mapResponsesTransactionMonthlyAmount(Transactions []*response.TransactionMonthAmountResponse) []*model.TransactionMonthAmountResponse {
	var responses []*model.TransactionMonthAmountResponse

	for _, Transaction := range Transactions {
		responses = append(responses, s.mapResponseTransactionMonthlyAmount(Transaction))
	}

	return responses
}

func (t *transactionResponseMapper) mapResponseTransactionYearlyAmount(s *response.TransactionYearlyAmountResponse) *model.TransactionYearlyAmountResponse {
	return &model.TransactionYearlyAmountResponse{
		Year:        s.Year,
		TotalAmount: int32(s.TotalAmount),
	}
}

func (s *transactionResponseMapper) mapResponsesTransactionYearlyAmount(Transactions []*response.TransactionYearlyAmountResponse) []*model.TransactionYearlyAmountResponse {
	var responses []*model.TransactionYearlyAmountResponse

	for _, Transaction := range Transactions {
		responses = append(responses, s.mapResponseTransactionYearlyAmount(Transaction))
	}

	return responses
}
