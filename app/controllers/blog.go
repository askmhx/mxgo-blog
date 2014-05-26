package controllers

import "github.com/menghx/mxgo"

type Blog struct{
	mxgo.Controller
}

func (blog Blog)View(bid string) mxgo.Result{
	return blog.Template()
}
