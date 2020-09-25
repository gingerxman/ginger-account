package corp

import (
	"context"
	"github.com/gingerxman/ginger-account/business/auth"
	"github.com/gingerxman/eel"
)

type FillCorpUserService struct {
	eel.ServiceBase
}

func NewFillCorpUserService(ctx context.Context) *FillCorpUserService {
	service := new(FillCorpUserService)
	service.Ctx = ctx
	return service
}

func (this *FillCorpUserService) FillOne(user *CorpUser, option eel.FillOption) {
	this.Fill([]*CorpUser{user}, option)
}

func (this *FillCorpUserService) Fill(users []*CorpUser, option eel.FillOption)  {
	if len(users) == 0 {
		return
	}
	
	userIds := make([]int, 0)
	for _, user := range users {
		userIds = append(userIds, user.Id)
	}
	
	this.fillCorp(users, userIds)
	if enableOption, ok := option["with_permission"]; ok && enableOption {
		this.fillPermission(users, userIds)
	}
}

func (this *FillCorpUserService) fillPermission(users []*CorpUser, userIds []int) {
	for _, user := range users {
		groups := auth.NewGroupRepository(this.Ctx).GetGroupsForCorpUser(user.Id)
		
		auth.NewFillGroupService(this.Ctx).Fill(groups, eel.FillOption{
			"with_permission": true,
		})
		
		user.Groups = groups
	}
}

func (this *FillCorpUserService) fillCorp(users []*CorpUser, userIds []int) {
	corpIds := make([]int, 0)
	for _, user := range users {
		corpIds = append(corpIds, user.CorpId)
	}
	
	corps := NewCorpRepository(this.Ctx).GetCorpsByIds(corpIds)
	id2corp := make(map[int]*Corp)
	for _, corp := range corps {
		id2corp[corp.Id] = corp
	}
	
	for _, user := range users {
		if corp, ok := id2corp[user.CorpId]; ok {
			user.Corp = corp
		}
	}
}

func init() {
}
