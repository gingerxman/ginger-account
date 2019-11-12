package corp

import (
	"context"
	"github.com/gingerxman/eel"
)

type EncodeCorpUserService struct {
	eel.ServiceBase
}

func NewEncodeCorpUserService(ctx context.Context) *EncodeCorpUserService {
	service := new(EncodeCorpUserService)
	service.Ctx = ctx
	return service
}

func (this *EncodeCorpUserService) Encode(corpUser *CorpUser) *RCorpUser {
	return &RCorpUser{
		Id: corpUser.Id,
		Name: corpUser.Username,
	}
}

func (this *EncodeCorpUserService) EncodeMany(corpUsers []*CorpUser) []*RCorpUser {
	rows := make([]*RCorpUser, 0)
	for _, corpUser := range corpUsers {
		rows = append(rows, this.Encode(corpUser))
	}
	
	return rows
}

func init() {
}
