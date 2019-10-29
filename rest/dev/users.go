package dev

import (
	"github.com/gingerxman/ginger-account/business/user"
	
	"github.com/gingerxman/eel"
)

type Users struct {
	eel.RestResource
}

func (this *Users) Resource() string {
	return "dev.users"
}

func (this *Users) SkipAuthCheck() bool {
	return true
}

func (r *Users) IsForDevTest() bool {
	return true
}

func (this *Users) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{},
	}
}

func (this *Users) Get(ctx *eel.Context) {
	bCtx := ctx.GetBusinessContext()
	users := user.NewUserRepository(bCtx).GetUsers(eel.Map{
		"id__gt": 0,
	})
	
	user.NewFillUserService(bCtx).Fill(users, eel.FillOption{})
	datas := user.NewEncodeUserService(bCtx).EncodeMany(users)
	
	ctx.Response.JSON(eel.Map{
		"users": datas,
	})
}