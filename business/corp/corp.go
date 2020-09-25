package corp

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/gingerxman/ginger-account/business/auth"
	"github.com/gingerxman/gorm"
	"io"
	
	"time"
	
	"github.com/gingerxman/eel"
	m_corp "github.com/gingerxman/ginger-account/models/corp"
)

func EncryptPassword(password string) string {
	hash := sha1.New()
	io.WriteString(hash, password)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

type Corp struct {
	eel.EntityBase
	Id int
	Code string
	Name string
	Logo string
	Remark string
	IsActive bool
	IsPlatform bool
	CreatedAt time.Time
	
	CorpUser *CorpUser
}

func (this *Corp) GetId() int {
	return this.Id
}

//AddCorpUser 为corp创建登录账号
func (this *Corp) AddCorpUser(username string, password string, realname string, groupNames []string, isManager bool) error {
	model := &m_corp.CorpUser{
		Username: username,
		RealName: realname,
		Password: EncryptPassword(password),
		CorpId: this.Id,
		IsManager: isManager,
	}
	o := eel.GetOrmFromContext(this.Ctx)
	db := o.Create(model)
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		panic(eel.NewBusinessError("corp_factory:create_corp_fail", "创建Corp失败"))
	}
	
	if len(groupNames) > 0 {
		groups := auth.NewGroupRepository(this.Ctx).GetGroupsByNames(groupNames)
		for _, group := range groups {
			model := &m_corp.UserHasGroup{
				CorpUserId: model.Id,
				GroupId: group.Id,
			}
			db := o.Create(model)
			if db.Error != nil {
				eel.Logger.Error(db.Error)
				panic(eel.NewBusinessError("corp_factory:create_corp_fail", "添加用户角色失败"))
			}
		}
	}
	
	return nil
}

func (this *Corp) GetUnionid() string {
	return fmt.Sprintf("corp_%d", this.Id)
}

func (this *Corp) Fill() {
	if this.Code != "" {
		return
	}
	
	model := m_corp.Corp{}
	db := eel.GetOrmFromContext(this.Ctx).Model(&m_corp.Corp{}).Where("id", this.Id).Take(&model)
	
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		panic(eel.NewBusinessError("corp:invalid_corp", fmt.Sprintf("无效corp(%d)", this.Id)))
	}
	
	this.FillWithModel(&model)
}

func (this *Corp) FillWithModel(model *m_corp.Corp) {
	if this.Code != "" {
		return
	}
	
	corp := this
	corp.Model = model
	corp.Id = model.Id
	corp.Code = model.Code
	corp.Name = model.Name
	corp.Logo = model.Logo
	corp.Remark = model.Remark
	corp.IsActive = model.IsActive
	corp.IsPlatform = model.IsPlatform
	corp.CreatedAt = model.CreatedAt
}

func (this *Corp) Delete() error {
	o := eel.GetOrmFromContext(this.Ctx)
	
	db := o.Model(&m_corp.Corp{}).Where("id", this.Id).Update(gorm.Params{
		"is_active": false,
	})
	
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return db.Error
	}
	
	return nil
}

func NewCorpFromId(ctx context.Context, id int) *Corp {
	corp := new(Corp)
	corp.Ctx = ctx
	corp.Model = nil
	corp.Id = id
	return corp
}

func GetCorpFromContext(ctx context.Context) *Corp {
	corp := ctx.Value("corp").(*Corp)
	return corp
}

func NewCorpFromModel(ctx context.Context, model *m_corp.Corp) *Corp {
	corp := new(Corp)
	corp.Ctx = ctx
	corp.FillWithModel(model)
	return corp
}

func init() {
}
