package graphql

import (
	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/response"
	"github.com/MamangRust/paymentgatewaygraphql/internal/graph/model"
)

type withdrawResponseMapper struct {
}

func NewWithdrawResponseMapper() *withdrawResponseMapper {
	return &withdrawResponseMapper{}
}

func (t *withdrawResponseMapper) ToGraphqlWithdrawAll(status, message string) *model.APIResponseWithdrawAll {
	return &model.APIResponseWithdrawAll{
		Status:  status,
		Message: message,
	}
}

func (t *withdrawResponseMapper) ToGraphqlWithdrawDelete(status, message string) *model.APIResponseWithdrawDelete {
	return &model.APIResponseWithdrawDelete{
		Status:  status,
		Message: message,
	}
}

func (t *withdrawResponseMapper) ToGraphqlResponseWithdraw(status, message string, data *response.WithdrawResponse) *model.APIResponseWithdraw {
	return &model.APIResponseWithdraw{
		Status:  status,
		Message: message,
		Data:    t.mapResponseWithdraw(data),
	}
}

func (t *withdrawResponseMapper) ToGraphqlResponseWithdraws(status, message string, data []*response.WithdrawResponse) *model.APIResponsesWithdraw {
	return &model.APIResponsesWithdraw{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesWithdraw(data),
	}
}

func (t *withdrawResponseMapper) ToGraphqlResponseWithdrawDeleteAt(status, message string, data *response.WithdrawResponseDeleteAt) *model.APIResponseWithdrawDeleteAt {
	return &model.APIResponseWithdrawDeleteAt{
		Status:  status,
		Message: message,
		Data:    t.mapResponseWithdrawDeleteAt(data),
	}
}

