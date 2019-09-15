package login

import (
	"fmt"
	"github.com/gingerxman/ginger-account/business/user"
	
	"github.com/gingerxman/eel"
)

const SECRET = "55e421ee9bdc0fda6b6c6518e590b0ee"

type LoginedUser struct {
	eel.RestResource
}

func (this *LoginedUser) Resource() string {
	return "login.logined_user"
}

func (this *LoginedUser) SkipAuthCheck() bool {
	return true
}

func (this *LoginedUser) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT": []string{"unionid", "secret"},
	}
}

func (this *LoginedUser) Put(ctx *eel.Context) {
	req := ctx.Request
	unionid := req.GetString("unionid")
	secret := req.GetString("secret")
	
	if secret != SECRET {
		ctx.Response.Error("logined_user:login_fail", "")
		return
	}

	bCtx := ctx.GetBusinessContext()
	user := user.NewUserRepository(bCtx).GetUserByUnionid(unionid)
	if user == nil {
		ctx.Response.Error("logined_user:invalid_user", fmt.Sprintf("unionid(%s)", unionid))
		return
	}
	
	ctx.Response.JSON(eel.Map{
		"id": user.Id,
		"jwt": user.GetJWTToken(),
	})
}

