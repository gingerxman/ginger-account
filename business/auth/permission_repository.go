package auth

import (
	"context"
	
	"github.com/gingerxman/eel"
	m_auth "github.com/gingerxman/ginger-account/models/corp"
)

type PermissionRepository struct {
	eel.RepositoryBase
}

func NewPermissionRepository(ctx context.Context) *PermissionRepository {
	service := new(PermissionRepository)
	service.Ctx = ctx
	return service
}

func (this *PermissionRepository) GetPermissions(filters eel.Map) []*Permission {
	o := eel.GetOrmFromContext(this.Ctx)
	
	var models []*m_auth.Permission
	db := o.Model(&m_auth.Permission{})
	if len(filters) > 0 {
		db = db.Where(filters)
	}
	db = db.Find(&models)
	
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return make([]*Permission, 0)
	}
	
	instances := make([]*Permission, 0)
	for _, model := range models {
		instances = append(instances, NewPermissionFromModel(this.Ctx, model))
	}
	return instances
}

func (this *PermissionRepository) GetPagedPermissions(filters eel.Map, page *eel.PageInfo) ([]*Permission, eel.INextPageInfo) {
	o := eel.GetOrmFromContext(this.Ctx)
	
	var models []*m_auth.Permission
	db := o.Model(&m_auth.Permission{})
	paginateResult, db := eel.Paginate(db, page, &models)
	
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return nil, paginateResult
	}
	
	instances := make([]*Permission, 0)
	for _, model := range models {
		instances = append(instances, NewPermissionFromModel(this.Ctx, model))
	}
	return instances, paginateResult
}

func (this *PermissionRepository) GetPermissionsByCodes(codes []string) []*Permission {
	filters := eel.Map{
		"code__in": codes,
	}
	
	return this.GetPermissions(filters)
}

func (this *PermissionRepository) GetPermissionsByIds(permissionIds []int) []*Permission {
	filters := eel.Map{
		"id__in": permissionIds,
	}
	
	return this.GetPermissions(filters)
}

func init() {
}
