package graphql

import (
	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/response"
	"github.com/MamangRust/paymentgatewaygraphql/internal/graph/model"
)

type AuthGraphqlMapper interface {
	ToGraphqlResponseLogin(status, message string, response *response.TokenResponse) *model.APIResponseLogin
	ToGraphqlResponseRegister(status, message string, response *response.UserResponse) *model.APIResponseRegister
	ToGraphqlResponseRefreshToken(status, message string, response *response.TokenResponse) *model.APIResponseRefreshToken
	ToGraphqlResponseGetMe(status, message string, response *response.UserResponse) *model.APIResponseGetMe
}

type UserGraphqlMapper interface {
	ToGraphqlResponseUser(status, message string, user *response.UserResponse) *model.APIResponseUserResponse
	ToGraphqlResponseUserDeleteAt(status, message string, user *response.UserResponseDeleteAt) *model.APIResponseUserResponseDeleteAt
	ToGraphqlResponseUsers(status, message string, user []*response.UserResponse) *model.APIResponsesUser
	ToGraphqlResponseUserDelete(status, message string) *model.APIResponseUserDelete
	ToGraphqlResponseUserAll(status, message string) *model.APIResponseUserAll
	ToGraphqlResponsePaginationUser(status, message string, user []*response.UserResponse, pagination *response.PaginationMeta) *model.APIResponsePaginationUser
	ToGraphqlResponsePaginationUserDeleteAt(status, message string, user []*response.UserResponseDeleteAt, pagination *response.PaginationMeta) *model.APIResponsePaginationUserDeleteAt
}

type RoleGraphqlMapper interface {
	ToGraphqlResponseRole(status, message string, role *response.RoleResponse) *model.APIResponseRole
	ToGraphqlResponseRoleDeleteAt(status, message string, role *response.RoleResponseDeleteAt) *model.APIResponseRoleDeleteAt
	ToGraphqlResponsesRole(status, message string, role []*response.RoleResponse) *model.APIResponsesRole
	ToGraphqlResponseDelete(status, message string) *model.APIResponseRoleDelete
	ToGraphqlResponseAll(status, message string) *model.APIResponseRoleAll
	ToGraphqlResponsePaginationRole(status, message string, role []*response.RoleResponse, pagination *response.PaginationMeta) *model.APIResponsePaginationRole
	ToGraphqlResponsePaginationRoleDeleteAt(status, message string, role []*response.RoleResponseDeleteAt, pagination *response.PaginationMeta) *model.APIResponsePaginationRoleDeleteAt
}

type CardGraphqlMapper interface {
	ToGraphqlResponsePaginationCard(status string, Message string, card []*response.CardResponse, pagination *response.PaginationMeta) *model.APIResponsePaginationCard
	ToGraphqlResponseCard(status string, Message string, card *response.CardResponse) *model.APIResponseCard
	ToGraphqlResponseAll(status, message string) *model.APIResponseCardAll
	ToGraphqlResponseDelete(status, message string) *model.APIResponseCardDelete
	ToGraphqlResponseCardDeleteAt(status string, Message string, card *response.CardResponseDeleteAt) *model.APIResponseCardDeleteAt
	ToGraphqlResponsePaginationCardDeleteAt(status, message string, card []*response.CardResponseDeleteAt, pagination *response.PaginationMeta) *model.APIResponsePaginationCardDeleteAt
	ToGraphqlDashboardCard(status, message string, dash *response.DashboardCard) *model.APIResponseDashboardCard
	ToGraphqlDashboardCardCardNumber(status, message string, dash *response.DashboardCardCardNumber) *model.APIResponseDashboardCardNumber
	ToGraphqlMonthlyBalances(status, message string, cards []*response.CardResponseMonthBalance) *model.APIResponseMonthlyBalance
	ToGraphqlYearlyBalances(status, message string, cards []*response.CardResponseYearlyBalance) *model.APIResponseYearlyBalance
	ToGraphqlMonthlyAmounts(status, message string, card []*response.CardResponseMonthAmount) *model.APIResponseMonthlyAmount
	ToGraphqlYearlyAmounts(status, message string, card []*response.CardResponseYearAmount) *model.APIResponseYearlyAmount
}

