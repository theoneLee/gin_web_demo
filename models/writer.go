package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Writer struct {
	Model

	//TagID int `json:"tag_id" gorm:"index"`
	ArticleList []Article `json:"article_list" gorm:"foreignkey:WriterId1"` //默认外键字段为Model名+Id

	Name string
}

//func ExistWriter(id int){
//	db.Where("id=?",id).
//}

func GetWriter(id int) (*Writer, error) {
	var w Writer
	//err := db.Debug().Where("id=?",id).First(&w).Error
	//if err != nil && err != gorm.ErrRecordNotFound {
	//	return nil, err
	//}
	//err = db.Model(&w).Related(&w.ArticleList).Error //todo 这样写直接返回&w会使得一对多的关联字段value是null，而不是像查询多个时返回的[],err为关联关系为空
	//err = db.Debug().Model(&w).Preload("Article").Find(&w.ArticleList).Error
	//err := db.Debug().Table("blog_writer").Where("blog_writer.id=?",id).Joins("left join blog_article a on a.user_id = writer.id").Scan(&w).Error

	err := db.Debug().Preload("ArticleList").Where("id=?", id).First(&w).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("get writer err:", err)
		w.ArticleList = make([]Article, 0)
		return &w, nil
	}
	return &w, nil
}

func AddWriter(data map[string]interface{}) bool {
	w := Writer{
		//ArticleList: data[],
		Name: data["name"].(string),
	}
	db.Create(&w)
	return true
}

func GetWriters(pageNum int, pageSize int) (wl []Writer) {
	db.Debug().Preload("ArticleList").Offset(pageNum).Limit(pageSize).Find(&wl)
	return
}

//关联由创建article时进行关联，
