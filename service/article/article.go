package article

import (
	"gin-blog/models"
	"gin-blog/pkg/e"
)

/*
service
*/

type Service struct {
	Dao *models.ArticleDao
}

func NewService(d models.ArticleDao) *Service {
	srv := Service{&d}
	return &srv
}

func (service Service) GetArticle(id int) (data interface{}, code int) {
	dao := *service.Dao
	if dao.ExistByID(id) {
		data = dao.Get(id)
		code = e.SUCCESS
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}
	return
}

// todo 在v1，article的除了参数获取，参数校验，数据打包和返回外的逻辑抽出来，作为Service接收器方法
