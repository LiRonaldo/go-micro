package ServicesImpl

import (
	"context"
	"go-micro/Services"
	"strconv"
	"time"
)

type ProdService struct {
}

func newProd(id int32, pname string) *Services.ProdModel {
	return &Services.ProdModel{ProdID: id, ProdName: pname}
}
func (*ProdService) GetProdsList(ctx context.Context, req *Services.ProdsRequest, resp *Services.ProdListResponse) error {
	time.Sleep(time.Second * 3)
	var models []*Services.ProdModel = make([]*Services.ProdModel, 0)
	var i int32
	for i = 0; i < req.Size; i++ {
		models = append(models, newProd(100+i, "prodname"+strconv.Itoa(100+int(i))))
	}
	resp.Data = models
	return nil
}
func (*ProdService) GetProdDetail(ctx context.Context, req *Services.ProdsRequest, resp *Services.ProdDetailResponse) error {
	time.Sleep(time.Second * 3)
	resp.Data = newProd(req.ProdId, "测试数据")
	return nil
}
