package login

import (
	b_user "github.com/gingerxman/ginger-account/business/user"
	
	"github.com/gingerxman/eel"
)

type MallVisitor struct {
	eel.RestResource
}

func (this *MallVisitor) Resource() string {
	return "login.mall_visitor"
}

func (this *MallVisitor) SkipAuthCheck() bool {
	return true
}

func (this *MallVisitor) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT": []string{"name", "avatar", "unionid", "corp_id:int"},
	}
}

func (this *MallVisitor) Put(ctx *eel.Context) {
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

