package repository

import (
	"context"

	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/record"
	"github.com/MamangRust/paymentgatewaygraphql/internal/domain/requests"
	recordmapper "github.com/MamangRust/paymentgatewaygraphql/internal/mapper/record"
	db "github.com/MamangRust/paymentgatewaygraphql/pkg/database/schema"
	userrole_errors "github.com/MamangRust/paymentgatewaygraphql/pkg/errors/user_role_errors"
)

type userRoleRepository struct {
	db      *db.Queries
	ctx     context.Context
	mapping recordmapper.UserRoleRecordMapping
}

func NewUserRoleRepository(db *db.Queries, ctx context.Context, mapping recordmapper.UserRoleRecordMapping) *userRoleRepository {
	return &userRoleRepository{
		db:      db,
		ctx:     ctx,
		mapping: mapping,
	}
}

func (r *userRoleRepository) AssignRoleToUser(req *requests.CreateUserRoleRequest) (*record.UserRoleRecord, error) {
	res, err := r.db.AssignRoleToUser(r.ctx, db.AssignRoleToUserParams{
		UserID: int32(req.UserId),
		RoleID: int32(req.RoleId),
	})

	if err != nil {
		return nil, userrole_errors.ErrAssignRoleToUser
	}

	return r.mapping.ToUserRoleRecord(res), nil
}

func (r *userRoleRepository) RemoveRoleFromUser(req *requests.RemoveUserRoleRequest) error {
	err := r.db.RemoveRoleFromUser(r.ctx, db.RemoveRoleFromUserParams{
		UserID: int32(req.UserId),
		RoleID: int32(req.RoleId),
	})

	if err != nil {
		return userrole_errors.ErrRemoveRole
	}

	return nil
}
