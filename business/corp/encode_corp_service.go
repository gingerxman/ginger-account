package corp

import (
	"context"
	"github.com/gingerxman/eel"
)

type EncodeCorpService struct {
	eel.ServiceBase
}

func NewEncodeCorpService(ctx context.Context) *EncodeCorpService {
	service := new(EncodeCorpService)
	service.Ctx = ctx
	return service
}

func (this *EncodeCorpService) Encode(corp *Corp) *RCorp {
	rCorp := &RCorp{
		Id: corp.Id,
		Code: corp.Code,
		Name: corp.Name,
		Logo: corp.Logo,
		Remark: corp.Remark,
		IsPlatform: corp.IsPlatform,
		CreatedAt: corp.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	
	return rCorp
}

func (this *EncodeCorpService) EncodeMany(corps []*Corp) []*RCorp {
	rows := make([]*RCorp, 0)
	for _, corp := range corps {
		rows = append(rows, this.Encode(corp))
	}
	
	return rows
}

func init() {
}
