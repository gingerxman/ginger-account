package corp

import (
	"context"
	
	"github.com/gingerxman/eel"
	m_corp "github.com/gingerxman/ginger-account/models/corp"
)

type CorpRepository struct {
	eel.RepositoryBase
}

func NewCorpRepository(ctx context.Context) *CorpRepository {
	service := new(CorpRepository)
	service.Ctx = ctx
	return service
}

func (this *CorpRepository) GetCorps(filters eel.Map) []*Corp {
	o := eel.GetOrmFromContext(this.Ctx)
	
	var models []*m_corp.Corp
	db := o.Model(&m_corp.Corp{})
	if len(filters) > 0 {
		db = db.Where(filters)
	}
	db = db.Find(&models)
	
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return make([]*Corp, 0)
	}
	
	instances := make([]*Corp, 0)
	for _, model := range models {
		instances = append(instances, NewCorpFromModel(this.Ctx, model))
	}
	return instances
}

func (this *CorpRepository) GetPagedCorps(filters eel.Map, page *eel.PageInfo) ([]*Corp, eel.INextPageInfo) {
	o := eel.GetOrmFromContext(this.Ctx)
	
	var models []*m_corp.Corp
	db := o.Model(&m_corp.Corp{})
	filters["is_active"] = true
	db = db.Where(filters).Order("id desc")
	paginateResult, db := eel.Paginate(db, page, &models)
	
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return nil, paginateResult
	}
	
	instances := make([]*Corp, 0)
	for _, model := range models {
		instances = append(instances, NewCorpFromModel(this.Ctx, model))
	}
	return instances, paginateResult
}

func (this *CorpRepository) GetCorpsByIds(ids []int) []*Corp {
	o := eel.GetOrmFromContext(this.Ctx)
	
	var models []*m_corp.Corp
	db := o.Model(&m_corp.Corp{}).Where("id__in", ids).Find(&models)
	
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return make([]*Corp, 0)
	}
	
	instances := make([]*Corp, 0)
	for _, model := range models {
		instances = append(instances, NewCorpFromModel(this.Ctx, model))
	}
	return instances
}

func (this *CorpRepository) GetCorpById(id int) *Corp {
	users := this.GetCorpsByIds([]int{id})
	
	if len(users) == 0 {
		return nil
	} else {
		return users[0]
	}
}

func init() {
}
