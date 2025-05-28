package controller

import (
	"context"
	models "product_srv/model"
	pb "product_srv/proto/product"
	"time"
	"user_srv/data_source"
)

type ProductHandler struct{}

func (*ProductHandler) ProductsList(ctx context.Context, req *pb.ProductsRequest, res *pb.ProductsResponse) error {
	currentPage := req.CurrentPage
	pagesize := req.Pagesize
	products := []models.ProductModel{}

	offsetNum := pagesize * (currentPage - 1)
	result := data_source.Db.Limit(int(pagesize)).Offset(int(offsetNum)).Find(&products)
	if result.Error != nil {
		res.Code = 500
		res.Msg = "获取数据失败"
		return result.Error
	}
	res.Code = 200
	res.Msg = "获取数据成功"
	var count int64
	products_count := []models.ProductModel{}
	data_source.Db.Find(&products_count).Count(&count) //用于获取总条数

	products_response := []*pb.Product{} // 定义一个空的切片用于存取待返回的产品数据
	for _, product := range products {
		product_tmp := &pb.Product{
			Id:         product.Id,
			Name:       product.Name,
			Price:      product.Price,
			Num:        product.Num,
			Unit:       product.Unit, // 单位
			Desc:       product.Desc,
			Pic:        product.Pic,
			CreateTime: product.CreateTime.Format("2006-01-02 15:04:05"),
		}
		products_response = append(products_response, product_tmp)
	}
	res.Products = products_response
	res.Total = count         // 总条数
	res.Current = currentPage // 当前页
	return nil
}

func (*ProductHandler) ProductAdd(ctx context.Context, req *pb.ProductAddRequest, res *pb.ProductAddResponse) error {
	product := models.ProductModel{
		Name:       req.Name,
		Num:        req.Num,
		Price:      req.Price,
		Unit:       req.Unit,
		Desc:       req.Desc,
		Pic:        req.Pic,
		CreateTime: time.Now(), // 创建时间,不去使用请求创建时间，避免用户传入错误的时间，避免数据不一致
	}
	result := data_source.Db.Create(&product)
	if result.Error != nil {
		res.Code = 500
		res.Msg = "添加失败"
		return result.Error
	}
	res.Code = 200
	res.Msg = "添加成功"
	return nil
}

func (*ProductHandler) ProductDel(ctx context.Context, req *pb.ProductDelRequest, res *pb.ProductDelResponse) error {
	result := data_source.Db.Where("id=?", req.Id).Delete(&models.ProductModel{}) // 删除数据，Delete参数用于指定删除的表
	if result.Error != nil {
		res.Code = 500
		res.Msg = "删除失败"
		return result.Error
	}
	res.Code = 200
	res.Msg = "删除成功"
	return nil
}

// 这个方法用于编辑产品数据,用于查找待编辑的产品数据
func (*ProductHandler) ProductToEdit(ctx context.Context, req *pb.ProductToEditRequest, res *pb.ProductToEditResponse) error {
	product := models.ProductModel{}
	result := data_source.Db.Where("id=?", req.Id).Find(&product)
	if result.Error != nil {
		res.Code = 500
		res.Msg = "没有对应的数据"
		return result.Error
	}
	res.Code = 200
	res.Msg = "获取数据成功"
	res.Product = &pb.Product{
		Id:         product.Id,
		Name:       product.Name,
		Num:        product.Num,
		Price:      product.Price,
		Unit:       product.Unit,
		Desc:       product.Desc,
		Pic:        product.Pic,
		CreateTime: product.CreateTime.Format("2006-01-02 15:04:05"),
	}
	return nil
}

// 编辑产品数据,
func (*ProductHandler) ProductDoEdit(ctx context.Context, req *pb.ProductEditRequest, res *pb.ProductEditResponse) error {
	product := models.ProductModel{
		Id:         req.Id,
		Name:       req.Name,
		Num:        req.Num,
		Price:      req.Price,
		Unit:       req.Unit,
		Desc:       req.Desc,
		Pic:        req.Pic,
		CreateTime: time.Now(),
	}
	result := data_source.Db.Save(&product) // Save方法用于更新数据
	if result.Error != nil {
		res.Code = 500
		res.Msg = "编辑失败"
		return result.Error
	}
	res.Code = 200
	res.Msg = "编辑成功"
	return nil
}
