package controllers

import "github.com/menghx/mxgo"

type App struct {
	*mxgo.Controller
}

func (app App) Index() Result{
	return  app.Template();
}

