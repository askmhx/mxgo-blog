package main

import (
	"github.com/menghx/mxgo"
	"os"
	"fmt"
	"blgo/app/filter"
)

func main(){
	fmt.Println(os.Getwd())
	mxgo.Router("/index","*","App.Index")
	mxgo.AddFilter(&filter.LoginFilter{},&filter.LoginFilter{})
	mxgo.EnableAdmin(true)
	mxgo.Run()
}

