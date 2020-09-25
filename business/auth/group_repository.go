package auth

import (
	"context"
	
	m_auth "github.com/gingerxman/ginger-account/models/corp"
	"github.com/gingerxman/eel"
)

type GroupRepository struct {
	eel.RepositoryBase
}

func NewGroupRepository(ctx context.Context) *GroupRepository {
	service := new(GroupRepository)
	service.Ctx = ctx
	return service
}

func (this *GroupRepository) GetGroups(filters eel.Map) []*Group {
	o := eel.GetOrmFromContext(this.Ctx)
	
	var models []*m_auth.Group
	db := o.Model(&m_auth.Group{})
	if len(filters) > 0 {
		db = db.Where(filters)
	}
	db = db.Find(&models)
	
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return make([]*Group, 0)
	}
	
	instances := make([]*Group, 0)
	for _, model := range models {
		instances = append(instances, NewGroupFromModel(this.Ctx, model))
	}
	return instances
}

func (this *GroupRepository) GetPagedGroups(filters eel.Map, page *eel.PageInfo) ([]*Group, eel.INextPageInfo) {
	o := eel.GetOrmFromContext(this.Ctx)
	
	var models []*m_auth.Group
	db := o.Model(&m_auth.Group{})
	paginateResult, db := eel.Paginate(db, page, &models)
	
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return nil, paginateResult
	}
	
	instances := make([]*Group, 0)
	for _, model := range models {
		instances = append(instances, NewGroupFromModel(this.Ctx, model))
	}
	return instances, paginateResult
}

func (this *GroupRepository) GetGroupByName(name string) *Group {
	filters := eel.Map{
		"name": name,
	}
	
	groups := this.GetGroups(filters)
	if len(groups) == 0 {
		return nil
	} else {
		return groups[0]
	}
}

func (this *GroupRepository) GetGroupsByNames(names []string) []*Group {
	filters := eel.Map{
		"name__in": names,
	}
	
	return this.GetGroups(filters)
}

func (this *GroupRepository) GetGroupsByIds(groupIds []int) []*Group {
	filters := eel.Map{
		"id__in": groupIds,
	}
	
	return this.GetGroups(filters)
}

func (this *GroupRepository) GetGroupsForCorpUser(corpUserId int) []*Group {
	var models []*m_auth.UserHasGroup
	eel.GetOrmFromContext(this.Ctx).Model(&m_auth.UserHasGroup{}).Where("corp_user_id", corpUserId).Find(&models)
	
	groupIds := make([]int, 0)
	for _, model := range models {
		groupIds = append(groupIds, model.GroupId)
	}
	
	return this.GetGroupsByIds(groupIds)
}

func init() {
}
