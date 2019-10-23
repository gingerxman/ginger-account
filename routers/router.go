package routers

import (
	"github.com/gingerxman/eel"
	"github.com/gingerxman/eel/handler/rest/console"
	"github.com/gingerxman/eel/handler/rest/op"
	"github.com/gingerxman/ginger-account/rest/dev"
	"github.com/gingerxman/ginger-account/rest/login"
	"github.com/gingerxman/ginger-account/rest/user"
	"github.com/gingerxman/ginger-account/rest/corp"
)

func init() {
	eel.RegisterResource(&console.Console{})
	eel.RegisterResource(&op.Health{})
	
	eel.RegisterResource(&user.NewUser{})
	eel.RegisterResource(&user.User{})
	
	eel.RegisterResource(&corp.Corp{})
	
	eel.RegisterResource(&login.LoginedUser{})
	eel.RegisterResource(&login.LoginedCorpUser{})
	eel.RegisterResource(&login.LoginedMallUser{})
	eel.RegisterResource(&login.LoginedBDDUser{})
	
	eel.RegisterResource(&dev.BDDReset{})
	eel.RegisterResource(&dev.AllCorps{})
}