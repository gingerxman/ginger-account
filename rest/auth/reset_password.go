package auth

import (
	"fmt"
	"github.com/gingerxman/ginger-account/business/corp"
	
	"github.com/gingerxman/eel"
)

/*
 @resource: 角色
 */

type ResetPassword struct {
	eel.RestResource
}

func (this *ResetPassword) Resource() string {
	return "auth.reset_password"
}

func (this *ResetPassword) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT": []string{"corp_user_id:int"},
	}
}

func (this *ResetPassword) Put(ctx *eel.Context) {
	req := ctx.Request
	corpUserId, _ := req.GetInt("corp_user_id")
	
	bCtx := ctx.GetBusinessContext()
	corpUser := corp.NewCorpUserRepository(bCtx).GetCorpUserById(corpUserId)
	if corpUser == nil {
		ctx.Response.Error("reset_password:invalid_corp_user", fmt.Sprintf("无效的id(%d)", corpUserId))
		return
	}
	
	newPassword, err := corpUser.ResetPassword()
	if err != nil {
		ctx.Response.Error("reset_password:reset_fail", err.Error())
		return
	} else {
		ctx.Response.JSON(eel.Map{
			"password": newPassword,
		})
	}
}

