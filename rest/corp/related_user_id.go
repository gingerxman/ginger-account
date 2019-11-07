package corp

import (
	b_corp "github.com/gingerxman/ginger-account/business/corp"
	"github.com/gingerxman/ginger-account/business/user"
	
	"github.com/gingerxman/eel"
)

type RelatedUserId struct {
	eel.RestResource
}

func (this *RelatedUserId) Resource() string {
	return "corp.related_user_id"
}

func (this *RelatedUserId) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"corp_id:int"},
	}
}

func (this *RelatedUserId) Get(ctx *eel.Context) {
	req := ctx.Request
	corpId, _ := req.GetInt("corp_id")
	bCtx := ctx.GetBusinessContext()
	corp := b_corp.NewCorpFromId(bCtx, corpId)
	
	ctx.Response.JSON(eel.Map{
		"id": user.NewUserRepository(bCtx).GetRelatedUserIdForCorp(corp),
	})
}
