package corp

type RCorp struct {
	Id     int    `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Logo string `json:"logo"`
	Remark string `json:"remark"`
	IsPlatform bool `json:"is_platform"`
	CreatedAt string `json:"created_at"`
}
