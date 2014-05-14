package main

import (
	"com.github/menghx/mxgo"
	"os"
	"fmt"
	"blgo/app/filters"
)

func main(){
	fmt.Println(os.Getwd())
	mxgo.AddRoute("/index","*","App.Index")
	mxgo.AddFilter(&LoginFilter{})
	mxgo.EnableAdmin(true)
	mxgo.Run()
}

