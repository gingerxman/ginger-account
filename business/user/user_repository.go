package user

import (
	"context"
	
	"github.com/gingerxman/eel"
	m_user "github.com/gingerxman/ginger-account/models/user"
)

type UserRepository struct {
	eel.RepositoryBase
}

func NewUserRepository(ctx context.Context) *UserRepository {
	service := new(UserRepository)
	service.Ctx = ctx
	return service
}

func (this *UserRepository) GetUsers(filters eel.Map) []*User {
	o := eel.GetOrmFromContext(this.Ctx)
	
	var models []*m_user.User
	db := o.Model(&m_user.User{})
	if len(filters) > 0 {
		db = db.Where(filters)
	}
	db = db.Find(&models)
	
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return make([]*User, 0)
	}
	
	instances := make([]*User, 0)
	for _, model := range models {
		instances = append(instances, NewUserFromModel(this.Ctx, model))
	}
	return instances
}

func (this *UserRepository) GetPagedUsers(filters eel.Map, page *eel.PageInfo) ([]*User, eel.INextPageInfo) {
	o := eel.GetOrmFromContext(this.Ctx)
	
	var models []*m_user.User
	db := o.Model(&m_user.User{})
	filters["is_active"] = true
	db = db.Where(filters).Order("id desc")
	paginateResult, db := eel.Paginate(db, page, &models)
	
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return nil, paginateResult
	}
	
	instances := make([]*User, 0)
	for _, model := range models {
		instances = append(instances, NewUserFromModel(this.Ctx, model))
	}
	return instances, paginateResult
}

func (this *UserRepository) GetUsersByIds(ids []int) []*User {
	o := eel.GetOrmFromContext(this.Ctx)
	
	var models []*m_user.User
	db := o.Model(&m_user.User{}).Where("id__in", ids).Find(&models)
	
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return make([]*User, 0)
	}
	
	instances := make([]*User, 0)
	for _, model := range models {
		instances = append(instances, NewUserFromModel(this.Ctx, model))
	}
	return instances
}

func (this *UserRepository) GetUserById(id int) *User {
	users := this.GetUsersByIds([]int{id})
	
	if len(users) == 0 {
		return nil
	} else {
		return users[0]
	}
}

func (this *UserRepository) GetUserByUnionid(unionid string) *User {
	filters := eel.Map{
		"unionid": unionid,
	}
	users := this.GetUsers(filters)
	
	if len(users) == 0 {
		return nil
	} else {
		return users[0]
	}
}

func init() {
}
