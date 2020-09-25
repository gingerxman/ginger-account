package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/gingerxman/eel"
	m_auth "github.com/gingerxman/ginger-account/models/corp"
)


type Permission struct {
	eel.EntityBase
	Id int
	Code string
	Name string
}

func NewPermission(ctx context.Context, code string, name string) (*Permission, error) {
	o := eel.GetOrmFromContext(ctx)
	if o.Model(&m_auth.Permission{}).Where("code", code).Exist() {
		return nil, errors.New(fmt.Sprintf("permission(%s) already exists", code))
	}
	
	model := &m_auth.Permission{
		Code: code,
		Name: name,
	}
	db := o.Create(model)
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return nil, db.Error
	}
	
	permission := NewPermissionFromModel(ctx, model)
	return permission, nil
}

func NewPermissionFromModel(ctx context.Context, model *m_auth.Permission) *Permission {
	instance := new(Permission)
	instance.Ctx = ctx
	instance.Id = model.Id
	instance.Code = model.Code
	instance.Name = model.Name
	
	return instance
}

func init() {
}
