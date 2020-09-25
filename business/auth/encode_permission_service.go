package auth

import (
	"context"
	"github.com/gingerxman/eel"
)

type EncodePermissionService struct {
	eel.ServiceBase
}

func NewEncodePermissionService(ctx context.Context) *EncodePermissionService {
	service := new(EncodePermissionService)
	service.Ctx = ctx
	return service
}

func (this *EncodePermissionService) Encode(permission *Permission) *RPermission {
	rPermission := &RPermission{
		Id: permission.Id,
		Code: permission.Code,
		Name: permission.Name,
	}
	
	return rPermission
}

func (this *EncodePermissionService) EncodeMany(permissions []*Permission) []*RPermission {
	rows := make([]*RPermission, 0)
	for _, permission := range permissions {
		rows = append(rows, this.Encode(permission))
	}
	
	return rows
}

func init() {
}
