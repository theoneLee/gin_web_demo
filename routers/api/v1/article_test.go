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

type mockArticleDao struct {
	m map[int]*models.Article
}

func (dao *mockArticleDao) ExistByID(id int) bool {
	article, ok := dao.m[id]

	if ok && article.ID > 0 {
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

	fmt.Printf("code:%v   data:%+v", code, data)

	//c.JSON(http.StatusOK, gin.H{
	//	"code": code,
	//	"msg":  e.GetMsg(code),
	//	"data": data,
	//})

}
