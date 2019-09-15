package user

import (
	"context"
	"github.com/gingerxman/eel"
	m_user "github.com/gingerxman/ginger-account/models/user"
)

type UserFactory struct {
	eel.ServiceBase
}

func NewUserFactory(ctx context.Context) *UserFactory {
	service := new(UserFactory)
	service.Ctx = ctx
	return service
}

func(this *UserFactory) CreateUser(name string, password string, unionid string) *User{
	userDbModel := &m_user.User{
		Sex: m_user.USER_SEX_UNKNOWN,
		Unionid: unionid,
		Password: password,
		Name: name,
	}
	o := eel.GetOrmFromContext(this.Ctx)
	db := o.Create(userDbModel)
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		panic(eel.NewBusinessError("user_factory:create_user_fail", "创建User失败"))
	}
	
	//userDbModel.Code = fmt.Sprintf("%d", constant.SNS_MAGIC_NUMBER+userDbModel.Id)
	//_, err = o.Update(userDbModel, "Code")
	//if err != nil {
	//	beego.Error(err)
	//	panic(vanilla.NewBusinessError("user_factory:update_user_code_fail", "更新User的Code失败"))
	//}
	user := NewUserFromModel(this.Ctx, userDbModel)
	return user
}

func init() {
}
