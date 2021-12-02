package login

import (
	b_user "github.com/gingerxman/ginger-account/business/user"
	
	"github.com/gingerxman/eel"
)

type LoginedMallUser struct {
	eel.RestResource
}

func (this *LoginedMallUser) Resource() string {
	return "login.logined_mall_user"
}

func (this *LoginedMallUser) SkipAuthCheck() bool {
	return true
}

func (this *LoginedMallUser) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT": []string{"name", "unionid", "?avatar", "corp_id:int"},
	}
}

func (this *LoginedMallUser) Put(ctx *eel.Context) {
	req := ctx.Request
	unionid := req.GetString("unionid")
	name := req.GetString("name")
	avatar := req.GetString("avatar")
	corpId, _ := req.GetInt("corp_id")

	bCtx := ctx.GetBusinessContext()
	user := b_user.NewUserRepository(bCtx).GetUserByUnionid(unionid)
	if user == nil {
		user = b_user.NewUserFactory(bCtx).CreateUser(name, "nopassword", unionid, "h5")
		user.Update(&b_user.UpdateUserParams{
			Avatar: avatar,
		})
	}
	
	ctx.Response.JSON(eel.Map{
		"id": user.Id,
		"jwt": user.GetJWTTokenInCorp(corpId),
	})
}

