package main

import (
	"SoftwareDesignModelByGo/Facade"
)

func main(){
	//fmt.Println(Factory.NewAPI(1).Say("HELLO"))
	//fmt.Println(Factory.NewAPI(2).Say("HELLO"))
	Facade.NewAPI().Test()
}