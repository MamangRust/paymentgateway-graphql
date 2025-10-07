package graphql

import (
	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/response"
	"github.com/MamangRust/paymentgatewaygraphql/internal/graph/model"
)

type topupResponseMapper struct {
}

func NewTopupResponseMapper() *topupResponseMapper {
	return &topupResponseMapper{}
}

func (t *topupResponseMapper) ToGraphqlResponseTopup(status, message string, data *response.TopupResponse) *model.APIResponseTopup {
	return &model.APIResponseTopup{
		Status:  status,
		Message: message,
		Data:    t.mapResponseTopup(data),
	}
}

func (t *topupResponseMapper) ToGraphqlResponseTopupDeleteAt(status, message string, data *response.TopupResponseDeleteAt) *model.APIResponseTopupDeleteAt {
	return &model.APIResponseTopupDeleteAt{
		Status:  status,
		Message: message,
		Data:    t.mapResponseTopupDeleteAt(data),
	}
}

func (t *topupResponseMapper) ToGraphqlTopupAll(status, message string) *model.APIResponseTopupAll {
	return &model.APIResponseTopupAll{
		Status:  status,
		Message: message,
	}
}

func (t *topupResponseMapper) ToGraphqlTopupDelete(status, message string) *model.APIResponseTopupDelete {
	return &model.APIResponseTopupDelete{
		Status:  status,
		Message: message,
	}
}

