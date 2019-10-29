package v1

import "C"
import (
	"gin-blog/models"
	"gin-blog/pkg/app"
	"gin-blog/pkg/e"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func AddWriter(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	//wid := com.StrTo(c.Query("id")).MustInt()

	valid := validation.Validation{}
	valid.Required(name, "name").Message("name should be not null")

	code := e.INVALID_PARAMS
	if valid.HasErrors() {
		appG.Response(http.StatusInternalServerError, code, nil)
		return
	}

	data := make(map[string]interface{})
	data["name"] = name
	models.AddWriter(data)

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

func GetWriter(c *gin.Context) {
	appG := app.Gin{C: c}

	//wid := com.StrTo(c.Query("id")).MustInt()
	wid := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(wid, 1, "id").Message("id should be more than 0")

	if valid.HasErrors() {
		appG.Response(http.StatusInternalServerError, e.INVALID_PARAMS, nil)
		return
	}

	data, err := models.GetWriter(wid)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, data)

}

func GetWriters(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})
	//maps := make(map[string]interface{})
	var wl []models.Writer
	wl = models.GetWriters(util.GetPage(c), setting.PageSize)
	data["lists"] = wl
	data["total"] = len(wl) //models.GetArticleTotal(maps)

	appG.Response(http.StatusOK, e.SUCCESS, data)

}
