package article

import (
	"fmt"
	"gin-blog/models"
	"gin-blog/pkg/e"
	"testing"
)

var mockdao = mockArticleDao{map[int]*models.Article{
	1: new(models.Article),
	2: new(models.Article),
	3: new(models.Article),
}}

//参考onenote笔记 "Go 方法，接口"
// https://onedrive.live.com/view.aspx?resid=54FDED4F8B6CD09C%218378&id=documents&wd=target%28%E4%BA%91%E8%97%8F%E6%95%B0%E6%8D%AE%2Fgo.one%7C3F7DD993-0459-A141-A68C-D6533CDC109C%2FGo%20%E6%96%B9%E6%B3%95%EF%BC%8C%E6%8E%A5%E5%8F%A3%7C6290CD4F-C9BC-ED41-B1C2-C8FE58A9FD9F%2F%29 onenote:https://d.docs.live.net/54fded4f8b6cd09c/文档/新的历程（为以后身边的人不痛苦而努力，姑丈，我不会忘记。）/云藏数据/go.one#Go%20方法，接口&section-id={3F7DD993-0459-A141-A68C-D6533CDC109C}&page-id={6290CD4F-C9BC-ED41-B1C2-C8FE58A9FD9F}&end

type mockArticleDao struct {
	m map[int]*models.Article
}

func (dao *mockArticleDao) ExistByID(id int) bool {
	_, ok := dao.m[id]

	if ok {
		return true
	}

	return false
}

func (dao *mockArticleDao) Get(id int) (article models.Article) {
	articlePtr, ok := dao.m[id]

	if ok {
		article = *articlePtr
		return
	}

	return
}

func (dao *mockArticleDao) GetArticleTotal(maps interface{}) (count int) {
	//db.Model(&Article{}).Where(maps).Count(&count)
	count = len(dao.m)
	return
}

func (dao *mockArticleDao) GetArticles(pageNum int, pageSize int, maps interface{}) (articles []models.Article) {
	//db.Debug().Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	// todo map convert to list
	return
}

func (dao *mockArticleDao) EditArticle(id int, data interface{}) bool {
	//db.Model(&Article{}).Where("id = ?", id).Updates(data)
	//update from map
	return true
}

func (dao *mockArticleDao) AddArticle(data map[string]interface{}) bool {
	//db.Create(&Article{
	//	TagID:     data["tag_id"].(int),
	//	Title:     data["title"].(string),
	//	Desc:      data["desc"].(string),
	//	Content:   data["content"].(string),
	//	CreatedBy: data["created_by"].(string),
	//	State:     data["state"].(int),
	//	WriterId1: data["writer_id"].(string),
	//})
	//todo add into map
	return true
}

func (dao *mockArticleDao) DeleteArticle(id int) bool {
	//db.Where("id = ?", id).Delete(Article{})
	//todo delete from map
	return true
}

// 组件依赖mock
func TestService_GetArticle(t *testing.T) {
	srv := NewService(&mockdao)
	id := 1 //1~3 will pass //todo 可做基于表的测试。整理web controller，service，dao 的最佳实践（上面的mockdao可以利用工具生成）
	data, code := srv.GetArticle(id)
	if code == e.ERROR_NOT_EXIST_ARTICLE {
		t.Fail()
	}
	fmt.Printf("data:%+v |code:%v\n", data, code)
}

//==========函数依赖============

//待测试函数如果这样写，会导致该函数依赖无法由外部赋值
func Reply(username, password string) bool {
	if Login(username, password) {
		fmt.Println("login success")
		return true
	}
	return false
}

//应当这样写，才可以处理函数依赖
var LoginStub = Login

func Reply2(username, password string) bool {
	if LoginStub(username, password) {
		fmt.Println("login success")
		return true
	}
	return false
}

func Login(username, password string) bool {
	if username == password {
		return true
	} else {
		return false
	}
}

// 下面是Reply2的测试代码，可以在外部注入Login行为
func TestSuccessReply(t *testing.T) {
	ori := LoginStub
	defer func() { LoginStub = ori }() //恢复LoginStub为未mock前的函数

	//mock函数
	LoginStub = func(username, password string) bool {
		return true
	}

	if !Reply2("a", "b") {
		t.Errorf("登陆成功，却回复失败") //正常情况是登陆成功，应该让他成功回复
	}

}
