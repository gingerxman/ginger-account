package corp

import (
	"context"
	
	"github.com/gingerxman/eel"
	m_corp "github.com/gingerxman/ginger-account/models/corp"
)

type CorpUserRepository struct {
	eel.RepositoryBase
}

func NewCorpUserRepository(ctx context.Context) *CorpUserRepository {
	service := new(CorpUserRepository)
	service.Ctx = ctx
	return service
}

func (this *CorpUserRepository) GetCorpUsers(filters eel.Map) []*CorpUser {
	o := eel.GetOrmFromContext(this.Ctx)
	
	var models []*m_corp.CorpUser
	db := o.Model(&m_corp.CorpUser{})
	if len(filters) > 0 {
		db = db.Where(filters)
	}
	db = db.Find(&models)
	
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return make([]*CorpUser, 0)
	}
	
	instances := make([]*CorpUser, 0)
	for _, model := range models {
		instances = append(instances, NewCorpUserFromModel(this.Ctx, model))
	}
	return instances
}

func (this *CorpUserRepository) GetPagedCorpUsers(filters eel.Map, page *eel.PageInfo) ([]*CorpUser, eel.INextPageInfo) {
	o := eel.GetOrmFromContext(this.Ctx)
	
	var models []*m_corp.CorpUser
	db := o.Model(&m_corp.CorpUser{})
	db = db.Where(filters).Order("id desc")
	paginateResult, db := eel.Paginate(db, page, &models)
	
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return nil, paginateResult
	}
	
	instances := make([]*CorpUser, 0)
	for _, model := range models {
		instances = append(instances, NewCorpUserFromModel(this.Ctx, model))
	}
	return instances, paginateResult
}

func (this *CorpUserRepository) GetCorpUsersByIds(ids []int) []*CorpUser {
	filters := eel.Map{
		"id__in": ids,
	}
	
	return this.GetCorpUsers(filters)
}

func init() {
}
