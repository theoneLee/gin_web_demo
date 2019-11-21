package corn

import (
	"fmt"
	"github.com/robfig/cron"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

/*
	开辟一个定时任务定期登陆账户获取cookie
	开辟一个定时任务，调用接口把数据爬出来。
*/

func test() {
	log.Println("Starting...")

	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		//models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		//models.CleanAllArticle()
	})

	c.Start()

	//阻塞主程序
	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
	//select{}
}

// cms 模拟登陆 获取cookie内容 可以通过postman生成下面代码
func login() {
	url := "http://127.0.0.1:8083/cms/loginController.do?checkuser=" //"http://127.0.0.1:8083/cms/loginController.do?checkuser&userKey=D1B5CC2FE46C4CC983C073BCA897935608D926CD32992B5900&userName=admin&password=hnxd123456"

	payload := strings.NewReader("userKey=D1B5CC2FE46C4CC983C073BCA897935608D926CD32992B5900&userName=admin&password=hnxd123456")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Origin", "http://192.168.12.53:8083")
	req.Header.Add("Referer", "http://192.168.12.53:8083/cms/loginController.do?login")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "c551c0ac-aedb-43b9-a5b3-5d832f5bb13b")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println("==================")
	fmt.Println("header:", res.Header)
	fmt.Println("==================")

	fmt.Println(string(body))
	// todo 拿到cookie，入库
}

func getData() {
	// todo 获取最新的cookie，
	//  构造和请求这个接口，获取成功就换页，放入自己数据库，然后记录当前分页
	//  http://127.0.0.1:8083/cms/newstopController.do?datagrid&field=id,docNo,dataid,ctitle,site,sourceurl,channel,issuedate,maincontent

}
