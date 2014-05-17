package controller

import "github.com/menghx/mxgo"

type App struct {
	*mxgo.Controller
}

func (app App) Index() Result{
	app.Data="hello"
	return  app.Template();
}

