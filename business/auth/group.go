package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/gingerxman/eel"
	m_auth "github.com/gingerxman/ginger-account/models/corp"
)


type Group struct {
	eel.EntityBase
	Id int
	Code string
	Name string
	
	Permissions []*Permission
}

func (this *Group) UpdatePermissions(codes []string) {
	if len(codes) == 0 {
		return
	}
	
	permissions := NewPermissionRepository(this.Ctx).GetPermissionsByCodes(codes)
	if len(permissions) == 0 {
		return
	}
	
	o := eel.GetOrmFromContext(this.Ctx)
	db := o.Model(&m_auth.GroupHasPermission{}).Where("group_id", this.Id).Delete(&m_auth.GroupHasPermission{})
	err := db.Error
	if err != nil {
		eel.Logger.Error(err)
	}
	
	for _, permission := range permissions {
		model := &m_auth.GroupHasPermission{
			GroupId: this.Id,
			PermissionId: permission.Id,
		}
		db = o.Create(model)
		err := db.Error
		if err != nil {
			eel.Logger.Error(err)
		}
	}
}

func NewGroup(ctx context.Context, code string, name string, permissionCodes []string) (*Group, error) {
	o := eel.GetOrmFromContext(ctx)
	if o.Model(&m_auth.Group{}).Where("code", code).Exist() {
		return nil, errors.New(fmt.Sprintf("group(%s) already exists", code))
	}
	
	model := &m_auth.Group{
		Code: code,
		Name: name,
	}
	db := o.Create(model)
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		panic(eel.NewBusinessError("corp_factory:create_corp_fail", "创建Corp失败"))
	}
	
	group := NewGroupFromModel(ctx, model)
	group.UpdatePermissions(permissionCodes)
	
	return group, nil
}

func NewGroupFromModel(ctx context.Context, model *m_auth.Group) *Group {
	instance := new(Group)
	instance.Ctx = ctx
	instance.Id = model.Id
	instance.Code = model.Code
	instance.Name = model.Name
	instance.Permissions = make([]*Permission, 0)
	
	return instance
}

func init() {
}
