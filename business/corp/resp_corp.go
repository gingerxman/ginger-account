package corp

type RCorpUser struct {
	Id int `json:"id"`
	Name string `json:"name"`
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
