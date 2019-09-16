package corp

import (
	"context"
	"github.com/gingerxman/eel"
)

type FillCorpService struct {
	eel.ServiceBase
}

func NewFillCorpService(ctx context.Context) *FillCorpService {
	service := new(FillCorpService)
	service.Ctx = ctx
	return service
}

func (this *FillCorpService) FillOne(corp *Corp, option eel.FillOption) {
	this.Fill([]*Corp{corp}, option)
}

func (this *FillCorpService) Fill(corps []*Corp, option eel.FillOption)  {
	if len(corps) == 0 {
		return
	}
	
	corpIds := make([]int, 0)
	for _, corp := range corps {
		corpIds = append(corpIds, corp.Id)
	}
}

func init() {
}
