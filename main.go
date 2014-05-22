package main

import (
	"blog/app/filter"
	"github.com/menghx/mxgo"
)

func main() {
	mxgo.Router("/index", "*", "App.Index")
	mxgo.AddFilter(&filter.LoginFilter{}, &filter.LoginFilter{})
	mxgo.EnableAdmin(true)
	mxgo.Run()
}
