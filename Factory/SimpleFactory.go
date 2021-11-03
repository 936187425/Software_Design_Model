
/**
 简单工厂模式：
	即外部只能调用接口,内部实现是封装的。在Go语言中没有真正的构造函数，
	一般是用NewXXX函数返回接口实现，封装接口。
 */
package Factory



//本文件中只有1 API接口  2定义的方法  3NewXXX函数是对外可见,其他都是不可见的
type API interface {
	Say(name string) string
}

//注意要在NewXXX中添加输入参数,来决定合适的类创建
func NewAPI(condition int) (res API) {
	if condition == 1 {
		//注意是hiAPI类型实现了Say方法
		res = *new(hiAPI)
	} else if condition == 2 {
		//注意是*hellpAPI实现了Say方法
		res = new(helloAPI)
	}
	return res
}
/**
以下是实现类
 */

type hiAPI struct{}

type helloAPI struct{}

func (self hiAPI)Say(name string) string{//注意是hiAPI类型实现了Say方法
	return "hiAPI"
}

func (self *helloAPI)Say(name string) string{//注意是*hellpAPI实现了Say方法
	return "helloAPI"
}
