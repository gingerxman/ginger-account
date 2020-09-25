package auth

import (
	b_auth "github.com/gingerxman/ginger-account/business/auth"
	
	"github.com/gingerxman/eel"
)

/*
 @resource: 权限
 */

type Permission struct {
	eel.RestResource
}

func (this *Permission) Resource() string {
	return "auth.permission"
}

func (this *Permission) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT": []string{"code", "name"},
	}
}

func (this *Permission) Put(ctx *eel.Context) {
	req := ctx.Request
	code := req.GetString("code")
	name := req.GetString("name")
	
	bCtx := ctx.GetBusinessContext()
	permission, err := b_auth.NewPermission(bCtx, code, name)
	if err != nil {
		eel.Logger.Error(err)
		ctx.Response.Error("permission:create_fail", "create fail")
		return
	}
	
	ctx.Response.JSON(eel.Map{
		"id": permission.Id,
	})
}

