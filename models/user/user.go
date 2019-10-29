package blog

import (
	"github.com/gingerxman/eel"
)


const USER_SEX_FEMALE = 0
const USER_SEX_MALE = 1
const USER_SEX_UNKNOWN = 2
var STR2SEX = map[string]int{
	"female": USER_SEX_FEMALE,
	"male": USER_SEX_MALE,
	"unknown": USER_SEX_UNKNOWN,
	"": USER_SEX_UNKNOWN,
}
var SEX2STR = map[int]string {
	USER_SEX_FEMALE: "female",
	USER_SEX_MALE: "male",
	USER_SEX_UNKNOWN: "unknown",
}
// User
type User struct {
	eel.Model
	Unionid string `gorm:"size:255"`
	IsActive bool `gorm:"default:true"`
	CorpUserId int `gorm:"default:0"`
	Code string `gorm:"size:32"`
	
	//基本信息
	Name string `gorm:"size:52"`
	Avatar string `gorm:"size:1024"`
	Password string `gorm:"size:52"`
	Sex int
	
	//其他信息
	Source string `gorm:"size:52"`
}
func (this *User) TableName() string {
	return "user_user"
}

func init() {
	eel.RegisterModel(new(User))
}
