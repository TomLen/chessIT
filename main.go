package main

import (
	"fmt"
	"./bot"
	"./config"
)

func main(){
	err := config.ReadConfig()

	if err != nil{
		fmt.Println(err.Errror())
		return
	}

	not.Start()

	<-make(chan struct{})
	return
}