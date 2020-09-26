package user

type RUser struct {
	Id     int    `json:"id"`
	Unionid string `json:"unionid"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Sex    string `json:"sex"`
	Code   string `json:"code"`
	Source string `json:"source"`
}