func (t *withdrawResponseMapper) ToGraphqlResponsePaginationWithdraw(status, message string, data []*response.WithdrawResponse, pagination *response.PaginationMeta) *model.APIResponsePaginationWithdraw {
	return &model.APIResponsePaginationWithdraw{
		Status:     status,
		Message:    message,
		Data:       t.mapResponsesWithdraw(data),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (t *withdrawResponseMapper) ToGraphqlResponsePaginationWithdrawDeleteAt(status, message string, data []*response.WithdrawResponseDeleteAt, pagination *response.PaginationMeta) *model.APIResponsePaginationWithdrawDeleteAt {
	return &model.APIResponsePaginationWithdrawDeleteAt{
		Status:     status,
		Message:    message,
		Data:       t.mapResponsesWithdrawDeleteAt(data),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (t *withdrawResponseMapper) ToGraphqlResponseWithdrawMonthAmount(status, message string, data []*response.WithdrawMonthlyAmountResponse) *model.APIResponseWithdrawMonthAmount {
	return &model.APIResponseWithdrawMonthAmount{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesWithdrawMonthAmount(data),
	}
}

func (t *withdrawResponseMapper) ToGraphqlResponseWithdrawYearAmount(status, message string, data []*response.WithdrawYearlyAmountResponse) *model.APIResponseWithdrawYearAmount {
	return &model.APIResponseWithdrawYearAmount{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesWithdrawYearAmount(data),
	}
}

func (t *withdrawResponseMapper) ToGraphqlResponseWithdrawMonthStatusSuccess(status, message string, data []*response.WithdrawResponseMonthStatusSuccess) *model.APIResponseWithdrawMonthStatusSuccess {
	return &model.APIResponseWithdrawMonthStatusSuccess{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesMonthStatusSuccess(data),
	}
}

func (t *withdrawResponseMapper) ToGraphqlResponseWithdrawYearStatusSuccess(status, message string, data []*response.WithdrawResponseYearStatusSuccess) *model.APIResponseWithdrawYearStatusSuccess {
	return &model.APIResponseWithdrawYearStatusSuccess{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesYearStatusSuccess(data),
	}
}

func (t *withdrawResponseMapper) ToGraphqlResponseWithdrawMonthStatusFailed(status, message string, data []*response.WithdrawResponseMonthStatusFailed) *model.APIResponseWithdrawMonthStatusFailed {
	return &model.APIResponseWithdrawMonthStatusFailed{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesMonthStatusFailed(data),
	}
}

func (t *withdrawResponseMapper) ToGraphqlResponseWithdrawYearStatusFailed(status, message string, data []*response.WithdrawResponseYearStatusFailed) *model.APIResponseWithdrawYearStatusFailed {
	return &model.APIResponseWithdrawYearStatusFailed{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesYearStatusFailed(data),
	}
}

func (t *withdrawResponseMapper) mapResponseWithdraw(Withdraw *response.WithdrawResponse) *model.WithdrawResponse {
	return &model.WithdrawResponse{
		ID:             int32(Withdraw.ID),
		WithdrawNo:     Withdraw.WithdrawNo,
		CardNumber:     Withdraw.CardNumber,
		WithdrawAmount: int32(Withdraw.WithdrawAmount),
		WithdrawTime:   Withdraw.WithdrawTime,
		CreatedAt:      Withdraw.CreatedAt,
		UpdatedAt:      Withdraw.UpdatedAt,
	}
}

func (t *withdrawResponseMapper) mapResponsesWithdraw(Withdraws []*response.WithdrawResponse) []*model.WithdrawResponse {
	var responses []*model.WithdrawResponse

	for _, Withdraw := range Withdraws {
		responses = append(responses, t.mapResponseWithdraw(Withdraw))
	}

	return responses
}

func (t *withdrawResponseMapper) mapResponseWithdrawDeleteAt(Withdraw *response.WithdrawResponseDeleteAt) *model.WithdrawResponseDeleteAt {
	return &model.WithdrawResponseDeleteAt{
		ID:             int32(Withdraw.ID),
		WithdrawNo:     Withdraw.WithdrawNo,
		CardNumber:     Withdraw.CardNumber,
		WithdrawAmount: int32(Withdraw.WithdrawAmount),
		WithdrawTime:   Withdraw.WithdrawTime,
		CreatedAt:      Withdraw.CreatedAt,
		UpdatedAt:      Withdraw.UpdatedAt,
		DeletedAt:      Withdraw.DeletedAt,
	}
}

func (t *withdrawResponseMapper) mapResponsesWithdrawDeleteAt(Withdraws []*response.WithdrawResponseDeleteAt) []*model.WithdrawResponseDeleteAt {
	var responses []*model.WithdrawResponseDeleteAt

	for _, Withdraw := range Withdraws {
		responses = append(responses, t.mapResponseWithdrawDeleteAt(Withdraw))
	}

	return responses
}

func (t *withdrawResponseMapper) mapResponseMonthStatusSuccess(data *response.WithdrawResponseMonthStatusSuccess) *model.WithdrawMonthStatusSuccessResponse {
	return &model.WithdrawMonthStatusSuccessResponse{
		Year:         data.Year,
		Month:        data.Month,
		TotalSuccess: int32(data.TotalSuccess),
		TotalAmount:  int32(data.TotalAmount),
	}
}

func (t *withdrawResponseMapper) mapResponsesMonthStatusSuccess(Withdraws []*response.WithdrawResponseMonthStatusSuccess) []*model.WithdrawMonthStatusSuccessResponse {
	var responses []*model.WithdrawMonthStatusSuccessResponse

	for _, Withdraw := range Withdraws {
		responses = append(responses, t.mapResponseMonthStatusSuccess(Withdraw))
	}

	return responses
}

func (t *withdrawResponseMapper) mapResponseYearStatusSuccess(data *response.WithdrawResponseYearStatusSuccess) *model.WithdrawYearStatusSuccessResponse {
	return &model.WithdrawYearStatusSuccessResponse{
		Year:         data.Year,
		TotalSuccess: int32(data.TotalSuccess),
		TotalAmount:  int32(data.TotalAmount),
	}
}

func (t *withdrawResponseMapper) mapResponsesYearStatusSuccess(Withdraws []*response.WithdrawResponseYearStatusSuccess) []*model.WithdrawYearStatusSuccessResponse {
	var responses []*model.WithdrawYearStatusSuccessResponse

	for _, Withdraw := range Withdraws {
		responses = append(responses, t.mapResponseYearStatusSuccess(Withdraw))
	}

	return responses
}

func (t *withdrawResponseMapper) mapResponseMonthStatusFailed(data *response.WithdrawResponseMonthStatusFailed) *model.WithdrawMonthStatusFailedResponse {
	return &model.WithdrawMonthStatusFailedResponse{
		Year:        data.Year,
		Month:       data.Month,
		TotalFailed: int32(data.TotalFailed),
		TotalAmount: int32(data.TotalAmount),
	}
}

func (t *withdrawResponseMapper) mapResponsesMonthStatusFailed(Withdraws []*response.WithdrawResponseMonthStatusFailed) []*model.WithdrawMonthStatusFailedResponse {
	var responses []*model.WithdrawMonthStatusFailedResponse

	for _, Withdraw := range Withdraws {
		responses = append(responses, t.mapResponseMonthStatusFailed(Withdraw))
	}

	return responses
}

func (t *withdrawResponseMapper) mapResponseYearStatusFailed(data *response.WithdrawResponseYearStatusFailed) *model.WithdrawYearStatusFailedResponse {
	return &model.WithdrawYearStatusFailedResponse{
		Year:        data.Year,
		TotalFailed: int32(data.TotalFailed),
		TotalAmount: int32(data.TotalAmount),
	}
}

func (t *withdrawResponseMapper) mapResponsesYearStatusFailed(Withdraws []*response.WithdrawResponseYearStatusFailed) []*model.WithdrawYearStatusFailedResponse {
	var responses []*model.WithdrawYearStatusFailedResponse

	for _, Withdraw := range Withdraws {
		responses = append(responses, t.mapResponseYearStatusFailed(Withdraw))
	}

	return responses
}

func (t *withdrawResponseMapper) mapResponseWithdrawMonthAmount(Withdraw *response.WithdrawMonthlyAmountResponse) *model.WithdrawMonthlyAmountResponse {
	return &model.WithdrawMonthlyAmountResponse{
		Month:       Withdraw.Month,
		TotalAmount: int32(Withdraw.TotalAmount),
	}
}

func (t *withdrawResponseMapper) mapResponsesWithdrawMonthAmount(Withdraws []*response.WithdrawMonthlyAmountResponse) []*model.WithdrawMonthlyAmountResponse {
	var responses []*model.WithdrawMonthlyAmountResponse

	for _, Withdraw := range Withdraws {
		responses = append(responses, t.mapResponseWithdrawMonthAmount(Withdraw))
	}

	return responses
}

func (t *withdrawResponseMapper) mapResponseWithdrawYearAmount(Withdraw *response.WithdrawYearlyAmountResponse) *model.WithdrawYearlyAmountResponse {
	return &model.WithdrawYearlyAmountResponse{
		Year:        Withdraw.Year,
		TotalAmount: int32(Withdraw.TotalAmount),
	}
}

func (t *withdrawResponseMapper) mapResponsesWithdrawYearAmount(Withdraws []*response.WithdrawYearlyAmountResponse) []*model.WithdrawYearlyAmountResponse {
	var responses []*model.WithdrawYearlyAmountResponse

	for _, Withdraw := range Withdraws {
		responses = append(responses, t.mapResponseWithdrawYearAmount(Withdraw))
	}

	return responses
}
