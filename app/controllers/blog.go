package controllers

import "github.com/menghx/mxgo"

type Blog struct{
	mxgo.Controller
}

func (blog Blog)view(bid string) mxgo.Result{
	return blog.Template()
}
