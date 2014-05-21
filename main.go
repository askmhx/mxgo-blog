package main

import (
	"github.com/menghx/mxgo"
	"blgo/app/filter"
)

func main(){
	mxgo.Router("/index","*","App.Index")
	mxgo.AddFilter(&filter.LoginFilter{},&filter.LoginFilter{})
	mxgo.EnableAdmin(true)
	mxgo.Run()
}

