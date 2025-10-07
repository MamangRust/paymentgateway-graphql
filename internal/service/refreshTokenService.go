package service

import (
	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/requests"
	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/response"
	responseservice "github.com/MamangRust/paymentgatewaygraphql/internal/mapper/response/service"
	"github.com/MamangRust/paymentgatewaygraphql/internal/repository"
	refreshtoken_errors "github.com/MamangRust/paymentgatewaygraphql/pkg/errors/refresh_token_errors"
	"github.com/MamangRust/paymentgatewaygraphql/pkg/logger"

	"go.uber.org/zap"
)

type refreshTokenService struct {
	refreshTokenRepository repository.RefreshTokenRepository
	logger                 logger.LoggerInterface
	mapping                responseservice.RefreshTokenResponseMapper
}

func NewRefreshTokenService(refreshTokenRepository repository.RefreshTokenRepository, logger logger.LoggerInterface, mapping responseservice.RefreshTokenResponseMapper) *refreshTokenService {
	return &refreshTokenService{
		refreshTokenRepository: refreshTokenRepository,
		logger:                 logger,
		mapping:                mapping,
	}
}

func (r *refreshTokenService) FindByToken(token string) (*response.RefreshTokenResponse, *response.ErrorResponse) {
	refreshToken, err := r.refreshTokenRepository.FindByToken(token)

	if err != nil {
		r.logger.Error("Failed to find refresh token", zap.Error(err))

		return nil, refreshtoken_errors.ErrFailedFindByToken
	}

	return r.mapping.ToRefreshTokenResponse(refreshToken), nil
}

func (r *refreshTokenService) FindByUserId(user_id int) (*response.RefreshTokenResponse, *response.ErrorResponse) {
	refreshToken, err := r.refreshTokenRepository.FindByUserId(user_id)

	if err != nil {
		r.logger.Error("Failed to find refresh token", zap.Error(err))
		return nil, refreshtoken_errors.ErrFailedFindByUserID
	}

	return r.mapping.ToRefreshTokenResponse(refreshToken), nil
}

func (r *refreshTokenService) UpdateRefreshToken(req *requests.UpdateRefreshToken) (*response.RefreshTokenResponse, *response.ErrorResponse) {
	refreshToken, err := r.refreshTokenRepository.UpdateRefreshToken(req)

	if err != nil {
		r.logger.Error("Failed to update refresh token", zap.Error(err))
		return nil, refreshtoken_errors.ErrFailedUpdateRefreshToken
	}

	return r.mapping.ToRefreshTokenResponse(refreshToken), nil
}

func (r *refreshTokenService) DeleteRefreshToken(token string) *response.ErrorResponse {
	err := r.refreshTokenRepository.DeleteRefreshToken(token)

	if err != nil {
		r.logger.Error("Failed to delete refresh token", zap.Error(err))
		return refreshtoken_errors.ErrFailedDeleteRefreshToken
	}

	return nil
}
