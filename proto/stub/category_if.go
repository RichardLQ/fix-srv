package stub

import (
	"context"
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
	return &rsp,nil
}


