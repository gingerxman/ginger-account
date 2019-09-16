package corp

import (
	"context"
	"errors"
	"fmt"
	"github.com/gingerxman/eel"
	m_corp "github.com/gingerxman/ginger-account/models/corp"
)

type CorpFactory struct {
	eel.ServiceBase
}

func NewCorpFactory(ctx context.Context) *CorpFactory {
	service := new(CorpFactory)
	service.Ctx = ctx
	return service
}

func(this *CorpFactory) createCorp(name string, isPlatform bool) (*Corp, error) {
	o := eel.GetOrmFromContext(this.Ctx)
	if o.Model(&m_corp.Corp{}).Where("name", name).Exist() {
		return nil, errors.New(fmt.Sprintf("corp(%s) already exists", name))
	}
	
	model := &m_corp.Corp{
		Name: name,
		Code: name,
		IsPlatform: isPlatform,
	}
	db := o.Create(model)
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		panic(eel.NewBusinessError("corp_factory:create_corp_fail", "创建Corp失败"))
	}
	
	corp := NewCorpFromModel(this.Ctx, model)
	return corp, nil
}

func(this *CorpFactory) CreateCorp(name string) (*Corp, error) {
	return this.createCorp(name, false)
}

func(this *CorpFactory) CreatePlatformCorp(name string) (*Corp, error) {
	return this.createCorp(name, true)
}

func init() {
}
