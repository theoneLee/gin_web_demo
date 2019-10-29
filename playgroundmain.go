package main

import pg "gin-blog/playerground"

func main() {
	var phone pg.Phone
	phone = new(pg.NokiaPhone) //实现Phone接口
	phone.Call()               // 包外使用需要Export 接口的方法

	phone2 := pg.NokiaPhone{} //实现Phone接口
	phone2.Call()

	pg.NokiaPhone{}.Call()

	//phone = new(IPhone) //IPhone结构体没有实现listen方法，所以不属于Phone接口
	//phone.call()

}
