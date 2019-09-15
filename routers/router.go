package routers

import (
	"github.com/gingerxman/eel"
	"github.com/gingerxman/eel/handler/rest/console"
	"github.com/gingerxman/ginger-account/rest/user"
	"github.com/gingerxman/ginger-account/rest/login"
)

func init() {
	eel.RegisterResource(&console.Console{})
	
	eel.RegisterResource(&user.NewUser{})
	eel.RegisterResource(&user.User{})
	
	eel.RegisterResource(&login.LoginedUser{})
}