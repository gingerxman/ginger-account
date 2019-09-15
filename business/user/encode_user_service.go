package user

import (
	"context"
	"github.com/gingerxman/eel"
)

type EncodeUserService struct {
	eel.ServiceBase
}

func NewEncodeUserService(ctx context.Context) *EncodeUserService {
	service := new(EncodeUserService)
	service.Ctx = ctx
	return service
}

func (this *EncodeUserService) Encode(user *User) *RUser {
	rUser := &RUser{
		Id: user.Id,
		Code: user.Code,
		
		//微信相关信息
		Unionid: user.Unionid,
		Name: user.Name,
		Avatar: user.Avatar,
		Sex: user.Sex,
	}
	
	return rUser
}

func (this *EncodeUserService) EncodeMany(users []*User) []*RUser {
	rows := make([]*RUser, 0)
	for _, user := range users {
		rows = append(rows, this.Encode(user))
	}
	
	return rows
}

func init() {
}
