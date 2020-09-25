package corp

import (
	"context"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-account/business/auth"
)

type EncodeCorpUserService struct {
	eel.ServiceBase
}

func NewEncodeCorpUserService(ctx context.Context) *EncodeCorpUserService {
	service := new(EncodeCorpUserService)
	service.Ctx = ctx
	return service
}

func (this *EncodeCorpUserService) Encode(corpUser *CorpUser) *RCorpUser {
	rGroups := auth.NewEncodeGroupService(this.Ctx).EncodeMany(corpUser.Groups)
	
	rPermissions := make([]*auth.RPermission, 0)
	for _, rGroup := range rGroups {
		rPermissions = append(rPermissions, rGroup.Permissions...)
	}
	
	return &RCorpUser{
		Id: corpUser.Id,
		CorpId: corpUser.CorpId,
		CorpName: corpUser.Corp.Name,
		IsActive: corpUser.IsActive,
		IsManager: corpUser.IsManager,
		Name: corpUser.Username,
		RealName: corpUser.RealName,
		Groups: rGroups,
		Permissions: rPermissions,
	}
}

func (this *EncodeCorpUserService) EncodeMany(corpUsers []*CorpUser) []*RCorpUser {
	rows := make([]*RCorpUser, 0)
	for _, corpUser := range corpUsers {
		rows = append(rows, this.Encode(corpUser))
	}
	
	return rows
}

func init() {
}
