package graphql

import (
	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/response"
	"github.com/MamangRust/paymentgatewaygraphql/internal/graph/model"
)

type transferResponseMapper struct {
}

func NewTransferResponseMapper() *transferResponseMapper {
	return &transferResponseMapper{}
}

func (t *transferResponseMapper) ToGraphqlTransferAll(status, message string) *model.APIResponseTransferAll {
	return &model.APIResponseTransferAll{
		Status:  status,
		Message: message,
	}
}

func (t *transferResponseMapper) ToGraphqlTransferDelete(status, message string) *model.APIResponseTransferDelete {
	return &model.APIResponseTransferDelete{
		Status:  status,
		Message: message,
	}
}

func (t *transferResponseMapper) ToGraphqlResponseTransfer(status, message string, data *response.TransferResponse) *model.APIResponseTransfer {
	return &model.APIResponseTransfer{
		Status:  status,
		Message: message,
		Data:    t.mapResponseTransfer(data),
	}
}

func (t *transferResponseMapper) ToGraphqlResponseTransfers(status, message string, data []*response.TransferResponse) *model.APIResponseTransfers {
	return &model.APIResponseTransfers{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTransfer(data),
	}
}

func (t *transferResponseMapper) ToGraphqlResponseTransferDeleteAt(status, message string, data *response.TransferResponseDeleteAt) *model.APIResponseTransferDeleteAt {
	return &model.APIResponseTransferDeleteAt{
		Status:  status,
		Message: message,
		Data:    t.mapResponseTransferDeleteAt(data),
	}
}

