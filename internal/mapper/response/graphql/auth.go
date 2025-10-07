package graphql

import (
	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/response"
	"github.com/MamangRust/paymentgatewaygraphql/internal/graph/model"
)

type authResponseMapper struct {
}

func NewAuthResponseMapper() *authResponseMapper {
	return &authResponseMapper{}
}

func (s *authResponseMapper) ToGraphqlResponseLogin(status, message string, response *response.TokenResponse) *model.APIResponseLogin {
	return &model.APIResponseLogin{
		Status:  status,
		Message: message,
		Data: &model.TokenResponse{
			AccessToken:  response.AccessToken,
			RefreshToken: response.RefreshToken,
		},
	}
}

func (s *authResponseMapper) ToGraphqlResponseRegister(status, message string, response *response.UserResponse) *model.APIResponseRegister {
	return &model.APIResponseRegister{
		Status:  status,
		Message: message,
		Data: &model.UserResponse{
			ID:        int32(response.ID),
			Lastname:  response.LastName,
			Email:     response.Email,
			CreatedAt: response.CreatedAt,
			UpdatedAt: response.UpdatedAt,
		},
	}
}

func (s *authResponseMapper) ToGraphqlResponseRefreshToken(status, message string, response *response.TokenResponse) *model.APIResponseRefreshToken {
	return &model.APIResponseRefreshToken{
		Status:  status,
		Message: message,
		Data: &model.TokenResponse{
			AccessToken:  response.AccessToken,
			RefreshToken: response.RefreshToken,
		},
	}
}

func (s *authResponseMapper) ToGraphqlResponseGetMe(status, message string, response *response.UserResponse) *model.APIResponseGetMe {
	return &model.APIResponseGetMe{
		Status:  status,
		Message: message,
		Data: &model.UserResponse{
			ID:        int32(response.ID),
			Firstname: response.FirstName,
			Lastname:  response.LastName,
			Email:     response.Email,
			CreatedAt: response.CreatedAt,
			UpdatedAt: response.UpdatedAt,
		},
	}
}
