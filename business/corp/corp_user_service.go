package corp

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/eel/config"
	m_corp "github.com/gingerxman/ginger-account/models/corp"
	"io"
)

var SUPER_PASSWORD = ""

type CorpUserService struct {
	eel.ServiceBase
}

func NewCorpUserService(ctx context.Context) *CorpUserService {
	service := new(CorpUserService)
	service.Ctx = ctx
	return service
}

func (this *CorpUserService) encryptPassword(password string) string {
	hash := sha1.New()
	io.WriteString(hash, password)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (this *CorpUserService) Auth(username string, password string) (*CorpUser, error) {
	o := eel.GetOrmFromContext(this.Ctx)
	
	model := m_corp.CorpUser{}
	db := o.Model(&model).Where("username", username).Take(&model)
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return nil, db.Error
	}
	
	return NewCorpUserFromModel(this.Ctx, &model), nil
}

func init() {
	SUPER_PASSWORD = config.ServiceConfig.String("system::SUPER_PASSWORD")
}
