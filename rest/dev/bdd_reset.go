package dev

import (
	"fmt"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-account/business/corp"
	"github.com/gingerxman/ginger-account/business/user"
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
	
	gingerCorp := corp.NewCorpRepository(bCtx).GetCorpByName("Ginger Corp")
	if gingerCorp == nil {
		var err error
		gingerCorp, err = corp.NewCorpFactory(bCtx).CreatePlatformCorp("Ginger Corp")
		if err == nil {
			user.NewUserFactory(bCtx).CreateUserForCorp(gingerCorp)
			gingerCorp.AddCorpUser("ginger", "test")
		}
		
		o.Exec("delete from user_user")
	} else {
		// 删除除了ginger corp之外的user
		_id := struct {
			Id int
		}{}
		//gingerCorpId := 0
		db := o.Raw("select id from corp_corp where name = 'Ginger Corp'").Scan(&_id)
		if db.Error != nil {
			panic(db.Error)
		}
		name := fmt.Sprintf("corp_%d", _id.Id)
		o.Exec("delete from user_user where name != ?", name)
	}
	
	ctx.Response.JSON(eel.Map{})
}

