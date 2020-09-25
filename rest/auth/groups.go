package auth

import (
	"github.com/gingerxman/ginger-account/business/auth"
	
	"github.com/gingerxman/eel"
)

type Groups struct {
	eel.RestResource
}

func (this *Groups) Resource() string {
	return "auth.groups"
}

func (this *Groups) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{},
	}
}

func (this *Groups) Get(ctx *eel.Context) {
	bCtx := ctx.GetBusinessContext()
	groups := auth.NewGroupRepository(bCtx).GetGroups(eel.Map{})
	
	auth.NewFillGroupService(bCtx).Fill(groups, eel.FillOption{})
	datas := auth.NewEncodeGroupService(bCtx).EncodeMany(groups)
	
	ctx.Response.JSON(eel.Map{
		"groups": datas,
	})
}
