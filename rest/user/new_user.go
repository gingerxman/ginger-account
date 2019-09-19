package user

import (
	b_user "github.com/gingerxman/ginger-account/business/user"
	
	"github.com/gingerxman/eel"
)

type NewUser struct {
	eel.RestResource
}

func (this *NewUser) Resource() string {
	return "user.new_user"
}

func (this *NewUser) SkipAuthCheck() bool {
	return true
}

func (this *NewUser) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT": []string{"name", "password", "?unionid"},
	}
}

func (this *NewUser) Put(ctx *eel.Context) {
	req := ctx.Request
	name := req.GetString("name")
	password := req.GetString("password")
	unionid := req.GetString("unionid", name)

	bCtx := ctx.GetBusinessContext()
	user := b_user.NewUserFactory(bCtx).CreateUser(name, password, unionid, "web")
	
	ctx.Response.JSON(eel.Map{
		"id": user.Id,
	})
}

