package main

import (
	"com.github/menghx/mxgo"
	"os"
	"fmt"
)

func main(){
	fmt.Println(os.Getwd())
	mxgo.Run()
}

