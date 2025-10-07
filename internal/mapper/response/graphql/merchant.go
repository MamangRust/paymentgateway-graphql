package graphql

import (
	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/response"
	"github.com/MamangRust/paymentgatewaygraphql/internal/graph/model"
)

type merchantResponse struct{}

func NewMerchantResponseMapper() *merchantResponse {
	return &merchantResponse{}
}

func (m *merchantResponse) ToGraphqlResponseMerchant(status, message string, merchant *response.MerchantResponse) *model.APIResponseMerchant {
	return &model.APIResponseMerchant{
		Status:  status,
		Message: message,
		Data:    m.mapMerchantResponse(merchant),
	}
}

func (m *merchantResponse) ToGraphqlResponsePaginationMerchant(status, message string, merchant []*response.MerchantResponse, pagination *response.PaginationMeta) *model.APIResponseMerchantPagination {
	return &model.APIResponseMerchantPagination{
		Status:     status,
		Message:    message,
		Data:       m.mapMerchantResponses(merchant),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (m *merchantResponse) ToGraphqlResponsesMerchant(status, message string, merchant []*response.MerchantResponse) *model.APIResponsesMerchant {
	return &model.APIResponsesMerchant{
		Status:  status,
		Message: message,
		Data:    m.mapMerchantResponses(merchant),
	}
}

func (m *merchantResponse) ToGraphqlResponseMerchantDeleteAt(status, message string, merchant *response.MerchantResponseDeleteAt) *model.APIResponseMerchantDeleteAt {
	return &model.APIResponseMerchantDeleteAt{
		Status:  status,
		Message: message,
		Data:    m.mapMerchantResponseDeleteAt(merchant),
	}
}

func (m *merchantResponse) ToGraphqlResponsePaginationMerchantDeleteAt(status, message string, merchant []*response.MerchantResponseDeleteAt, pagination *response.PaginationMeta) *model.APIResponseMerchantDeleteAtPagination {
	return &model.APIResponseMerchantDeleteAtPagination{
		Status:     status,
		Message:    message,
		Data:       m.mapMerchantResponsesDeleteAt(merchant),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (m *merchantResponse) ToGraphqlResponsePaginationTransaction(status, message string, merchant []*response.MerchantTransactionResponse, pagination *response.PaginationMeta) *model.APIResponseMerchantTransactionPagination {

	return &model.APIResponseMerchantTransactionPagination{
		Status:     status,
		Message:    message,
		Data:       m.mapMerchantTransactionResponses(merchant),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (m *merchantResponse) ToGraphqlMonthlyPaymentMethods(status, message string, merchant []*response.MerchantResponseMonthlyPaymentMethod) *model.APIResponseMerchantMonthlyPaymentMethod {
	return &model.APIResponseMerchantMonthlyPaymentMethod{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesMonthlyPaymentMethod(merchant),
	}
}

func (m *merchantResponse) ToGraphqlYearlyPaymentMethods(status, message string, merchant []*response.MerchantResponseYearlyPaymentMethod) *model.APIResponseMerchantYearlyPaymentMethod {
	return &model.APIResponseMerchantYearlyPaymentMethod{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesYearlyPaymentMethod(merchant),
	}
}

func (m *merchantResponse) ToGraphqlMonthlyAmounts(status, message string, ms []*response.MerchantResponseMonthlyAmount) *model.APIResponseMerchantMonthlyAmount {
	return &model.APIResponseMerchantMonthlyAmount{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesMonthlyAmount(ms),
	}
}

func (m *merchantResponse) ToGraphqlYearlyAmounts(status, message string, ms []*response.MerchantResponseYearlyAmount) *model.APIResponseMerchantYearlyAmount {
	return &model.APIResponseMerchantYearlyAmount{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesYearlyAmount(ms),
	}
}

func (m *merchantResponse) ToGraphqlMonthlyTotalAmounts(status, message string, ms []*response.MerchantResponseMonthlyTotalAmount) *model.APIResponseMerchantMonthlyTotalAmount {
	return &model.APIResponseMerchantMonthlyTotalAmount{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesMonthlyTotalAmount(ms),
	}
}

func (m *merchantResponse) ToGraphqlYearlyTotalAmounts(status, message string, ms []*response.MerchantResponseYearlyTotalAmount) *model.APIResponseMerchantYearlyTotalAmount {
	return &model.APIResponseMerchantYearlyTotalAmount{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesYearlyTotalAmount(ms),
	}
}

func (s *merchantResponse) ToGraphqlMerchantDeleteAll(status, message string) *model.APIResponseMerchantDelete {
	return &model.APIResponseMerchantDelete{
		Status:  status,
		Message: message,
	}
}

func (s *merchantResponse) ToGraphqlMerchantAll(status, message string) *model.APIResponseMerchantAll {
	return &model.APIResponseMerchantAll{
		Status:  status,
		Message: message,
	}
}

func (m *merchantResponse) mapMerchantResponse(merchant *response.MerchantResponse) *model.MerchantResponse {
	return &model.MerchantResponse{
		ID:        int32(merchant.ID),
		Name:      merchant.Name,
		UserID:    int32(merchant.UserID),
		Status:    merchant.Status,
		APIKey:    merchant.ApiKey,
		CreatedAt: merchant.CreatedAt,
		UpdatedAt: merchant.UpdatedAt,
	}
}

func (m *merchantResponse) mapMerchantResponses(merchants []*response.MerchantResponse) []*model.MerchantResponse {
	var responseMerchants []*model.MerchantResponse

	for _, merchant := range merchants {
		responseMerchants = append(responseMerchants, m.mapMerchantResponse(merchant))
	}

	return responseMerchants
}

func (m *merchantResponse) mapMerchantResponseDeleteAt(merchant *response.MerchantResponseDeleteAt) *model.MerchantResponseDeleteAt {
	return &model.MerchantResponseDeleteAt{
		ID:        int32(merchant.ID),
		Name:      merchant.Name,
		UserID:    int32(merchant.UserID),
		Status:    merchant.Status,
		APIKey:    merchant.ApiKey,
		CreatedAt: merchant.CreatedAt,
		UpdatedAt: merchant.UpdatedAt,
		DeletedAt: *merchant.DeletedAt,
	}
}

func (m *merchantResponse) mapMerchantResponsesDeleteAt(merchants []*response.MerchantResponseDeleteAt) []*model.MerchantResponseDeleteAt {
	var responseMerchants []*model.MerchantResponseDeleteAt

	for _, merchant := range merchants {
		responseMerchants = append(responseMerchants, m.mapMerchantResponseDeleteAt(merchant))
	}

	return responseMerchants
}

func (m *merchantResponse) mapMerchantTransactionResponse(merchant *response.MerchantTransactionResponse) *model.MerchantTransactionResponse {
	return &model.MerchantTransactionResponse{
		ID:              int32(merchant.ID),
		CardNumber:      merchant.CardNumber,
		Amount:          merchant.Amount,
		PaymentMethod:   merchant.PaymentMethod,
		MerchantID:      merchant.MerchantID,
		MerchantName:    merchant.MerchantName,
		TransactionTime: merchant.TransactionTime,
		CreatedAt:       merchant.CreatedAt,
		UpdatedAt:       merchant.UpdatedAt,
	}
}

func (m *merchantResponse) mapMerchantTransactionResponses(r []*response.MerchantTransactionResponse) []*model.MerchantTransactionResponse {
	var responseMerchants []*model.MerchantTransactionResponse
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapMerchantTransactionResponse(merchant))
	}

	return responseMerchants
}

func (m *merchantResponse) mapResponseMonthlyPaymentMethod(ms *response.MerchantResponseMonthlyPaymentMethod) *model.MerchantMonthlyPaymentMethodResponse {
	return &model.MerchantMonthlyPaymentMethodResponse{
		Month:         ms.Month,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int32(ms.TotalAmount),
	}
}

func (m *merchantResponse) mapResponsesMonthlyPaymentMethod(r []*response.MerchantResponseMonthlyPaymentMethod) []*model.MerchantMonthlyPaymentMethodResponse {
	var responseMerchants []*model.MerchantMonthlyPaymentMethodResponse
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapResponseMonthlyPaymentMethod(merchant))
	}

	return responseMerchants
}

func (m *merchantResponse) mapResponseYearlyPaymentMethod(ms *response.MerchantResponseYearlyPaymentMethod) *model.MerchantYearlyPaymentMethodResponse {
	return &model.MerchantYearlyPaymentMethodResponse{
		Year:          ms.Year,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int32(ms.TotalAmount),
	}
}

func (m *merchantResponse) mapResponsesYearlyPaymentMethod(r []*response.MerchantResponseYearlyPaymentMethod) []*model.MerchantYearlyPaymentMethodResponse {
	var responseMerchants []*model.MerchantYearlyPaymentMethodResponse
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapResponseYearlyPaymentMethod(merchant))
	}

	return responseMerchants
}

