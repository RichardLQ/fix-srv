package logic

import (
	"github.com/RichardLQ/fix-srv/auth"
	"github.com/RichardLQ/fix-srv/refer"
	"github.com/tidwall/gjson"
)



//GetPicCategory 获取图片分类
func GetPicCategory() ([]gjson.Result,error) {
	url := refer.API_HOST + "v1/vertical/category"
	listString, err := auth.SendHttpRequest(url, "", "", nil, nil)
	if err != nil {
		return []gjson.Result{},err
	}
	data := gjson.Get(listString, "res.category").Array()

	return data,err
}
