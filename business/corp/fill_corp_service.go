package corp

import (
	"context"
	"github.com/gingerxman/eel"
	m_corp "github.com/gingerxman/ginger-account/models/corp"
)

type FillCorpService struct {
	eel.ServiceBase
}

func NewFillCorpService(ctx context.Context) *FillCorpService {
	service := new(FillCorpService)
	service.Ctx = ctx
	return service
}

func (this *FillCorpService) FillOne(corp *Corp, option eel.FillOption) {
	this.Fill([]*Corp{corp}, option)
}

func (this *FillCorpService) Fill(corps []*Corp, option eel.FillOption)  {
	if len(corps) == 0 {
		return
	}
	
	corpIds := make([]int, 0)
	for _, corp := range corps {
		corpIds = append(corpIds, corp.Id)
	}
	if enableOption, ok := option["with_corp_user"]; ok && enableOption {
		this.fillCorpUser(corps, corpIds)
	}
}

func (this *FillCorpService) fillCorpUser(corps []*Corp, ids []int) {
	//构造<id, corp>
	id2corp := make(map[int]*Corp)
	for _, corp := range corps {
		id2corp[corp.Id] = corp
	}
	
	//填充corp.CorpUser
	var models []*m_corp.CorpUser
	eel.GetOrmFromContext(this.Ctx).Model(&m_corp.CorpUser{}).Where("corp_id__in", ids).Find(&models)
	
	for _, model := range models {
		if corp, ok := id2corp[model.CorpId]; ok {
			corp.CorpUser = NewCorpUserFromModel(this.Ctx, model)
			
		}
	}
}

func init() {
}