func (t *transferResponseMapper) ToGraphqlResponsePaginationTransfer(status, message string, data []*response.TransferResponse, pagination *response.PaginationMeta) *model.APIResponsePaginationTransfer {
	return &model.APIResponsePaginationTransfer{
		Status:     status,
		Message:    message,
		Data:       t.mapResponsesTransfer(data),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (t *transferResponseMapper) ToGraphqlResponsePaginationTransferDeleteAt(status, message string, data []*response.TransferResponseDeleteAt, pagination *response.PaginationMeta) *model.APIResponsePaginationTransferDeleteAt {
	return &model.APIResponsePaginationTransferDeleteAt{
		Status:     status,
		Message:    message,
		Data:       t.mapResponsesTransferDeleteAt(data),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (t *transferResponseMapper) ToGraphqlResponseTransferMonthAmount(status, message string, data []*response.TransferMonthAmountResponse) *model.APIResponseTransferMonthAmount {
	return &model.APIResponseTransferMonthAmount{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTransferMonthAmount(data),
	}
}

func (t *transferResponseMapper) ToGraphqlResponseTransferYearAmount(status, message string, data []*response.TransferYearAmountResponse) *model.APIResponseTransferYearAmount {
	return &model.APIResponseTransferYearAmount{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTransferYearAmount(data),
	}
}

func (t *transferResponseMapper) ToGraphqlResponseTransferMonthStatusSuccess(status, message string, data []*response.TransferResponseMonthStatusSuccess) *model.APIResponseTransferMonthStatusSuccess {
	return &model.APIResponseTransferMonthStatusSuccess{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesMonthStatusSuccess(data),
	}
}

func (t *transferResponseMapper) ToGraphqlResponseTransferYearStatusSuccess(status, message string, data []*response.TransferResponseYearStatusSuccess) *model.APIResponseTransferYearStatusSuccess {
	return &model.APIResponseTransferYearStatusSuccess{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesYearStatusSuccess(data),
	}
}

func (t *transferResponseMapper) ToGraphqlResponseTransferMonthStatusFailed(status, message string, data []*response.TransferResponseMonthStatusFailed) *model.APIResponseTransferMonthStatusFailed {
	return &model.APIResponseTransferMonthStatusFailed{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesMonthStatusFailed(data),
	}
}

func (t *transferResponseMapper) ToGraphqlResponseTransferYearStatusFailed(status, message string, data []*response.TransferResponseYearStatusFailed) *model.APIResponseTransferYearStatusFailed {
	return &model.APIResponseTransferYearStatusFailed{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesYearStatusFailed(data),
	}
}

func (t *transferResponseMapper) mapResponseTransfer(transfer *response.TransferResponse) *model.TransferResponse {
	return &model.TransferResponse{
		ID:             int32(transfer.ID),
		TransferNo:     transfer.TransferNo,
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: int32(transfer.TransferAmount),
		TransferTime:   transfer.TransferTime,
		CreatedAt:      transfer.CreatedAt,
		UpdatedAt:      transfer.UpdatedAt,
	}
}

func (t *transferResponseMapper) mapResponsesTransfer(transfers []*response.TransferResponse) []*model.TransferResponse {
	var responses []*model.TransferResponse

	for _, transfer := range transfers {
		responses = append(responses, t.mapResponseTransfer(transfer))
	}

	return responses
}

func (t *transferResponseMapper) mapResponseTransferDeleteAt(transfer *response.TransferResponseDeleteAt) *model.TransferResponseDeleteAt {
	return &model.TransferResponseDeleteAt{
		ID:             int32(transfer.ID),
		TransferNo:     transfer.TransferNo,
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: int32(transfer.TransferAmount),
		TransferTime:   transfer.TransferTime,
		CreatedAt:      transfer.CreatedAt,
		UpdatedAt:      transfer.UpdatedAt,
		DeletedAt:      transfer.DeletedAt,
	}
}

func (t *transferResponseMapper) mapResponsesTransferDeleteAt(transfers []*response.TransferResponseDeleteAt) []*model.TransferResponseDeleteAt {
	var responses []*model.TransferResponseDeleteAt

	for _, transfer := range transfers {
		responses = append(responses, t.mapResponseTransferDeleteAt(transfer))
	}

	return responses
}

func (t *transferResponseMapper) mapResponseMonthStatusSuccess(data *response.TransferResponseMonthStatusSuccess) *model.TransferMonthStatusSuccessResponse {
	return &model.TransferMonthStatusSuccessResponse{
		Year:         data.Year,
		Month:        data.Month,
		TotalSuccess: int32(data.TotalSuccess),
		TotalAmount:  int32(data.TotalAmount),
	}
}

func (t *transferResponseMapper) mapResponsesMonthStatusSuccess(transfers []*response.TransferResponseMonthStatusSuccess) []*model.TransferMonthStatusSuccessResponse {
	var responses []*model.TransferMonthStatusSuccessResponse

	for _, transfer := range transfers {
		responses = append(responses, t.mapResponseMonthStatusSuccess(transfer))
	}

	return responses
}

func (t *transferResponseMapper) mapResponseYearStatusSuccess(data *response.TransferResponseYearStatusSuccess) *model.TransferYearStatusSuccessResponse {
	return &model.TransferYearStatusSuccessResponse{
		Year:         data.Year,
		TotalSuccess: int32(data.TotalSuccess),
		TotalAmount:  int32(data.TotalAmount),
	}
}

func (t *transferResponseMapper) mapResponsesYearStatusSuccess(transfers []*response.TransferResponseYearStatusSuccess) []*model.TransferYearStatusSuccessResponse {
	var responses []*model.TransferYearStatusSuccessResponse

	for _, transfer := range transfers {
		responses = append(responses, t.mapResponseYearStatusSuccess(transfer))
	}

	return responses
}

func (t *transferResponseMapper) mapResponseMonthStatusFailed(data *response.TransferResponseMonthStatusFailed) *model.TransferMonthStatusFailedResponse {
	return &model.TransferMonthStatusFailedResponse{
		Year:        data.Year,
		Month:       data.Month,
		TotalFailed: int32(data.TotalFailed),
		TotalAmount: int32(data.TotalAmount),
	}
}

func (t *transferResponseMapper) mapResponsesMonthStatusFailed(transfers []*response.TransferResponseMonthStatusFailed) []*model.TransferMonthStatusFailedResponse {
	var responses []*model.TransferMonthStatusFailedResponse

	for _, transfer := range transfers {
		responses = append(responses, t.mapResponseMonthStatusFailed(transfer))
	}

	return responses
}

func (t *transferResponseMapper) mapResponseYearStatusFailed(data *response.TransferResponseYearStatusFailed) *model.TransferYearStatusFailedResponse {
	return &model.TransferYearStatusFailedResponse{
		Year:        data.Year,
		TotalFailed: int32(data.TotalFailed),
		TotalAmount: int32(data.TotalAmount),
	}
}

func (t *transferResponseMapper) mapResponsesYearStatusFailed(transfers []*response.TransferResponseYearStatusFailed) []*model.TransferYearStatusFailedResponse {
	var responses []*model.TransferYearStatusFailedResponse

	for _, transfer := range transfers {
		responses = append(responses, t.mapResponseYearStatusFailed(transfer))
	}

	return responses
}

func (t *transferResponseMapper) mapResponseTransferMonthAmount(transfer *response.TransferMonthAmountResponse) *model.TransferMonthAmountResponse {
	return &model.TransferMonthAmountResponse{
		Month:       transfer.Month,
		TotalAmount: int32(transfer.TotalAmount),
	}
}

func (t *transferResponseMapper) mapResponsesTransferMonthAmount(transfers []*response.TransferMonthAmountResponse) []*model.TransferMonthAmountResponse {
	var responses []*model.TransferMonthAmountResponse

	for _, transfer := range transfers {
		responses = append(responses, t.mapResponseTransferMonthAmount(transfer))
	}

	return responses
}

func (t *transferResponseMapper) mapResponseTransferYearAmount(transfer *response.TransferYearAmountResponse) *model.TransferYearAmountResponse {
	return &model.TransferYearAmountResponse{
		Year:        transfer.Year,
		TotalAmount: int32(transfer.TotalAmount),
	}
}

func (t *transferResponseMapper) mapResponsesTransferYearAmount(transfers []*response.TransferYearAmountResponse) []*model.TransferYearAmountResponse {
	var responses []*model.TransferYearAmountResponse

	for _, transfer := range transfers {
		responses = append(responses, t.mapResponseTransferYearAmount(transfer))
	}

	return responses
}
