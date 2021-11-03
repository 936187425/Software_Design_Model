/**
	完成适配器模式的基本框架
	在Adapter中匿名组合Adaptee接口，所以Adapter类也拥有SpecificRequest实例方法，
    又因为Go语言中非入侵式接口特征，其实Adapter也适配Adaptee接口。
 */
package Adapter

import (
	"fmt"
)

//1客户端想要调用的接口,此处适配器可以当作是Target接口的一种实现类
type Target interface {
	Request()
}

func NewTarget()Target{
	return &adapter{NewAdaptee()}
}

//3适配器(不对外可见)
type adapter struct{
	Adaptee
}
//实现Request 从而实现Target接口,实现方式是调用被适配的接口
func (self *adapter)Request(){
	self.Adaptee.SpecificRequest()
}



//2需要被适配的接口
type Adaptee interface {
	SpecificRequest()
}
func NewAdaptee() Adaptee{
	return new(adapteeImpl)
}
type adapteeImpl struct{}
func (self *adapteeImpl) SpecificRequest() {
	fmt.Println("I am adapteeImpl")
}

