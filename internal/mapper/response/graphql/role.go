package graphql

import (
	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/response"
	"github.com/MamangRust/paymentgatewaygraphql/internal/graph/model"
)

type roleResponseMapper struct {
}

func NewRoleResponseMapper() *roleResponseMapper {
	return &roleResponseMapper{}
}

func (s *roleResponseMapper) ToGraphqlResponseAll(status, message string) *model.APIResponseRoleAll {
	return &model.APIResponseRoleAll{
		Status:  status,
		Message: message,
	}
}

func (s *roleResponseMapper) ToGraphqlResponseDelete(status, message string) *model.APIResponseRoleDelete {
	return &model.APIResponseRoleDelete{
		Status:  status,
		Message: message,
	}
}

func (s *roleResponseMapper) ToGraphqlResponseRole(status, message string, role *response.RoleResponse) *model.APIResponseRole {
	return &model.APIResponseRole{
		Status:  status,
		Message: message,
		Data:    s.mapResponseRole(role),
	}
}

func (s *roleResponseMapper) ToGraphqlResponseRoleDeleteAt(status, message string, role *response.RoleResponseDeleteAt) *model.APIResponseRoleDeleteAt {
	return &model.APIResponseRoleDeleteAt{
		Status:  status,
		Message: message,
		Data:    s.mapResponseRoleDeleteAt(role),
	}
}

func (s *roleResponseMapper) ToGraphqlResponsesRole(status, message string, role []*response.RoleResponse) *model.APIResponsesRole {
	return &model.APIResponsesRole{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesRole(role),
	}
}

func (s *roleResponseMapper) ToGraphqlResponsePaginationRole(status, message string, role []*response.RoleResponse, pagination *response.PaginationMeta) *model.APIResponsePaginationRole {
	return &model.APIResponsePaginationRole{
		Status:     status,
		Message:    message,
		Data:       s.mapResponsesRole(role),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (s *roleResponseMapper) ToGraphqlResponsePaginationRoleDeleteAt(status, message string, role []*response.RoleResponseDeleteAt, pagination *response.PaginationMeta) *model.APIResponsePaginationRoleDeleteAt {
	return &model.APIResponsePaginationRoleDeleteAt{
		Status:     status,
		Message:    message,
		Data:       s.mapResponsesRoleDeleteAt(role),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (s *roleResponseMapper) mapResponseRole(role *response.RoleResponse) *model.RoleResponse {
	return &model.RoleResponse{
		ID:        int32(role.ID),
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

func (s *roleResponseMapper) mapResponsesRole(roles []*response.RoleResponse) []*model.RoleResponse {
	var responseRoles []*model.RoleResponse

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapResponseRole(role))
	}

	return responseRoles
}

func (s *roleResponseMapper) mapResponseRoleDeleteAt(role *response.RoleResponseDeleteAt) *model.RoleResponseDeleteAt {
	return &model.RoleResponseDeleteAt{
		ID:        int32(role.ID),
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
		DeletedAt: role.DeletedAt,
	}
}

func (s *roleResponseMapper) mapResponsesRoleDeleteAt(roles []*response.RoleResponseDeleteAt) []*model.RoleResponseDeleteAt {
	var responseRoles []*model.RoleResponseDeleteAt

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapResponseRoleDeleteAt(role))
	}

	return responseRoles
}
