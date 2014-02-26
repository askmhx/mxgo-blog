package controllers

import "com.github/menghx/mxgo"

type App struct {
	*mxgo.Controller
}

func (app App) Index() Result{
	return  &TemplateResult{};
}

