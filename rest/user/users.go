package user

import (
	"github.com/gingerxman/ginger-account/business/user"
	
	"github.com/gingerxman/eel"
)

type Users struct {
	eel.RestResource
}

func (this *Users) Resource() string {
	return "user.users"
}

func (this *Users) SkipAuthCheck() bool {
	return true
}

func (r *Users) IsForDevTest() bool {
	return true
}

func (this *Users) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"ids:json-array"},
	}
}

func (this *Users) Get(ctx *eel.Context) {
	bCtx := ctx.GetBusinessContext()
	req := ctx.Request
	userIds := req.GetIntArray("ids")
	users := user.NewUserRepository(bCtx).GetUsersByIds(userIds)
	
	user.NewFillUserService(bCtx).Fill(users, eel.FillOption{})
	datas := user.NewEncodeUserService(bCtx).EncodeMany(users)
	
	ctx.Response.JSON(eel.Map{
		"users": datas,
	})
}