package models

import (
	"github.com/jinzhu/gorm"

	"time"
)

type ArticleCRUD interface {
	ExistByID(id int) bool
	Get(id int) (t Article) //可以看TagCRUD接口，拥有同名方法也没关系
	GetArticleTotal(maps interface{}) (count int)
	GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article)
	//GetArticle(id int) (article Article)
	EditArticle(id int, data interface{}) bool
	AddArticle(data map[string]interface{}) bool
	DeleteArticle(id int) bool
}

//ArticleDao 实现ArticleCRUD接口，后续写单元测试时可以很方便的将数据库依赖替换
type ArticleDao struct {
}

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag" gorm:"PRELOAD:false"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`

	WriterId1 string `json:"writer_id_1" gorm:"column:writer_id"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

func (ArticleDao) ExistByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true
	}

	return false
}

func (ArticleDao) Get(id int) (article Article) {
	db.Debug().Where("id = ?", id).First(&article)
	db.Debug().Model(&article).Related(&article.Tag)

	return
}

func (ArticleDao) GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

func (ArticleDao) GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Debug().Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

// models包 的方法
func GetArticle(id int) (article Article) {
	db.Debug().Where("id = ?", id).First(&article)
	db.Debug().Model(&article).Related(&article.Tag)

	return
}

func (ArticleDao) EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

func (ArticleDao) AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
		WriterId1: data["writer_id"].(string),
	})

	return true
}

func (ArticleDao) DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})

	return true
}
