package corp

import (
	"fmt"
	b_corp "github.com/gingerxman/ginger-account/business/corp"

	"github.com/gingerxman/eel"
)

type Corp struct {
	eel.RestResource
}

func (this *Corp) Resource() string {
	return "corp.corp"
}

func (this *Corp) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"id:int"},
		"PUT": []string{"corp_name", "username", "password"},
		"POST": []string{"id:int", "name"},
		"DELETE": []string{"id:int"},
	}
}

func (this *Corp) Get(ctx *eel.Context) {
	req := ctx.Request
	id, _ := req.GetInt("id")

	bCtx := ctx.GetBusinessContext()
	corp := b_corp.NewCorpRepository(bCtx).GetCorpById(id)
	
	if corp == nil {
		ctx.Response.Error("corp:invalid_corp", fmt.Sprintf("id=%d", id))
		return
	}

	b_corp.NewFillCorpService(bCtx).FillOne(corp, eel.FillOption{})
	respData := b_corp.NewEncodeCorpService(bCtx).Encode(corp)
	
	ctx.Response.JSON(respData)
}

func (this *Corp) Put(ctx *eel.Context) {
	req := ctx.Request
	corpName := req.GetString("corp_name")
	
	bCtx := ctx.GetBusinessContext()
	corp, err := b_corp.NewCorpFactory(bCtx).CreateCorp(corpName)
	if err != nil {
		ctx.Response.Error("corp:create_fail", err.Error())
		return
	}
	
	username := req.GetString("username")
	password := req.GetString("password")
	err = corp.AddCorpUser(username, password)
	if err != nil {
		ctx.Response.Error("corp:create_fail", err.Error())
		return
	}
	
	ctx.Response.JSON(eel.Map{})
}

func (this *Corp) Post(ctx *eel.Context) {
	//req := ctx.Request
	//id, _ := req.GetInt("id")
	//title := req.GetString("title")
	//content := req.GetString("content")
	//
	//bCtx := ctx.GetBusinessContext()
	//blogRepository := b_blog.NewBlogRepository(bCtx)
	//blog := blogRepository.GetBlog(id)
	//
	//blog.Update(title, content)

	ctx.Response.JSON(eel.Map{})
}

func (this *Corp) Delete(ctx *eel.Context) {
	req := ctx.Request
	id, _ := req.GetInt("id")

	bCtx := ctx.GetBusinessContext()
	corp := b_corp.NewCorpRepository(bCtx).GetCorpById(id)
	corp.Delete()
	if corp == nil {
		ctx.Response.Error("corp:invalid_corp", fmt.Sprintf("corp_id(%d)", id))
	} else {
		ctx.Response.JSON(eel.Map{})
	}
}
