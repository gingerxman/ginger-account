package auth

type RPermission struct {
	Id int `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type RGroup struct {
	Id     int    `json:"id"`
	Code string `json:"code"`
	Name   string `json:"name"`
	Permissions []*RPermission `json:"permissions"`
}
