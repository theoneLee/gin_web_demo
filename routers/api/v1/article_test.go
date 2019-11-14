package v1

import (
	"fmt"
	"gin-blog/models"
	"gin-blog/pkg/e"
	"github.com/astaxie/beego/validation"
	"testing"
)

var dao = mockArticleDao{map[int]*models.Article{
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

// 有个bug。使得该测试无法运行。使用go test ./... 无法通过，除去多个main的问题，还有：
//Fail to parse 'conf/app.ini': open conf/app.ini: no such file or directory
//检查test的working directory 为项目根目录，而不是v1包
func TestGetArticle(t *testing.T) {
	id := 1

	valid := validation.Validation{}
	//valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if dao.ExistByID(id) {
			data = dao.Get(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			fmt.Println(err) //log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	fmt.Printf("code:%v   data:%+v \n", code, data)

	//c.JSON(http.StatusOK, gin.H{
	//	"code": code,
	//	"msg":  e.GetMsg(code),
	//	"data": data,
	//})

}