type MerchantGraphqlMapper interface {
	ToGraphqlResponseMerchant(status, message string, merchant *response.MerchantResponse) *model.APIResponseMerchant
	ToGraphqlResponsesMerchant(status, message string, merchant []*response.MerchantResponse) *model.APIResponsesMerchant
	ToGraphqlResponsePaginationMerchant(status, message string, merchant []*response.MerchantResponse, pagination *response.PaginationMeta) *model.APIResponseMerchantPagination
	ToGraphqlResponseMerchantDeleteAt(status, message string, merchant *response.MerchantResponseDeleteAt) *model.APIResponseMerchantDeleteAt
	ToGraphqlResponsePaginationMerchantDeleteAt(status, message string, merchant []*response.MerchantResponseDeleteAt, pagination *response.PaginationMeta) *model.APIResponseMerchantDeleteAtPagination
	ToGraphqlResponsePaginationTransaction(status, message string, merchant []*response.MerchantTransactionResponse, pagination *response.PaginationMeta) *model.APIResponseMerchantTransactionPagination
	ToGraphqlMonthlyPaymentMethods(status, message string, merchant []*response.MerchantResponseMonthlyPaymentMethod) *model.APIResponseMerchantMonthlyPaymentMethod
	ToGraphqlYearlyPaymentMethods(status, message string, merchant []*response.MerchantResponseYearlyPaymentMethod) *model.APIResponseMerchantYearlyPaymentMethod
	ToGraphqlMonthlyAmounts(status, message string, ms []*response.MerchantResponseMonthlyAmount) *model.APIResponseMerchantMonthlyAmount
	ToGraphqlYearlyAmounts(status, message string, ms []*response.MerchantResponseYearlyAmount) *model.APIResponseMerchantYearlyAmount
	ToGraphqlMonthlyTotalAmounts(status, message string, ms []*response.MerchantResponseMonthlyTotalAmount) *model.APIResponseMerchantMonthlyTotalAmount
	ToGraphqlYearlyTotalAmounts(status, message string, ms []*response.MerchantResponseYearlyTotalAmount) *model.APIResponseMerchantYearlyTotalAmount
	ToGraphqlMerchantDeleteAll(status, message string) *model.APIResponseMerchantDelete
	ToGraphqlMerchantAll(status, message string) *model.APIResponseMerchantAll
}

type SaldoGraphqMapper interface {
	ToGraphqlResponseSaldo(status, message string, saldo *response.SaldoResponse) *model.APIResponseSaldoResponse
	ToGraphqlResponsesSaldo(status, message string, saldo []*response.SaldoResponse) *model.APIResponsesSaldo
	ToGraphqlResponseSaldoDeleteAt(status, message string, saldo *response.SaldoResponseDeleteAt) *model.APIResponseSaldoResponseDeleteAt

	ToGraphqlResponsePaginationSaldo(status, message string, saldo []*response.SaldoResponse, pagination *response.PaginationMeta) *model.APIResponsePaginationSaldo
	ToGraphqlResponsePaginationSaldoDeleteAt(status, message string, saldo []*response.SaldoResponseDeleteAt, pagination *response.PaginationMeta) *model.APIResponsePaginationSaldoDeleteAt
	ToGraphqlResponseDelete(status, message string) *model.APIResponseSaldoDelete
	ToGraphqlResponseAll(status, message string) *model.APIResponseSaldoAll
	ToGraphqlResponseMonthTotalSaldo(status, message string, response []*response.SaldoMonthTotalBalanceResponse) *model.APIResponseMonthTotalSaldo
	ToGraphqlResponseYearTotalSaldo(status, message string, response []*response.SaldoYearTotalBalanceResponse) *model.APIResponseYearTotalSaldo
	ToGraphqlResponseMonthSaldoBalances(status, message string, response []*response.SaldoMonthBalanceResponse) *model.APIResponseMonthSaldoBalances
	ToGraphqlResponseYearBalance(status, message string, response []*response.SaldoYearBalanceResponse) *model.APIResponseYearSaldoBalances
}

