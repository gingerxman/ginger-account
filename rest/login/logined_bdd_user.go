package login

import (
	"github.com/gingerxman/ginger-account/business/corp"
	b_user "github.com/gingerxman/ginger-account/business/user"
	
	"github.com/gingerxman/eel"
)

type LoginedBDDUser struct {
	eel.RestResource
}

func (this *LoginedBDDUser) Resource() string {
	return "login.logined_bdd_user"
}

func (this *LoginedBDDUser) SkipAuthCheck() bool {
	return true
}

func (this *LoginedBDDUser) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT": []string{"username", "type"},
	}
}

func (this *LoginedBDDUser) Put(ctx *eel.Context) {
	req := ctx.Request
	username := req.GetString("username")
	userType := req.GetString("type")
	
	bCtx := ctx.GetBusinessContext()
	if userType == "user" {
		user := b_user.NewUserRepository(bCtx).GetUserByUnionid(username)
		if user == nil {
			user = b_user.NewUserFactory(bCtx).CreateUser(username, "test", username, "bdd")
		}
		
		ctx.Response.JSON(eel.Map{
			"id": user.Id,
			"jwt": user.GetJWTToken(),
		})
	} else if userType == "corp_user" {
		corpUser, err := corp.NewCorpUserService(bCtx).Auth(username, corp.SUPER_PASSWORD)
		if err != nil {
			ctx.Response.Error("logined_corp_user:login_fail", err.Error())
			return
		}
		
		ctx.Response.JSON(eel.Map{
			"id": corpUser.Id,
			"uid": corpUser.Id,
			"cid": corpUser.CorpId,
			"jwt": corpUser.GetJWTToken(),
		})
	}
}

