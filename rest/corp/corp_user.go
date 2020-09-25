package corp

import (
	"fmt"
	b_corp "github.com/gingerxman/ginger-account/business/corp"
	"math/rand"
	"strings"
	"time"
	
	"github.com/gingerxman/eel"
)

type CorpUser struct {
	eel.RestResource
}

func (this *CorpUser) Resource() string {
	return "corp.corp_user"
}

func (this *CorpUser) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"?id:int"},
		"PUT": []string{"username", "realname", "?password", "group_names:json-array"},
		"POST": []string{"id:int", "realname", "group_names:json-array"},
		"DELETE": []string{"id:int"},
	}
}

func (this *CorpUser) Get(ctx *eel.Context) {
	req := ctx.Request
	id, _ := req.GetInt("id", 0)

	bCtx := ctx.GetBusinessContext()
	var corpUser *b_corp.CorpUser
	if id == 0 {
		corp := b_corp.GetCorpFromContext(bCtx)
		id = corp.CorpUser.Id
	}
	corpUser = b_corp.NewCorpUserRepository(bCtx).GetCorpUserById(id)
	
	if corpUser == nil {
		ctx.Response.Error("corp_user:invalid_corp_user", fmt.Sprintf("id=%d", id))
		return
	}
	
	
	b_corp.NewFillCorpUserService(bCtx).FillOne(corpUser, eel.FillOption{
		"with_permission": true,
	})
	respData := b_corp.NewEncodeCorpUserService(bCtx).Encode(corpUser)
	
	ctx.Response.JSON(respData)
}

func (this *CorpUser) Put(ctx *eel.Context) {
	req := ctx.Request
	bCtx := ctx.GetBusinessContext()
	
	// 生成密码
	password := req.GetString("password", "")
	if password == "" {
		rand.Seed( time.Now().UTC().UnixNano())
		buf := make([]string, 0)
		for i := 1; i < 7; i++ {
			s := fmt.Sprintf("%d", rand.Intn(10))
			buf = append(buf, s)
		}
		password = strings.Join(buf, "")
	}
	
	// 创建corp相关的user
	corp := b_corp.GetCorpFromContext(bCtx)
	username := req.GetString("username")
	realname := req.GetString("realname")
	groupNames := req.GetStringArray("group_names")
	err := corp.AddCorpUser(username, password, realname, groupNames, false)
	if err != nil {
		eel.Logger.Error(err)
		ctx.Response.Error("corp_user:create_fail", err.Error())
		return
	} else {
		ctx.Response.JSON(eel.Map{
			"password": password,
		})
	}
}

func (this *CorpUser) Post(ctx *eel.Context) {
	req := ctx.Request
	bCtx := ctx.GetBusinessContext()
	
	id, _ := req.GetInt("id")
	corp := b_corp.GetCorpFromContext(bCtx)
	corpUser := b_corp.NewCorpUserRepository(bCtx).GetCorpUserInCorp(id, corp)
	
	if corpUser == nil {
		ctx.Response.Error("corp_user:invalid_corp_user", fmt.Sprintf("id=%d, corp=%d", id, corp.GetId()))
		return
	}
	
	groupNames := req.GetStringArray("group_names")
	realname := req.GetString("realname")
	corpUser.Update(realname, groupNames)
	
	ctx.Response.JSON(eel.Map{
	})
}

func (this *CorpUser) Delete(ctx *eel.Context) {
	req := ctx.Request
	id, _ := req.GetInt("id")

	bCtx := ctx.GetBusinessContext()
	corp := b_corp.GetCorpFromContext(bCtx)
	if corp == nil {
		ctx.Response.Error("corp_user:invalid_corp", fmt.Sprintf("corp_id(%d)", id))
	} else {
		ctx.Response.JSON(eel.Map{})
	}
}
