package egb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func httpGet() {
	str, err := Get("http://www.baidu/com").Param("name", "anlina").Exec().ToString()
	fmt.Println(str, err)
}

func httpPost() {
	postparams := map[string]interface{}{
		"name": "angelina",
		"age":  "21",
	}
	json, _ := json.Marshal(&postparams)
	str, err := Post("expamleurl.com/post").Json(string(json)).Exec().ToString()
	fmt.Println(str, err)
}

func setHeader() {
	str, err := Get("http://www.baidu/com").
		//设置url参数
		Param("name", "anlina").
		//设置header
		Set("header1", "value").
		//设置请求的client
		Use(&http.Client{}).
		//执行httpl请求
		Exec().
		//结果输出为string
		ToString()
	fmt.Println(str, err)
}
