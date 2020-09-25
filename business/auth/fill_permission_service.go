package auth

import (
	"context"
	"github.com/gingerxman/eel"
)

type FillPermissionService struct {
	eel.ServiceBase
}

func NewFillPermissionService(ctx context.Context) *FillPermissionService {
	service := new(FillPermissionService)
	service.Ctx = ctx
	return service
}

func (this *FillPermissionService) FillOne(permission *Permission, option eel.FillOption) {
	this.Fill([]*Permission{permission}, option)
}

func (this *FillPermissionService) Fill(permissions []*Permission, option eel.FillOption)  {
	if len(permissions) == 0 {
		return
	}
	
}

func init() {
}
