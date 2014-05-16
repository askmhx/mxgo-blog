package main

import (
	"github.com/menghx/mxgo"
	"os"
	"fmt"
	"blgo/app/filters"
)

func main(){
	fmt.Println(os.Getwd())
	mxgo.Router("/index","*","App.Index")
	mxgo.AddFilter(&filters.LoginFilter{},&filters.LoginFilter{})
	mxgo.EnableAdmin(true)
	mxgo.Run()
}

