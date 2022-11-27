package logic

import (
	"fmt"
	"github.com/RichardLQ/fix-srv/auth"
	"github.com/RichardLQ/fix-srv/proto/stub"
	"github.com/RichardLQ/fix-srv/refer"
	"github.com/tidwall/gjson"
)



//GetPicCategory 获取图片分类
func GetPicCategory(req []*stub.Photo) error {
	url:= refer.API_HOST + "v1/vertical/category"
	listString,err:=auth.SendHttpRequest(url,"","",nil,nil)
	if err != nil {
		return err
	}
	data := gjson.Get(listString,"res.category").String()
	//list := []*stub.Photo{}
	gjson.ForEachLine(data, func(line gjson.Result) bool {
		fmt.Println("name:", gjson.Get(line.String(), "name"))
		return true
	})
	//for i, datum := range data {
	//	//temp := &stub.Photo{
	//	//	Ename: gjson.Get(datum.String(),"ename"),
	//	//}
	//	fmt.Println(i)
	//	fmt.Println(datum)
	//}
	return err
}