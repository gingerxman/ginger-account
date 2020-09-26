package user

import (
	"fmt"
	b_user "github.com/gingerxman/ginger-account/business/user"

	"github.com/gingerxman/eel"
)

type User struct {
	eel.RestResource
}

func (this *User) Resource() string {
	return "user.user"
}

func (this *User) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"id:int"},
		"POST": []string{"id:int", "?name", "?avatar", "?sex"},
		"DELETE": []string{"id:int"},
	}
}

func (this *User) Get(ctx *eel.Context) {
	req := ctx.Request
	id, _ := req.GetInt("id")

	bCtx := ctx.GetBusinessContext()
	user := b_user.NewUserRepository(bCtx).GetUserById(id)
	
	if user == nil {
		ctx.Response.Error("user:invalid_user", fmt.Sprintf("id=%d", id))
		return
	}

	b_user.NewFillUserService(bCtx).FillOne(user, eel.FillOption{})
	respData := b_user.NewEncodeUserService(bCtx).Encode(user)
	
	ctx.Response.JSON(respData)
}

func (this *User) Post(ctx *eel.Context) {
	req := ctx.Request
	id, _ := req.GetInt("id")
	name := req.GetString("name")
	avatar := req.GetString("avatar")
	sex := req.GetString("sex", "")
	
	bCtx := ctx.GetBusinessContext()
	user := b_user.NewUserRepository(bCtx).GetUserById(id)
	if user == nil {
		ctx.Response.Error("user:invalid_user", fmt.Sprintf("id=%d", id))
		return
	}
	
	err := user.Update(&b_user.UpdateUserParams{
		Name: name,
		Avatar: avatar,
		Sex: sex,
	})
	if err != nil {
		ctx.Response.Error("user:update_fail", err.Error())
	} else {
		ctx.Response.JSON(eel.Map{})
	}
}

func (this *User) Delete(ctx *eel.Context) {
	req := ctx.Request
	id, _ := req.GetInt("id")

	bCtx := ctx.GetBusinessContext()
	user := b_user.NewUserRepository(bCtx).GetUserById(id)
	if user == nil {
		ctx.Response.Error("user:invalid_user", fmt.Sprintf("id=%d", id))
		return
	}
	
	err := user.Delete()
	if err != nil {
		ctx.Response.Error("user:delete_fail", err.Error())
	} else {
		ctx.Response.JSON(eel.Map{})
	}
}
