package corp

import (
	"context"
	"github.com/gingerxman/ginger-account/business/constant"
	
	"time"
	
	"github.com/gingerxman/eel"
	m_corp "github.com/gingerxman/ginger-account/models/corp"
)


type CorpUser struct {
	eel.EntityBase
	Id int
	CorpId int
	Username string
	RealName string
	CreatedAt time.Time
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

func NewCorpUserFromModel(ctx context.Context, model *m_corp.CorpUser) *CorpUser {
	corpUser := new(CorpUser)
	corpUser.Ctx = ctx
	corpUser.Id = model.Id
	corpUser.CorpId = model.CorpId
	corpUser.Username = model.Username
	corpUser.RealName = model.RealName
	corpUser.CreatedAt = model.CreatedAt
	return corpUser
}

func init() {
}