type TopupGraphqlMapper interface {
	ToGraphqlResponseTopup(status, message string, data *response.TopupResponse) *model.APIResponseTopup
	ToGraphqlResponseTopupDeleteAt(status, message string, data *response.TopupResponseDeleteAt) *model.APIResponseTopupDeleteAt
	ToGraphqlTopupAll(status, message string) *model.APIResponseTopupAll
	ToGraphqlTopupDelete(status, message string) *model.APIResponseTopupDelete
	ToGraphqlResponsePaginationTopup(status, message string, data []*response.TopupResponse, pagination *response.PaginationMeta) *model.APIResponsePaginationTopup
	ToGraphqlResponsePaginationTopupDeleteAt(status, message string, data []*response.TopupResponseDeleteAt, pagination *response.PaginationMeta) *model.APIResponsePaginationTopupDeleteAt
	ToGraphqlResponseTopupMonthStatusSuccess(status, message string, data []*response.TopupResponseMonthStatusSuccess) *model.APIResponseTopupMonthStatusSuccess
	ToGraphqlResponseTopupYearStatusSuccess(status, message string, data []*response.TopupResponseYearStatusSuccess) *model.APIResponseTopupYearStatusSuccess
	ToGraphqlResponseTopupMonthStatusFailed(status, message string, data []*response.TopupResponseMonthStatusFailed) *model.APIResponseTopupMonthStatusFailed
	ToGraphqlResponseTopupYearStatusFailed(status, message string, data []*response.TopupResponseYearStatusFailed) *model.APIResponseTopupYearStatusFailed
	ToGraphqlResponseTopupMonthMethod(status, message string, data []*response.TopupMonthMethodResponse) *model.APIResponseTopupMonthMethod
	ToGraphqlResponseTopupYearMethod(status, message string, data []*response.TopupYearlyMethodResponse) *model.APIResponseTopupYearMethod
	ToGraphqlResponseTopupMonthAmount(status, message string, data []*response.TopupMonthAmountResponse) *model.APIResponseTopupMonthAmount
	ToGraphqlResponseTopupYearAmount(status, message string, data []*response.TopupYearlyAmountResponse) *model.APIResponseTopupYearAmount
}

type TransactionGraphqlMapper interface {
	ToGraphqlResponseTransactionMonthStatusSuccess(status, message string, data []*response.TransactionResponseMonthStatusSuccess) *model.APIResponseTransactionMonthStatusSuccess
	ToGraphqlResponseTransactionYearStatusSuccess(status, message string, data []*response.TransactionResponseYearStatusSuccess) *model.APIResponseTransactionYearStatusSuccess
	ToGraphqlResponseTransactionMonthStatusFailed(status, message string, data []*response.TransactionResponseMonthStatusFailed) *model.APIResponseTransactionMonthStatusFailed
	ToGraphqlResponseTransactionYearStatusFailed(status, message string, data []*response.TransactionResponseYearStatusFailed) *model.APIResponseTransactionYearStatusFailed
	ToGraphqlResponseTransactionMonthMethod(status, message string, data []*response.TransactionMonthMethodResponse) *model.APIResponseTransactionMonthMethod
	ToGraphqlResponseTransactionYearMethod(status, message string, data []*response.TransactionYearMethodResponse) *model.APIResponseTransactionYearMethod
	ToGraphqlResponseTransactionMonthAmount(status, message string, data []*response.TransactionMonthAmountResponse) *model.APIResponseTransactionMonthAmount
	ToGraphqlResponseTransactionYearAmount(status, message string, data []*response.TransactionYearlyAmountResponse) *model.APIResponseTransactionYearAmount
	ToGraphqlTransactionAll(status, message string) *model.APIResponseTransactionAll
	ToGraphqlTransactionDelete(status, message string) *model.APIResponseTransactionDelete
	ToGraphqlPaginationTransaction(status, message string, data []*response.TransactionResponse, pagination *response.PaginationMeta) *model.APIResponsePaginationTransaction
	ToGraphqlPaginationTransactionDeleteAt(status, message string, data []*response.TransactionResponseDeleteAt, pagination *response.PaginationMeta) *model.APIResponsePaginationTransactionDeleteAt
	ToGraphqlResponseTransaction(status, message string, data *response.TransactionResponse) *model.APIResponseTransaction
	ToGraphqlResponseTransactions(status, message string, data []*response.TransactionResponse) *model.APIResponseTransactions
	ToGraphqlResponseTransactionDeleteAt(status, message string, data *response.TransactionResponseDeleteAt) *model.APIResponseTransactionDeleteAt
}

