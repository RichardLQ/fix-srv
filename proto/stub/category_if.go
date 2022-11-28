package stub

import (
	"context"
	"github.com/RichardLQ/fix-srv/logic"
	"github.com/RichardLQ/fix-srv/refer"
	"github.com/tidwall/gjson"
)

var UserService = &categoryServiceServer{}

type categoryServiceServer struct {
	FixServiceServer
}

//GetPicCategory 获取分类列表
func (this *categoryServiceServer) GetPicCategory(ctx context.Context, req *GetPicCategoryReq) (*GetPicCategoryRsp, error) {
	rsp := GetPicCategoryRsp{
		Code: 200,
		Msg: "请求成功！",
	}
	data,err :=logic.GetPicCategory()
	if err != nil {
		rsp.Code = refer.Find_Piclist_Err
		rsp.Msg = "请求失败！"
	}
	for i, datum := range data {
		temp := &Photo{
			Id: int32(i+1),
			Pid: gjson.Get(datum.String(),"id").String(),
			Ename: gjson.Get(datum.String(),"ename").String(),
			Address: gjson.Get(datum.String(),"cover").String(),
			Zname: gjson.Get(datum.String(),"name").String(),
		}
		rsp.List = append(rsp.List,temp)
	}
	return &rsp,nil
}


