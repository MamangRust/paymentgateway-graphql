package graphql

import (
	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/response"
	"github.com/MamangRust/paymentgatewaygraphql/internal/graph/model"
)

type cardResponseMapper struct {
}

func NewCardResponseMapper() *cardResponseMapper {
	return &cardResponseMapper{}
}

func (s *cardResponseMapper) ToGraphqlResponseCard(status string, Message string, card *response.CardResponse) *model.APIResponseCard {
	return &model.APIResponseCard{
		Status:  status,
		Message: Message,
		Data:    s.mapCardResponse(card),
	}
}

func (s *cardResponseMapper) ToGraphqlResponsePaginationCard(status string, Message string, card []*response.CardResponse, pagination *response.PaginationMeta) *model.APIResponsePaginationCard {
	return &model.APIResponsePaginationCard{
		Status:     status,
		Message:    Message,
		Data:       s.mapCardResponses(card),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (s *cardResponseMapper) ToGraphqlResponseAll(status, message string) *model.APIResponseCardAll {
	return &model.APIResponseCardAll{
		Status:  status,
		Message: message,
	}
}

func (s *cardResponseMapper) ToGraphqlResponseDelete(status, message string) *model.APIResponseCardDelete {
	return &model.APIResponseCardDelete{
		Status:  status,
		Message: message,
	}
}

func (s *cardResponseMapper) ToGraphqlResponseCardDeleteAt(status string, Message string, card *response.CardResponseDeleteAt) *model.APIResponseCardDeleteAt {
	return &model.APIResponseCardDeleteAt{
		Status:  status,
		Message: Message,
		Data:    s.mapCardResponseDeleteAt(card),
	}
}

func (s *cardResponseMapper) ToGraphqlResponsePaginationCardDeleteAt(status, message string, card []*response.CardResponseDeleteAt, pagination *response.PaginationMeta) *model.APIResponsePaginationCardDeleteAt {
	return &model.APIResponsePaginationCardDeleteAt{
		Status:     status,
		Message:    message,
		Data:       s.mapCardDeleteAtResponses(card),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (s *cardResponseMapper) ToGraphqlDashboardCard(status, message string, dash *response.DashboardCard) *model.APIResponseDashboardCard {
	return &model.APIResponseDashboardCard{
		Status:  status,
		Message: message,
		Data:    s.mapDashboardCard(dash),
	}
}

func (s *cardResponseMapper) ToGraphqlDashboardCardCardNumber(status, message string, dash *response.DashboardCardCardNumber) *model.APIResponseDashboardCardNumber {
	return &model.APIResponseDashboardCardNumber{
		Status:  status,
		Message: message,
		Data:    s.mapDashboardCardCardNumber(dash),
	}
}

func (s *cardResponseMapper) ToGraphqlMonthlyBalances(status, message string, cards []*response.CardResponseMonthBalance) *model.APIResponseMonthlyBalance {
	return &model.APIResponseMonthlyBalance{
		Status:  status,
		Message: message,
		Data:    s.mapMonthlyBalances(cards),
	}
}

func (s *cardResponseMapper) ToGraphqlYearlyBalances(status, message string, cards []*response.CardResponseYearlyBalance) *model.APIResponseYearlyBalance {
	return &model.APIResponseYearlyBalance{
		Status:  status,
		Message: message,
		Data:    s.mapYearlyBalances(cards),
	}
}

func (s *cardResponseMapper) ToGraphqlMonthlyAmounts(status, message string, card []*response.CardResponseMonthAmount) *model.APIResponseMonthlyAmount {
	return &model.APIResponseMonthlyAmount{
		Status:  status,
		Message: message,
		Data:    s.mapMonthlyAmounts(card),
	}
}

func (s *cardResponseMapper) ToGraphqlYearlyAmounts(status, message string, card []*response.CardResponseYearAmount) *model.APIResponseYearlyAmount {
	return &model.APIResponseYearlyAmount{
		Status:  status,
		Message: message,
		Data:    s.mapYearlyAmounts(card),
	}
}

// map

func (s *cardResponseMapper) mapCardResponse(card *response.CardResponse) *model.CardResponse {
	return &model.CardResponse{
		ID:           int32(card.ID),
		UserID:       int32(card.UserID),
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate,
		Cvv:          card.CVV,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt,
		UpdatedAt:    card.UpdatedAt,
	}
}

func (s *cardResponseMapper) mapCardResponses(cards []*response.CardResponse) []*model.CardResponse {
	var responseCards []*model.CardResponse

	for _, card := range cards {
		responseCards = append(responseCards, s.mapCardResponse(card))
	}

	return responseCards
}

func (s *cardResponseMapper) mapCardResponseDeleteAt(card *response.CardResponseDeleteAt) *model.CardResponseDeleteAt {
	return &model.CardResponseDeleteAt{
		ID:           int32(card.ID),
		UserID:       int32(card.UserID),
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate,
		Cvv:          card.CVV,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt,
		UpdatedAt:    card.UpdatedAt,
		DeletedAt:    card.DeletedAt,
	}
}

func (s *cardResponseMapper) mapCardDeleteAtResponses(cards []*response.CardResponseDeleteAt) []*model.CardResponseDeleteAt {
	var responseCards []*model.CardResponseDeleteAt

	for _, card := range cards {
		responseCards = append(responseCards, s.mapCardResponseDeleteAt(card))
	}

	return responseCards
}

func (s *cardResponseMapper) mapDashboardCard(dash *response.DashboardCard) *model.CardDashboardResponse {
	return &model.CardDashboardResponse{
		TotalBalance:     int32(*dash.TotalBalance),
		TotalWithdraw:    int32(*dash.TotalWithdraw),
		TotalTopup:       int32(*dash.TotalTopup),
		TotalTransfer:    int32(*dash.TotalTransfer),
		TotalTransaction: int32(*dash.TotalTransaction),
	}
}

func (s *cardResponseMapper) mapDashboardCardCardNumber(dash *response.DashboardCardCardNumber) *model.CardDashboardByNumberResponse {
	return &model.CardDashboardByNumberResponse{
		TotalBalance:          int32(*dash.TotalBalance),
		TotalWithdraw:         int32(*dash.TotalWithdraw),
		TotalTopup:            int32(*dash.TotalTopup),
		TotalTransferSend:     int32(*dash.TotalTransferSend),
		TotalTransferReceiver: int32(*dash.TotalTransferReceiver),
		TotalTransaction:      int32(*dash.TotalTransaction),
	}
}

func (s *cardResponseMapper) mapMonthlyBalance(cards *response.CardResponseMonthBalance) *model.CardMonthlyBalanceResponse {
	return &model.CardMonthlyBalanceResponse{
		Month:        cards.Month,
		TotalBalance: int32(cards.TotalBalance),
	}
}

func (s *cardResponseMapper) mapMonthlyBalances(cards []*response.CardResponseMonthBalance) []*model.CardMonthlyBalanceResponse {
	var responseCards []*model.CardMonthlyBalanceResponse

	for _, role := range cards {
		responseCards = append(responseCards, s.mapMonthlyBalance(role))
	}

	return responseCards
}

func (s *cardResponseMapper) mapYearlyBalance(cards *response.CardResponseYearlyBalance) *model.CardYearlyBalanceResponse {
	return &model.CardYearlyBalanceResponse{
		Year:         cards.Year,
		TotalBalance: int32(cards.TotalBalance),
	}
}

func (s *cardResponseMapper) mapYearlyBalances(cards []*response.CardResponseYearlyBalance) []*model.CardYearlyBalanceResponse {
	var responseCards []*model.CardYearlyBalanceResponse

	for _, role := range cards {
		responseCards = append(responseCards, s.mapYearlyBalance(role))
	}

	return responseCards
}

func (s *cardResponseMapper) mapMonthlyAmount(cards *response.CardResponseMonthAmount) *model.CardMonthlyAmountResponse {
	return &model.CardMonthlyAmountResponse{
		Month:       cards.Month,
		TotalAmount: int32(cards.TotalAmount),
	}
}

func (s *cardResponseMapper) mapMonthlyAmounts(cards []*response.CardResponseMonthAmount) []*model.CardMonthlyAmountResponse {
	var responseCards []*model.CardMonthlyAmountResponse

	for _, role := range cards {
		responseCards = append(responseCards, s.mapMonthlyAmount(role))
	}

	return responseCards
}

func (s *cardResponseMapper) mapYearlyAmount(cards *response.CardResponseYearAmount) *model.CardYearlyAmountResponse {
	return &model.CardYearlyAmountResponse{
		Year:        cards.Year,
		TotalAmount: int32(cards.TotalAmount),
	}
}

func (s *cardResponseMapper) mapYearlyAmounts(cards []*response.CardResponseYearAmount) []*model.CardYearlyAmountResponse {
	var responseCards []*model.CardYearlyAmountResponse

	for _, role := range cards {
		responseCards = append(responseCards, s.mapYearlyAmount(role))
	}

	return responseCards
}

func mapPaginationMeta(s *response.PaginationMeta) *model.PaginationMeta {
	return &model.PaginationMeta{
		CurrentPage:  int32(s.CurrentPage),
		PageSize:     int32(s.PageSize),
		TotalRecords: int32(s.TotalRecords),
		TotalPages:   int32(s.TotalPages),
	}
}
