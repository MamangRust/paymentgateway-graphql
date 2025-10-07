package graphql

import (
	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/response"
	"github.com/MamangRust/paymentgatewaygraphql/internal/graph/model"
)

type userResponseMapper struct {
}

func NewUserResponseMapper() *userResponseMapper {
	return &userResponseMapper{}
}

func (u *userResponseMapper) ToGraphqlResponseUserDelete(status, message string) *model.APIResponseUserDelete {
	return &model.APIResponseUserDelete{
		Status:  status,
		Message: message,
	}
}

func (u *userResponseMapper) ToGraphqlResponseUserAll(status, message string) *model.APIResponseUserAll {
	return &model.APIResponseUserAll{
		Status:  status,
		Message: message,
	}
}

func (u *userResponseMapper) ToGraphqlResponseUser(status, message string, user *response.UserResponse) *model.APIResponseUserResponse {
	return &model.APIResponseUserResponse{
		Status:  status,
		Message: message,
		Data:    u.mapUserResponse(user),
	}
}

func (u *userResponseMapper) ToGraphqlResponseUserDeleteAt(status, message string, user *response.UserResponseDeleteAt) *model.APIResponseUserResponseDeleteAt {
	return &model.APIResponseUserResponseDeleteAt{
		Status:  status,
		Message: message,
		Data:    u.mapUserResponseDeleteAt(user),
	}
}

func (u *userResponseMapper) ToGraphqlResponseUsers(status, message string, user []*response.UserResponse) *model.APIResponsesUser {
	return &model.APIResponsesUser{
		Status:  status,
		Message: message,
		Data:    u.mapUserResponses(user),
	}
}

func (u *userResponseMapper) ToGraphqlResponsePaginationUser(status, message string, user []*response.UserResponse, pagination *response.PaginationMeta) *model.APIResponsePaginationUser {
	return &model.APIResponsePaginationUser{
		Status:     status,
		Message:    message,
		Data:       u.mapUserResponses(user),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (u *userResponseMapper) ToGraphqlResponsePaginationUserDeleteAt(status, message string, user []*response.UserResponseDeleteAt, pagination *response.PaginationMeta) *model.APIResponsePaginationUserDeleteAt {
	return &model.APIResponsePaginationUserDeleteAt{
		Status:     status,
		Message:    message,
		Data:       u.mapUserResponsesDeleteAt(user),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (u *userResponseMapper) mapUserResponse(user *response.UserResponse) *model.UserResponse {
	return &model.UserResponse{
		ID:        int32(user.ID),
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (u *userResponseMapper) mapUserResponses(users []*response.UserResponse) []*model.UserResponse {
	var responses []*model.UserResponse

	for _, user := range users {
		responses = append(responses, u.mapUserResponse(user))
	}

	return responses
}

func (u *userResponseMapper) mapUserResponseDeleteAt(user *response.UserResponseDeleteAt) *model.UserResponseDeleteAt {
	return &model.UserResponseDeleteAt{
		ID:        int32(user.ID),
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}

func (u *userResponseMapper) mapUserResponsesDeleteAt(users []*response.UserResponseDeleteAt) []*model.UserResponseDeleteAt {
	var responses []*model.UserResponseDeleteAt

	for _, user := range users {
		responses = append(responses, u.mapUserResponseDeleteAt(user))
	}

	return responses
}
