package stub

import (
	"context"
	"fmt"
	"github.com/RichardLQ/fix-srv/logic"
)

//GetSearchList 热门分类图片内容
func (this *categoryServiceServer) GetSearchList(ctx context.Context, req *GetSearchReq) (*GetSearchRsp, error) {
	rsp := GetSearchRsp{
		Code: 200,
		Msg: "请求成功！",
	}
	data,err :=logic.GetSearchList(req.Keyword,req.Start)
	fmt.Println(data,err)
	return &rsp,nil
}






