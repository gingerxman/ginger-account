package dev

import (
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-account/business/corp"
)

type BDDReset struct {
	eel.RestResource
}

func (this *BDDReset) Resource() string {
	return "dev.bdd_reset"
}

func (this *BDDReset) SkipAuthCheck() bool {
	return true
}

func (r *BDDReset) IsForDevTest() bool {
	return true
}

func (this *BDDReset) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT":  []string{},
	}
}

func (this *BDDReset) Put(ctx *eel.Context) {
	bCtx := ctx.GetBusinessContext()
	o := eel.GetOrmFromContext(bCtx)
	
	o.Exec("delete from corp_corp where name != 'Ginger Corp'")
	o.Exec("delete from corp_user where username != 'ginger'")
	
	o.Exec("delete from user_user")
	
	corp, err := corp.NewCorpFactory(bCtx).CreatePlatformCorp("Ginger Corp")
	if err == nil {
		corp.AddCorpUser("ginger", "test")
	}
	
	ctx.Response.JSON(eel.Map{})
}

