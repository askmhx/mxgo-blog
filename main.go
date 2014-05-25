package main

import (
	"github.com/menghx/mxgo"
	"blog/app/filters"
	"blog/app/controllers"
)

func main() {
	mxgo.Router("GET:/test",&controllers.Index{},"Test")
	mxgo.Router("*:/index",&controllers.Index{},"Index")
	mxgo.Router("*:/blog/(P?)",&controllers.Blog{},"View")
	mxgo.AddFilter(&filters.LoginFilter{}, &filters.LoginFilter{})
	mxgo.EnableAdmin(true)
	mxgo.Run()
}