func (t *topupResponseMapper) ToGraphqlResponsePaginationTopup(status, message string, data []*response.TopupResponse, pagination *response.PaginationMeta) *model.APIResponsePaginationTopup {
	return &model.APIResponsePaginationTopup{
		Status:     status,
		Message:    message,
		Data:       t.mapResponsesTopup(data),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (t *topupResponseMapper) ToGraphqlResponsePaginationTopupDeleteAt(status, message string, data []*response.TopupResponseDeleteAt, pagination *response.PaginationMeta) *model.APIResponsePaginationTopupDeleteAt {
	return &model.APIResponsePaginationTopupDeleteAt{
		Status:     status,
		Message:    message,
		Data:       t.mapResponsesTopupDeleteAt(data),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (t *topupResponseMapper) ToGraphqlResponseTopupMonthStatusSuccess(status, message string, data []*response.TopupResponseMonthStatusSuccess) *model.APIResponseTopupMonthStatusSuccess {
	return &model.APIResponseTopupMonthStatusSuccess{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTopupMonthStatusSuccess(data),
	}
}

func (t *topupResponseMapper) ToGraphqlResponseTopupYearStatusSuccess(status, message string, data []*response.TopupResponseYearStatusSuccess) *model.APIResponseTopupYearStatusSuccess {
	return &model.APIResponseTopupYearStatusSuccess{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTopupYearStatusSuccess(data),
	}
}

func (t *topupResponseMapper) ToGraphqlResponseTopupMonthStatusFailed(status, message string, data []*response.TopupResponseMonthStatusFailed) *model.APIResponseTopupMonthStatusFailed {
	return &model.APIResponseTopupMonthStatusFailed{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTopupMonthStatusFailed(data),
	}
}

func (t *topupResponseMapper) ToGraphqlResponseTopupYearStatusFailed(status, message string, data []*response.TopupResponseYearStatusFailed) *model.APIResponseTopupYearStatusFailed {
	return &model.APIResponseTopupYearStatusFailed{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTopupYearStatusFailed(data),
	}
}

func (t *topupResponseMapper) ToGraphqlResponseTopupMonthMethod(status, message string, data []*response.TopupMonthMethodResponse) *model.APIResponseTopupMonthMethod {
	return &model.APIResponseTopupMonthMethod{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTopupMonthlyMethod(data),
	}
}

func (t *topupResponseMapper) ToGraphqlResponseTopupYearMethod(status, message string, data []*response.TopupYearlyMethodResponse) *model.APIResponseTopupYearMethod {
	return &model.APIResponseTopupYearMethod{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTopupYearlyMethod(data),
	}
}

func (t *topupResponseMapper) ToGraphqlResponseTopupMonthAmount(status, message string, data []*response.TopupMonthAmountResponse) *model.APIResponseTopupMonthAmount {
	return &model.APIResponseTopupMonthAmount{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTopupMonthlyAmount(data),
	}
}

func (t *topupResponseMapper) ToGraphqlResponseTopupYearAmount(status, message string, data []*response.TopupYearlyAmountResponse) *model.APIResponseTopupYearAmount {
	return &model.APIResponseTopupYearAmount{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTopupYearlyAmount(data),
	}
}

func (t *topupResponseMapper) mapResponseTopup(topup *response.TopupResponse) *model.TopupResponse {
	return &model.TopupResponse{
		ID:          int32(topup.ID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: int32(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   &topup.TopupTime,
		CreatedAt:   topup.CreatedAt,
		UpdatedAt:   topup.UpdatedAt,
	}
}

func (t *topupResponseMapper) mapResponsesTopup(topups []*response.TopupResponse) []*model.TopupResponse {
	var responses []*model.TopupResponse

	for _, topup := range topups {
		responses = append(responses, t.mapResponseTopup(topup))
	}

	return responses
}

func (t *topupResponseMapper) mapResponseTopupDeleteAt(topup *response.TopupResponseDeleteAt) *model.TopupResponseDeleteAt {
	return &model.TopupResponseDeleteAt{
		ID:          int32(topup.ID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: int32(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   &topup.TopupTime,
		CreatedAt:   topup.CreatedAt,
		UpdatedAt:   topup.UpdatedAt,
		DeletedAt:   topup.DeletedAt,
	}
}

func (t *topupResponseMapper) mapResponsesTopupDeleteAt(topups []*response.TopupResponseDeleteAt) []*model.TopupResponseDeleteAt {
	var responses []*model.TopupResponseDeleteAt

	for _, topup := range topups {
		responses = append(responses, t.mapResponseTopupDeleteAt(topup))
	}

	return responses
}

func (t *topupResponseMapper) mapResponseTopupMonthStatusSuccess(s *response.TopupResponseMonthStatusSuccess) *model.TopupMonthStatusSuccessResponse {
	return &model.TopupMonthStatusSuccessResponse{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

func (t *topupResponseMapper) mapResponsesTopupMonthStatusSuccess(topups []*response.TopupResponseMonthStatusSuccess) []*model.TopupMonthStatusSuccessResponse {
	var responses []*model.TopupMonthStatusSuccessResponse

	for _, topup := range topups {
		responses = append(responses, t.mapResponseTopupMonthStatusSuccess(topup))
	}

	return responses
}

func (t *topupResponseMapper) mapResponseTopupMonthStatusFailed(s *response.TopupResponseMonthStatusFailed) *model.TopupMonthStatusFailedResponse {
	return &model.TopupMonthStatusFailedResponse{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

func (t *topupResponseMapper) mapResponsesTopupMonthStatusFailed(topups []*response.TopupResponseMonthStatusFailed) []*model.TopupMonthStatusFailedResponse {
	var responses []*model.TopupMonthStatusFailedResponse

	for _, topup := range topups {
		responses = append(responses, t.mapResponseTopupMonthStatusFailed(topup))
	}

	return responses
}

func (t *topupResponseMapper) mapResponseTopupYearStatusSuccess(s *response.TopupResponseYearStatusSuccess) *model.TopupYearStatusSuccessResponse {
	return &model.TopupYearStatusSuccessResponse{
		Year:         s.Year,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

func (t *topupResponseMapper) mapResponsesTopupYearStatusSuccess(topups []*response.TopupResponseYearStatusSuccess) []*model.TopupYearStatusSuccessResponse {
	var responses []*model.TopupYearStatusSuccessResponse

	for _, topup := range topups {
		responses = append(responses, t.mapResponseTopupYearStatusSuccess(topup))
	}

	return responses
}

func (t *topupResponseMapper) mapResponseTopupYearStatusFailed(s *response.TopupResponseYearStatusFailed) *model.TopupYearStatusFailedResponse {
	return &model.TopupYearStatusFailedResponse{
		Year:        s.Year,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

func (t *topupResponseMapper) mapResponsesTopupYearStatusFailed(topups []*response.TopupResponseYearStatusFailed) []*model.TopupYearStatusFailedResponse {
	var responses []*model.TopupYearStatusFailedResponse

	for _, topup := range topups {
		responses = append(responses, t.mapResponseTopupYearStatusFailed(topup))
	}

	return responses
}

func (t *topupResponseMapper) mapResponseTopupMonthlyMethod(s *response.TopupMonthMethodResponse) *model.TopupMonthMethodResponse {
	return &model.TopupMonthMethodResponse{
		Month:       s.Month,
		TopupMethod: s.TopupMethod,
		TotalTopups: int32(s.TotalTopups),
		TotalAmount: int32(s.TotalAmount),
	}
}

func (s *topupResponseMapper) mapResponsesTopupMonthlyMethod(topups []*response.TopupMonthMethodResponse) []*model.TopupMonthMethodResponse {
	var responses []*model.TopupMonthMethodResponse

	for _, topup := range topups {
		responses = append(responses, s.mapResponseTopupMonthlyMethod(topup))
	}

	return responses
}

func (t *topupResponseMapper) mapResponseTopupYearlyMethod(s *response.TopupYearlyMethodResponse) *model.TopupYearMethodResponse {
	return &model.TopupYearMethodResponse{
		Year:        s.Year,
		TopupMethod: s.TopupMethod,
		TotalTopups: int32(s.TotalTopups),
		TotalAmount: int32(s.TotalAmount),
	}
}

func (s *topupResponseMapper) mapResponsesTopupYearlyMethod(topups []*response.TopupYearlyMethodResponse) []*model.TopupYearMethodResponse {
	var responses []*model.TopupYearMethodResponse

	for _, topup := range topups {
		responses = append(responses, s.mapResponseTopupYearlyMethod(topup))
	}

	return responses
}

func (t *topupResponseMapper) mapResponseTopupMonthlyAmount(s *response.TopupMonthAmountResponse) *model.TopupMonthAmountResponse {
	return &model.TopupMonthAmountResponse{
		Month:       s.Month,
		TotalAmount: int32(s.TotalAmount),
	}
}

func (s *topupResponseMapper) mapResponsesTopupMonthlyAmount(topups []*response.TopupMonthAmountResponse) []*model.TopupMonthAmountResponse {
	var responses []*model.TopupMonthAmountResponse

	for _, topup := range topups {
		responses = append(responses, s.mapResponseTopupMonthlyAmount(topup))
	}

	return responses
}

func (t *topupResponseMapper) mapResponseTopupYearlyAmount(s *response.TopupYearlyAmountResponse) *model.TopupYearAmountResponse {
	return &model.TopupYearAmountResponse{
		Year:        s.Year,
		TotalAmount: int32(s.TotalAmount),
	}
}

func (s *topupResponseMapper) mapResponsesTopupYearlyAmount(topups []*response.TopupYearlyAmountResponse) []*model.TopupYearAmountResponse {
	var responses []*model.TopupYearAmountResponse

	for _, topup := range topups {
		responses = append(responses, s.mapResponseTopupYearlyAmount(topup))
	}

	return responses
}
