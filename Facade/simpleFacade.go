package Facade

import "fmt"

/**
	首先要客户端与子系统中的一个Facade进行交互,
	然后Facade与子系统中的其他模块进行交互。
	实现了API接口作为客户端与外界的facade，然后子系统中的其他模块采用工厂模式
 */

//作为facade接口,提供一组接口的界面(可以采用简单工厂模式)
type API interface {
	//子系统中要实现的方法
	Test()
}
func NewAPI() API{
	//return new(apiImpl) 对apiImpl结构体中的两个变量并没有实例化
	return &apiImpl{NewAModuleAPI(),NewBModuleAPI()}
}

type apiImpl struct{ //接口实现
	a AModuleAPI
	b BModuleAPI
}
func (self *apiImpl)Test() {
	self.a.Printf()
	self.b.Printf()
}

//API facade所在系统的子系统中子模块AModuleAPI 采用简单工厂模式
type AModuleAPI interface{
	Printf() //此处
}

func NewAModuleAPI() AModuleAPI{
	return new(amoduleImpl)
}

type amoduleImpl struct{}

func (self *amoduleImpl)Printf(){
	fmt.Println("I am amoduleImpl")
}

//API facade所在系统的子系统中子模块BModuleAPI 采用简单工厂模式
type BModuleAPI interface {
	Printf()
}

func NewBModuleAPI() BModuleAPI{
	return new(bmoduleImpl)
}

type bmoduleImpl struct{}

func (self *bmoduleImpl) Printf(){
	fmt.Println("I am bmoduleImpl")
}