type TransferGraphqlMapper interface {
	ToGraphqlTransferAll(status, message string) *model.APIResponseTransferAll
	ToGraphqlTransferDelete(status, message string) *model.APIResponseTransferDelete
	ToGraphqlResponseTransfer(status, message string, data *response.TransferResponse) *model.APIResponseTransfer
	ToGraphqlResponseTransfers(status, message string, data []*response.TransferResponse) *model.APIResponseTransfers
	ToGraphqlResponseTransferDeleteAt(status, message string, data *response.TransferResponseDeleteAt) *model.APIResponseTransferDeleteAt
	ToGraphqlResponsePaginationTransfer(status, message string, data []*response.TransferResponse, pagination *response.PaginationMeta) *model.APIResponsePaginationTransfer
	ToGraphqlResponsePaginationTransferDeleteAt(status, message string, data []*response.TransferResponseDeleteAt, pagination *response.PaginationMeta) *model.APIResponsePaginationTransferDeleteAt
	ToGraphqlResponseTransferMonthAmount(status, message string, data []*response.TransferMonthAmountResponse) *model.APIResponseTransferMonthAmount
	ToGraphqlResponseTransferYearAmount(status, message string, data []*response.TransferYearAmountResponse) *model.APIResponseTransferYearAmount
	ToGraphqlResponseTransferMonthStatusSuccess(status, message string, data []*response.TransferResponseMonthStatusSuccess) *model.APIResponseTransferMonthStatusSuccess
	ToGraphqlResponseTransferYearStatusSuccess(status, message string, data []*response.TransferResponseYearStatusSuccess) *model.APIResponseTransferYearStatusSuccess
	ToGraphqlResponseTransferMonthStatusFailed(status, message string, data []*response.TransferResponseMonthStatusFailed) *model.APIResponseTransferMonthStatusFailed
	ToGraphqlResponseTransferYearStatusFailed(status, message string, data []*response.TransferResponseYearStatusFailed) *model.APIResponseTransferYearStatusFailed
}

type WithdrawGraphqlMapper interface {
	ToGraphqlWithdrawAll(status, message string) *model.APIResponseWithdrawAll
	ToGraphqlWithdrawDelete(status, message string) *model.APIResponseWithdrawDelete
	ToGraphqlResponseWithdraw(status, message string, data *response.WithdrawResponse) *model.APIResponseWithdraw
	ToGraphqlResponseWithdraws(status, message string, data []*response.WithdrawResponse) *model.APIResponsesWithdraw
	ToGraphqlResponseWithdrawDeleteAt(status, message string, data *response.WithdrawResponseDeleteAt) *model.APIResponseWithdrawDeleteAt
	ToGraphqlResponsePaginationWithdraw(status, message string, data []*response.WithdrawResponse, pagination *response.PaginationMeta) *model.APIResponsePaginationWithdraw
	ToGraphqlResponsePaginationWithdrawDeleteAt(status, message string, data []*response.WithdrawResponseDeleteAt, pagination *response.PaginationMeta) *model.APIResponsePaginationWithdrawDeleteAt
	ToGraphqlResponseWithdrawMonthAmount(status, message string, data []*response.WithdrawMonthlyAmountResponse) *model.APIResponseWithdrawMonthAmount
	ToGraphqlResponseWithdrawYearAmount(status, message string, data []*response.WithdrawYearlyAmountResponse) *model.APIResponseWithdrawYearAmount
	ToGraphqlResponseWithdrawMonthStatusSuccess(status, message string, data []*response.WithdrawResponseMonthStatusSuccess) *model.APIResponseWithdrawMonthStatusSuccess
	ToGraphqlResponseWithdrawYearStatusSuccess(status, message string, data []*response.WithdrawResponseYearStatusSuccess) *model.APIResponseWithdrawYearStatusSuccess
	ToGraphqlResponseWithdrawMonthStatusFailed(status, message string, data []*response.WithdrawResponseMonthStatusFailed) *model.APIResponseWithdrawMonthStatusFailed
	ToGraphqlResponseWithdrawYearStatusFailed(status, message string, data []*response.WithdrawResponseYearStatusFailed) *model.APIResponseWithdrawYearStatusFailed
}
