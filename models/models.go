package models

import (
	"fmt"
	"gin-blog/pkg/setting"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

// 该函数会在模块被加载时默认执行，但最好还是自己去做init工作，换一个名字，然后手动调用
//func init() {
//	var (
//		err                                               error
//		dbType, dbName, user, password, host, tablePrefix string
//	)
//
//	sec, err := setting.Cfg.GetSection("database")
//	if err != nil {
//		log.Fatal(2, "Fail to get section 'database': %v", err)
//	}
//
//	dbType = sec.Key("TYPE").String()
//	dbName = sec.Key("NAME").String()
//	user = sec.Key("USER").String()
//	password = sec.Key("PASSWORD").String()
//	host = sec.Key("HOST").String()
//	tablePrefix = sec.Key("TABLE_PREFIX").String()
//
//	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
//		user,
//		password,
//		host,
//		dbName))
//
//	if err != nil {
//		log.Println(err)
//	}
//
//	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
//		return tablePrefix + defaultTableName
//	}
//
//	db.SingularTable(true)
//	db.DB().SetMaxIdleConns(10)
//	db.DB().SetMaxOpenConns(100)
//}
func InitMysql() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}

func EditTag(id int, data map[string]interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)

	return true
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
