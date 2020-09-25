package corp

import (
	"context"
	"fmt"
	"github.com/gingerxman/ginger-account/business/auth"
	"github.com/gingerxman/ginger-account/business/constant"
	"github.com/gingerxman/gorm"
	"math/rand"
	"strings"
	
	"time"
	
	"github.com/gingerxman/eel"
	m_corp "github.com/gingerxman/ginger-account/models/corp"
)


type CorpUser struct {
	eel.EntityBase
	Id int
	CorpId int
	IsActive bool
	IsManager bool
	Username string
	RealName string
	CreatedAt time.Time
	
	Groups []*auth.Group
	Corp *Corp
}

func (this *CorpUser) Update(realname string, groupNames []string) {
	o := eel.GetOrmFromContext(this.Ctx)
	
	db := o.Model(&m_corp.CorpUser{}).Where("id", this.Id).Update(gorm.Params{
		"real_name": realname,
	})
	if db.Error != nil {
		eel.Logger.Error(db.Error)
	}
	
	o.Where("corp_user_id", this.Id).Delete(&m_corp.UserHasGroup{})
	if len(groupNames) > 0 {
		groups := auth.NewGroupRepository(this.Ctx).GetGroupsByNames(groupNames)
		for _, group := range groups {
			model := &m_corp.UserHasGroup{
				CorpUserId: this.Id,
				GroupId: group.Id,
			}
			db := o.Create(model)
			if db.Error != nil {
				eel.Logger.Error(db.Error)
			}
		}
	}
}

func (this *CorpUser) GetJWTTokenData() eel.Map {
	data := eel.Map {
		"v": constant.CORP_USER_JWT_VERSION,
		"uid": this.Id,
		"cid": this.CorpId,
		"name": this.RealName,
		"type": 1,
	}
	
	return data
}

func (this *CorpUser) GetJWTToken() string {
	return eel.EncodeJWT(this.GetJWTTokenData())
}

func (this *CorpUser) ResetPassword () (string, error) {
	o := eel.GetOrmFromContext(this.Ctx)
	
	// 生成密码
	rand.Seed( time.Now().UTC().UnixNano())
	buf := make([]string, 0)
	for i := 1; i < 7; i++ {
		s := fmt.Sprintf("%d", rand.Intn(10))
		buf = append(buf, s)
	}
	password := strings.Join(buf, "")
	encryptedPassword := EncryptPassword(password)
	
	db := o.Model(&m_corp.CorpUser{}).Where("id", this.Id).Update(gorm.Params{
		"password": encryptedPassword,
	})
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return "", db.Error
	}
	
	return password, nil
}

func NewCorpUserFromModel(ctx context.Context, model *m_corp.CorpUser) *CorpUser {
	corpUser := new(CorpUser)
	corpUser.Ctx = ctx
	corpUser.Id = model.Id
	corpUser.CorpId = model.CorpId
	corpUser.IsActive = model.IsActive
	corpUser.IsManager = model.IsManager
	corpUser.Username = model.Username
	corpUser.RealName = model.RealName
	corpUser.CreatedAt = model.CreatedAt
	return corpUser
}

func NewCorpUserFromId(ctx context.Context, corpUserId int) *CorpUser{
	instance := new(CorpUser)
	instance.Ctx = ctx
	instance.Id = corpUserId
	return instance
}

func init() {
}
