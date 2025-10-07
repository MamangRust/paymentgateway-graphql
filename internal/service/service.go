package service

import (
	responseservice "github.com/MamangRust/paymentgatewaygraphql/internal/mapper/response/service"
	"github.com/MamangRust/paymentgatewaygraphql/internal/repository"
	"github.com/MamangRust/paymentgatewaygraphql/pkg/auth"
	"github.com/MamangRust/paymentgatewaygraphql/pkg/hash"
	"github.com/MamangRust/paymentgatewaygraphql/pkg/logger"
)

type Service struct {
	Auth        AuthService
	User        UserService
	Role        RoleService
	Saldo       SaldoService
	Topup       TopupService
	Transfer    TransferService
	Withdraw    WithdrawService
	Card        CardService
	Merchant    MerchantService
	Transaction TransactionService
}

type Deps struct {
	Repositories *repository.Repositories
	Token        auth.TokenManager
	Hash         hash.HashPassword
	Logger       logger.LoggerInterface
	Mapper       responseservice.ResponseServiceMapper
}

func NewService(deps Deps) *Service {
	return &Service{
		Auth:        NewAuthService(deps.Repositories.User, deps.Repositories.RefreshToken, deps.Repositories.Role, deps.Repositories.UserRole, deps.Hash, deps.Token, deps.Logger, deps.Mapper.UserResponseMapper),
		User:        NewUserService(deps.Repositories.User, deps.Logger, deps.Mapper.UserResponseMapper, deps.Hash),
		Role:        NewRoleService(deps.Repositories.Role, deps.Logger, deps.Mapper.RoleResponseMapper),
		Saldo:       NewSaldoService(deps.Repositories.Saldo, deps.Repositories.Card, deps.Logger, deps.Mapper.SaldoResponseMapper),
		Topup:       NewTopupService(deps.Repositories.Card, deps.Repositories.Topup, deps.Repositories.Saldo, deps.Logger, deps.Mapper.TopupResponseMapper),
		Transfer:    NewTransferService(deps.Repositories.User, deps.Repositories.Card, deps.Repositories.Transfer, deps.Repositories.Saldo, deps.Logger, deps.Mapper.TransferResponseMapper),
		Withdraw:    NewWithdrawService(deps.Repositories.User, deps.Repositories.Withdraw, deps.Repositories.Saldo, deps.Logger, deps.Mapper.WithdrawResponseMapper),
		Card:        NewCardService(deps.Repositories.Card, deps.Repositories.User, deps.Logger, deps.Mapper.CardResponseMapper),
		Merchant:    NewMerchantService(deps.Repositories.Merchant, deps.Logger, deps.Mapper.MerchantResponseMapper),
		Transaction: NewTransactionService(deps.Repositories.Merchant, deps.Repositories.Card, deps.Repositories.Saldo, deps.Repositories.Transaction, deps.Logger, deps.Mapper.TransactionResponseMapper),
	}
}
