package user

import (
	mock_person "gin-blog/pkg/mock"
	"github.com/golang/mock/gomock"
	"testing"
)

/*
https://book.eddycjy.com/golang/talk/gomock.html
 要测试user的接收器方法GetUserInfo。

$ mockgen -source=./person/male.go -destination=./mock/male_mock.go -package=mock
-source：设置需要模拟（mock）的接口文件
-destination：设置 mock 文件输出的地方，若不设置则打印到标准输出中
-package：设置 mock 文件的包名，若不设置则为 mock_ 前缀加上文件名（如本文的包名会为 mock_person）
*/

func TestUser_GetUserInfo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	var id int64 = 1
	mockMale := mock_person.NewMockMale(ctl) //创建一个新的 mock 实例
	gomock.InOrder(                          //声明给定的调用应按顺序进行（是对 gomock.After 的二次封装）
		mockMale.EXPECT().Get(id).Return(nil), //EXPECT()返回一个允许调用者设置期望和返回值的对象。Get(id) 是设置入参并调用 mock 实例中的方法。Return(nil) 是设置先前调用的方法出参。简单来说，就是设置入参并调用，最后设置返回值
	)

	user := NewUser(mockMale) //将mock的Male对象作为NewUser构造函数的参数，测试所需要的函数查看返回情况
	err := user.GetUserInfo(id)
	if err != nil {
		t.Errorf("user.GetUserInfo err: %v", err)
	}
}
