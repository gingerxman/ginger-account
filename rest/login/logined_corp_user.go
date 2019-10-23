package login

import (
	"github.com/gingerxman/ginger-account/business/corp"
	
	"github.com/gingerxman/eel"
)

type LoginedCorpUser struct {
	eel.RestResource
}

func (this *LoginedCorpUser) Resource() string {
	return "login.logined_corp_user"
}

func (this *LoginedCorpUser) SkipAuthCheck() bool {
	return true
}

func (this *LoginedCorpUser) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT": []string{"username", "password"},
	}
}

func (this *LoginedCorpUser) Put(ctx *eel.Context) {
	req := ctx.Request
	username := req.GetString("username")
	password := req.GetString("password")

	bCtx := ctx.GetBusinessContext()
	corpUser, err := corp.NewCorpUserService(bCtx).Auth(username, password)
	if err != nil {
		ctx.Response.Error("logined_corp_user:login_fail", err.Error())
		return
	}
	
	corp := corp.NewCorpRepository(bCtx).GetCorpById(corpUser.CorpId)
	
	ctx.Response.JSON(eel.Map{
		"id": corpUser.Id,
		"uid": corpUser.Id,
		"cid": corpUser.CorpId,
		"username": corpUser.Username,
		"corp_name": corp.Name,
		"jwt": corpUser.GetJWTToken(),
	})
}

