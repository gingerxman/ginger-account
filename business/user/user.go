package user

import (
	"context"
	"fmt"
	"github.com/gingerxman/ginger-account/business/constant"
	"github.com/gingerxman/gorm"
	
	"time"
	
	"github.com/gingerxman/eel"
	m_user "github.com/gingerxman/ginger-account/models/user"
)

type UpdateUserParams struct {
	Name string
	Avatar string
}

type User struct {
	eel.EntityBase
	Id int
	Unionid string
	IsActive bool
	Name string
	Avatar string
	Sex string
	Code string
	CreatedAt time.Time
}

func NewUserFromOnlyId(ctx context.Context, id int) *User {
	user := new(User)
	user.Ctx = ctx
	user.Model = nil
	user.Id = id
	return user
}

func (this *User) GetId() int {
	return this.Id
}

func (this *User) Fill() {
	if this.Code != "" {
		return
	}
	
	model := m_user.User{}
	db := eel.GetOrmFromContext(this.Ctx).Model(&m_user.User{}).Where("id", this.Id).Take(&model)
	
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		panic(eel.NewBusinessError("user:invalid_user", fmt.Sprintf("无效user(%d)", this.Id)))
	}
	
	this.FillWithModel(&model)
}

func (this *User) FillWithModel(model *m_user.User) {
	if this.Code != "" {
		return
	}
	
	user := this
	user.Model = model
	user.Id = model.Id
	user.Code = model.Code
	user.Unionid = model.Unionid
	user.Name = model.Name
	user.Avatar = model.Thumbnail
	user.Sex = "unknown"
	if model.Sex == m_user.USER_SEX_MALE {
		user.Sex = "male"
	} else if model.Sex == m_user.USER_SEX_FEMALE {
		user.Sex = "female"
	}
	
	user.CreatedAt = model.CreatedAt
}

func (this *User) Delete() error {
	o := eel.GetOrmFromContext(this.Ctx)
	
	db := o.Model(&m_user.User{}).Where("id", this.Id).Update(gorm.Params{
		"is_active": false,
	})
	
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return db.Error
	}
	
	return nil
}

func (this *User) Update(params *UpdateUserParams) error {
	gormParams := gorm.Params{}
	if params.Name != "" {
		gormParams["name"] = params.Name
	}
	if params.Avatar != "" {
		gormParams["avatar"] = params.Avatar
	}
	
	db := eel.GetOrmFromContext(this.Ctx).Model(&m_user.User{}).Where("id", this.Id).Update(gormParams)
	
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return db.Error
	}
	
	return nil
}

func (this *User) GetJWTToken() string {
	return eel.EncodeJWT(this.GetJWTTokenData())
}

func (this *User) GetJWTTokenInCorp(corpId int) string {
	data := this.GetJWTTokenData()
	data["cid"] = corpId
	return eel.EncodeJWT(data)
}

func (this *User) GetJWTTokenData() eel.Map {
	this.Fill()
	data := eel.Map {
		"v": constant.USER_JWT_VERSION,
		"uid": this.Id,
		"code": this.Code,
		"type": 2,
	}
	
	return data
}

func GetUserFromContext(ctx context.Context) *User {
	user := ctx.Value("user").(*User)
	return user
}

func NewUserFromModel(ctx context.Context, model *m_user.User) *User {
	user := new(User)
	user.Ctx = ctx
	user.FillWithModel(model)
	return user
}

func init() {
}
