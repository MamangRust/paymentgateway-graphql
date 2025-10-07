package graphql

type GraphqlMapper struct {
	AuthGraphqlMapper
	RoleGraphqlMapper
	UserGraphqlMapper
	CardGraphqlMapper
	MerchantGraphqlMapper
	SaldoGraphqMapper
	TopupGraphqlMapper
	TransactionGraphqlMapper
	TransferGraphqlMapper
	WithdrawGraphqlMapper
}

func NewGraphqlMapper() *GraphqlMapper {
	return &GraphqlMapper{
		AuthGraphqlMapper:        NewAuthResponseMapper(),
		UserGraphqlMapper:        NewUserResponseMapper(),
		RoleGraphqlMapper:        NewRoleResponseMapper(),
		MerchantGraphqlMapper:    NewMerchantResponseMapper(),
		CardGraphqlMapper:        NewCardResponseMapper(),
		SaldoGraphqMapper:        NewSaldoResponseMapper(),
		TopupGraphqlMapper:       NewTopupResponseMapper(),
		TransactionGraphqlMapper: NewTransactionResponseMapper(),
		TransferGraphqlMapper:    NewTransferResponseMapper(),
		WithdrawGraphqlMapper:    NewWithdrawResponseMapper(),
	}
}