func (m *merchantResponse) mapResponseMonthlyAmount(ms *response.MerchantResponseMonthlyAmount) *model.MerchantMonthlyAmountResponse {
	return &model.MerchantMonthlyAmountResponse{
		Month:       ms.Month,
		TotalAmount: int32(ms.TotalAmount),
	}
}

func (m *merchantResponse) mapResponsesMonthlyAmount(r []*response.MerchantResponseMonthlyAmount) []*model.MerchantMonthlyAmountResponse {
	var responseMerchants []*model.MerchantMonthlyAmountResponse
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapResponseMonthlyAmount(merchant))
	}

	return responseMerchants
}

func (m *merchantResponse) mapResponseYearlyAmount(ms *response.MerchantResponseYearlyAmount) *model.MerchantYearlyAmountResponse {
	return &model.MerchantYearlyAmountResponse{
		Year:        ms.Year,
		TotalAmount: int32(ms.TotalAmount),
	}
}

func (m *merchantResponse) mapResponsesYearlyAmount(r []*response.MerchantResponseYearlyAmount) []*model.MerchantYearlyAmountResponse {
	var responseMerchants []*model.MerchantYearlyAmountResponse
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapResponseYearlyAmount(merchant))
	}

	return responseMerchants
}

func (m *merchantResponse) mapResponseMonthlyTotalAmount(ms *response.MerchantResponseMonthlyTotalAmount) *model.MerchantMonthlyTotalAmountResponse {
	return &model.MerchantMonthlyTotalAmountResponse{
		Month:       ms.Month,
		Year:        ms.Year,
		TotalAmount: int32(ms.TotalAmount),
	}
}

func (m *merchantResponse) mapResponsesMonthlyTotalAmount(r []*response.MerchantResponseMonthlyTotalAmount) []*model.MerchantMonthlyTotalAmountResponse {
	var responseMerchants []*model.MerchantMonthlyTotalAmountResponse
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapResponseMonthlyTotalAmount(merchant))
	}

	return responseMerchants
}

func (m *merchantResponse) mapResponseYearlyTotalAmount(ms *response.MerchantResponseYearlyTotalAmount) *model.MerchantYearlyTotalAmountResponse {
	return &model.MerchantYearlyTotalAmountResponse{
		Year:        ms.Year,
		TotalAmount: int32(ms.TotalAmount),
	}
}

func (m *merchantResponse) mapResponsesYearlyTotalAmount(r []*response.MerchantResponseYearlyTotalAmount) []*model.MerchantYearlyTotalAmountResponse {
	var responseMerchants []*model.MerchantYearlyTotalAmountResponse
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapResponseYearlyTotalAmount(merchant))
	}

	return responseMerchants
}
