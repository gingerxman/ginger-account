package blog

import (
	"github.com/gingerxman/eel"
)

type Corp struct {
	eel.Model
	Code string `gorm:"size:128"`
	Name string `gorm:"size:128"`
	Logo string `gorm:"size:1024"`
	Remark string `gorm:"size:128"`
	IsActive bool `orm:"default:true"`
	IsPlatform bool `orm:"default:false"`
}
func (self *Corp) TableName() string {
	return "corp_corp"
}

type CorpUser struct {
	eel.Model
	CorpId int
	Username string `gorm:"size:128"`
	RealName string `gorm:"size:128"`
	IsActive bool `gorm:"default:true"`
	Password string `gorm:"size:256"`
}
func (self *CorpUser) TableName() string {
	return "corp_user"
}

type Permission struct {
	eel.Model
	Code string `gorm:"size:80"`
}
func (this *Permission) TableName() string {
	return "auth_permission"
}

type UserHasPermission struct {
	eel.Model
	CorpUserId int
	PermissionId int
}
func (this *UserHasPermission) TableName() string {
	return "auth_user_has_permission"
}

type Group struct {
	eel.Model
	Name string `gorm:"size:80;unique"`
}
func (this *Group) TableName() string {
	return "auth_group"
}

type UserHasGroup struct {
	eel.Model
	CorpUserId int
	GroupId int
}
func (this *UserHasGroup) TableName() string {
	return "auth_user_has_group"
}

type GroupHasPermission struct {
	eel.Model
	PermissionId int
	GroupId int
}
func (this *GroupHasPermission) TableName() string {
	return "auth_group_has_permission"
}

func init() {
	eel.RegisterModel(new(Corp))
	eel.RegisterModel(new(CorpUser))
	eel.RegisterModel(new(Permission))
	eel.RegisterModel(new(UserHasPermission))
	eel.RegisterModel(new(Group))
	eel.RegisterModel(new(GroupHasPermission))
	eel.RegisterModel(new(UserHasPermission))
}
