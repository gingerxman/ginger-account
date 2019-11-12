package corp

import (
	b_corp "github.com/gingerxman/ginger-account/business/corp"
	
	"github.com/gingerxman/eel"
)

type CorpUsers struct {
	eel.RestResource
}

func (this *CorpUsers) Resource() string {
	return "corp.corp_users"
}

func (this *CorpUsers) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"ids:json-array"},
	}
}

func (this *CorpUsers) Get(ctx *eel.Context) {
	req := ctx.Request
	ids := req.GetIntArray("ids")

	bCtx := ctx.GetBusinessContext()
	corpUsers := b_corp.NewCorpUserRepository(bCtx).GetCorpUsersByIds(ids)
	datas := b_corp.NewEncodeCorpUserService(bCtx).EncodeMany(corpUsers)
	
	ctx.Response.JSON(eel.Map{
		"corp_users": datas,
	})
}
