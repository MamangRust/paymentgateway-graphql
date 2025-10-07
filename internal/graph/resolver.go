package graph

import (
	"github.com/MamangRust/paymentgatewaygraphql/internal/mapper/response/graphql"
	"github.com/MamangRust/paymentgatewaygraphql/internal/permission"
	"github.com/MamangRust/paymentgatewaygraphql/internal/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AuthGraphql        AuthHandleGraphql
	RoleGraphql        RoleHandleGraphql
	UserGraphql        UserHandleGraphql
	CardGraphql        CardHandleGraphql
	MerchantGraphql    MerchantHandleGraphql
	SaldoGraphql       SaldoHandleGraphql
	TopupGraphql       TopupHandleGraphql
	TransactionGraphql TransactionHandleGraphql
	TransferGraphql    TransferHandleGraphql
	WithdrawGraphql    WithdrawHandleGraphql
}

type AuthHandleGraphql struct {
	AuthService service.AuthService
	Mapping     graphql.AuthGraphqlMapper
}

type RoleHandleGraphql struct {
	RoleService service.RoleService
	Mapping     graphql.RoleGraphqlMapper
}

type UserHandleGraphql struct {
	UserService service.UserService
	Mapping     graphql.UserGraphqlMapper
}

type CardHandleGraphql struct {
	CardService service.CardService
	Mapping     graphql.CardGraphqlMapper
}

type MerchantHandleGraphql struct {
	MerchantService service.MerchantService
	Mapping         graphql.MerchantGraphqlMapper
}

type SaldoHandleGraphql struct {
	SaldoService service.SaldoService
	Mapping      graphql.SaldoGraphqMapper
}

type TopupHandleGraphql struct {
	TopupService service.TopupService
	Mapping      graphql.TopupGraphqlMapper
}

type TransactionHandleGraphql struct {
	TransactionService service.TransactionService
	Mapping            graphql.TransactionGraphqlMapper
	Permission         permission.Permission
}

type TransferHandleGraphql struct {
	TransferService service.TransferService
	Mapping         graphql.TransferGraphqlMapper
}

type WithdrawHandleGraphql struct {
	WithdrawService service.WithdrawService
	Mapping         graphql.WithdrawGraphqlMapper
}

func NewResolver(
	authService service.AuthService,
	roleService service.RoleService,
	userService service.UserService,
	cardService service.CardService,
	merchantService service.MerchantService,
	saldoService service.SaldoService,
	topupService service.TopupService,
	transactionService service.TransactionService,
	transferService service.TransferService,
	withdrawService service.WithdrawService,
	mapper *graphql.GraphqlMapper,
	permission permission.Permission,
) *Resolver {
	return &Resolver{
		AuthGraphql: AuthHandleGraphql{
			AuthService: authService,
			Mapping:     mapper.AuthGraphqlMapper,
		},
		RoleGraphql: RoleHandleGraphql{
			RoleService: roleService,
			Mapping:     mapper.RoleGraphqlMapper,
		},
		UserGraphql: UserHandleGraphql{
			UserService: userService,
			Mapping:     mapper.UserGraphqlMapper,
		},
		CardGraphql: CardHandleGraphql{
			CardService: cardService,
			Mapping:     mapper.CardGraphqlMapper,
		},
		MerchantGraphql: MerchantHandleGraphql{
			MerchantService: merchantService,
			Mapping:         mapper.MerchantGraphqlMapper,
		},
		SaldoGraphql: SaldoHandleGraphql{
			SaldoService: saldoService,
			Mapping:      mapper.SaldoGraphqMapper,
		},
		TopupGraphql: TopupHandleGraphql{
			TopupService: topupService,
			Mapping:      mapper.TopupGraphqlMapper,
		},
		TransactionGraphql: TransactionHandleGraphql{
			TransactionService: transactionService,
			Mapping:            mapper.TransactionGraphqlMapper,
			Permission:         permission,
		},
		TransferGraphql: TransferHandleGraphql{
			TransferService: transferService,
			Mapping:         mapper.TransferGraphqlMapper,
		},
		WithdrawGraphql: WithdrawHandleGraphql{
			WithdrawService: withdrawService,
			Mapping:         mapper.WithdrawGraphqlMapper,
		},
	}
}
