package controllers

import "github.com/menghx/mxgo"

type Index struct {
	mxgo.Controller
}

func (app Index) Index() mxgo.Result{
	app.Data="hello"
	return  app.Template();
}

func (app Index) Test() mxgo.Result{
	app.Data="test"
	return  app.Template();
}

