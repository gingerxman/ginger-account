package routers

import (
	"github.com/gingerxman/eel"
	"github.com/gingerxman/eel/handler/rest/console"
	"github.com/gingerxman/eel/handler/rest/op"
	"github.com/gingerxman/ginger-account/rest/auth"
	"github.com/gingerxman/ginger-account/rest/dev"
	"github.com/gingerxman/ginger-account/rest/login"
	"github.com/gingerxman/ginger-account/rest/user"
	"github.com/gingerxman/ginger-account/rest/corp"
	"github.com/gingerxman/ginger-account/rest/area"
)

func init() {
	eel.RegisterResource(&console.Console{})
	eel.RegisterResource(&op.Health{})
	
	// auth
	eel.RegisterResource(&auth.Permission{})
	eel.RegisterResource(&auth.Groups{})
	eel.RegisterResource(&auth.Group{})
	eel.RegisterResource(&auth.ResetPassword{})
	
	// user
	eel.RegisterResource(&user.NewUser{})
	eel.RegisterResource(&user.User{})
	eel.RegisterResource(&user.Users{})
	
	// corp
	eel.RegisterResource(&corp.Corp{})
	eel.RegisterResource(&corp.RelatedUserId{})
	eel.RegisterResource(&corp.CorpUser{})
	eel.RegisterResource(&corp.CorpUsers{})
	
	// login
	eel.RegisterResource(&login.LoginedUser{})
	eel.RegisterResource(&login.LoginedCorpUser{})
	eel.RegisterResource(&login.LoginedMallUser{})
	eel.RegisterResource(&login.MallVisitor{})
	eel.RegisterResource(&login.LoginedBDDUser{})
	
	/*
	 area
	*/
	eel.RegisterResource(&area.Area{})
	eel.RegisterResource(&area.AreaCode{})
	eel.RegisterResource(&area.YouzanAreaList{})
	
	// dev
	eel.RegisterResource(&dev.BDDReset{})
	eel.RegisterResource(&dev.AllCorps{})
	eel.RegisterResource(&dev.Users{})
}