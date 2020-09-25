package corp

import "github.com/gingerxman/ginger-account/business/auth"

type RCorpUser struct {
	Id int `json:"id"`
	CorpId int `json:"corp_id"`
	CorpName string `json:"corp_name"`
	Name string `json:"name"`
	RealName string `json:"real_name"`
	IsActive bool `json:"is_active"`
	IsManager bool `json:"is_manager"`
	Groups []*auth.RGroup `json:"groups"`
	Permissions []*auth.RPermission `json:"permissions"`
}

type RCorp struct {
	Id     int    `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Logo string `json:"logo"`
	Remark string `json:"remark"`
	IsPlatform bool `json:"is_platform"`
	CorpUser *RCorpUser `json:"corp_user"`
	CreatedAt string `json:"created_at"`
}
