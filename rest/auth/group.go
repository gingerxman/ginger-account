package auth

import (
	b_auth "github.com/gingerxman/ginger-account/business/auth"
	
	"github.com/gingerxman/eel"
)

/*
 @resource: 角色
 */

type Group struct {
	eel.RestResource
}

func (this *Group) Resource() string {
	return "auth.group"
}

func (this *Group) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT": []string{"code", "name", "permissions:json-array"},
	}
}

func (this *Group) Put(ctx *eel.Context) {
	req := ctx.Request
	code := req.GetString("code")
	name := req.GetString("name")
	permissionCodes := req.GetStringArray("permissions")
	
	bCtx := ctx.GetBusinessContext()
	group, err := b_auth.NewGroup(bCtx, code, name, permissionCodes)
	if err != nil {
		eel.Logger.Error(err)
		ctx.Response.Error("group:create_fail", "create fail")
		return
	}
	
	ctx.Response.JSON(eel.Map{
		"id": group.Id,
	})
}

