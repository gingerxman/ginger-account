package auth

import (
	"context"
	"github.com/gingerxman/eel"
)

type EncodeGroupService struct {
	eel.ServiceBase
}

func NewEncodeGroupService(ctx context.Context) *EncodeGroupService {
	service := new(EncodeGroupService)
	service.Ctx = ctx
	return service
}

func (this *EncodeGroupService) Encode(group *Group) *RGroup {
	rPermissions := make([]*RPermission, 0)
	for _, permission := range group.Permissions {
		rPermissions = append(rPermissions, &RPermission{
			Id: permission.Id,
			Code: permission.Code,
		})
	}
	
	rGroup := &RGroup{
		Id: group.Id,
		Code: group.Code,
		Name: group.Name,
		Permissions: rPermissions,
	}
	
	return rGroup
}

func (this *EncodeGroupService) EncodeMany(groups []*Group) []*RGroup {
	rows := make([]*RGroup, 0)
	for _, group := range groups {
		rows = append(rows, this.Encode(group))
	}
	
	return rows
}

func init() {
}
