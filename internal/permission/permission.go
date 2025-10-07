package permission

import (
	"errors"

	"github.com/MamangRust/paymentgatewaygraphql/internal/service"
)

type Permission interface {
	HasRole(userID int, allowedRoles ...string) (bool, error)
	ValidateApiKey(apiKey string) (bool, error)
}

type permission struct {
	roleService     service.RoleService
	merchantService service.MerchantService
}

func NewPermission(roleService service.RoleService, merchantService service.MerchantService) *permission {
	return &permission{
		roleService:     roleService,
		merchantService: merchantService,
	}
}

func (s *permission) HasRole(userID int, allowedRoles ...string) (bool, error) {
	roleResponse, errResp := s.roleService.FindByUserId(userID)
	if errResp != nil {
		return false, errors.New("failed to fetch user role")
	}

	for _, role := range roleResponse {
		for _, allowed := range allowedRoles {
			if role.Name == allowed {
				return true, nil
			}
		}
	}
	return false, nil
}

func (p *permission) ValidateApiKey(apiKey string) (bool, error) {
	if apiKey == "" {
		return false, errors.New("missing API key")
	}

	_, err := p.merchantService.FindByApiKey(apiKey)
	if err != nil {
		return false, errors.New("invalid API key")
	}

	return true, nil
}
