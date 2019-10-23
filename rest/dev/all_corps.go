package dev

import (
	b_corp "github.com/gingerxman/ginger-account/business/corp"
	
	"github.com/gingerxman/eel"
)

type AllCorps struct {
	eel.RestResource
}

func (this *AllCorps) Resource() string {
	return "dev.all_corps"
}

func (this *AllCorps) SkipAuthCheck() bool {
	return true
}

func (r *AllCorps) IsForDevTest() bool {
	return true
}

func (this *AllCorps) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{},
	}
}

func (this *AllCorps) Get(ctx *eel.Context) {
	bCtx := ctx.GetBusinessContext()
	allCorps := b_corp.NewCorpRepository(bCtx).GetCorps(eel.Map{
		"id__gt": 0,
	})
	
	corps := make([]*b_corp.Corp, 0)
	for _, corp := range allCorps {
		if corp.Name != "Ginger Corp" {
			corps = append(corps, corp)
		}
	}

	b_corp.NewFillCorpService(bCtx).Fill(corps, eel.FillOption{
		"with_corp_user": true,
	})
	rows := b_corp.NewEncodeCorpService(bCtx).EncodeMany(corps)
	
	ctx.Response.JSON(eel.Map{
		"corps": rows,
	})
}