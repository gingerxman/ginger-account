package user

import (
	"context"
	"github.com/gingerxman/eel"
)

type FillUserService struct {
	eel.ServiceBase
}

func NewFillUserService(ctx context.Context) *FillUserService {
	service := new(FillUserService)
	service.Ctx = ctx
	return service
}

func (this *FillUserService) FillOne(user *User, option eel.FillOption) {
	this.Fill([]*User{user}, option)
}

func (this *FillUserService) Fill(users []*User, option eel.FillOption)  {
	if len(users) == 0 {
		return
	}
	
	userIds := make([]int, 0)
	for _, user := range users {
		userIds = append(userIds, user.Id)
	}
}

func init() {
}
