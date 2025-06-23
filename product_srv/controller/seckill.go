package controller

import (
	"context"
	"product_srv/data_source"
	models "product_srv/models"
	pb "product_srv/proto/seckill"
	"time"
)

type SecKillsHandler struct{}

func (s *SecKillsHandler) SecKillList(ctx context.Context, req *pb.SecKillsRequest, res *pb.SecKillsResponse) error {

	// 获取秒杀列表
	currentPage := req.CurrentPage
	pageSize := req.Pagesize

	seckills := []models.SecKills{}

	offsetNum := pageSize * (currentPage - 1)
	result := data_source.Db.Offset(int(offsetNum)).Limit(int(pageSize)).Find(&seckills)

	if result.Error != nil {
		res.Code = 500
		res.Msg = "获取秒杀列表失败"
		return result.Error
	}

	res.Code = 200
	res.Msg = "获取秒杀列表成功"

	var count int64
	seckills_count := []models.SecKills{}
	data_source.Db.Find(&seckills_count).Count(&count)
	seckills_rep := []*pb.SecKill{}

	for _, seckill := range seckills {
		seckill_rep := pb.SecKill{}
		seckill_rep.Id = int64(seckill.Id)
		seckill_rep.Name = seckill.Name
		seckill_rep.Price = seckill.Price
		seckill_rep.Num = int64(seckill.Num)
		product := models.ProductModel{
			Id: seckill.PId,
		}
		data_source.Db.First(&product) // 根据ID查询产品，结果存储在product中，因为gorm中一行数据也算一个抽象表，这个表的ID就是产品的ID
		seckill_rep.Pid = int64(seckill.PId)
		seckill_rep.Pname = product.Name
		seckill_rep.StartTime = seckill.StartTime.Format("2006-01-02 15:04:05")
		seckill_rep.EndTime = seckill.EndTime.Format("2006-01-02 15:04:05")
		seckill_rep.CreateTime = seckill.CreateTime.Format("2006-01-02 15:04:05")

		seckills_rep = append(seckills_rep, &seckill_rep)
	}

	res.Current = currentPage
	res.PageSize = pageSize
	res.Total = int64(count)
	res.Seckills = seckills_rep
	return nil

}
func (*SecKillsHandler) GetProducts(ctx context.Context, req *pb.ProductRequest, res *pb.ProductResponse) error {
	// 获取产品列表
	products := []models.ProductModel{}
	result := data_source.Db.Find(&products)
	if result.Error != nil {
		res.Code = 500
		res.Msg = "获取产品列表失败"
		return result.Error
	}
	res.Code = 200
	res.Msg = "获取产品列表成功"
	var products_rep []*pb.SecKillProductRequest
	for _, product := range products {
		product_rep := pb.SecKillProductRequest{}
		product_rep.Id = int64(product.Id)
		product_rep.Pname = product.Name
		products_rep = append(products_rep, &product_rep)
		products_rep = append(products_rep, &product_rep)
	}
	res.Products = products_rep
	return nil
}

// 请求把秒杀产品添加到数据库
func (*SecKillsHandler) SecKillAdd(ctx context.Context, req *pb.SecKill, res *pb.SecKillResponse) error {
	name := req.Name
	price := req.Price
	num := req.Num
	pid := req.Pid
	start_time := req.StartTime
	end_time := req.EndTime

	time_start_time, _ := time.Parse("2006-01-02 15:04:05", start_time)
	time_end_time, _ := time.Parse("2006-01-02 15:04:05", end_time)
	seckill := models.SecKills{
		Name:       name,
		Price:      price,
		Num:        num,
		PId:        pid,
		StartTime:  time_start_time,
		EndTime:    time_end_time,
		Status:     0,
		CreateTime: time.Now(),
	}
	result := data_source.Db.Create(&seckill)

	if result.Error != nil {
		res.Code = 500
		res.Msg = "添加失败"
	}

	res.Code = 200
	res.Msg = "添加成功"
	return nil
}

func (*SecKillsHandler) SecKillDel(ctx context.Context, req *pb.SecKillDelRequest, res *pb.SecKillResponse) error {
	id := req.Id
	result := data_source.Db.Delete(&models.SecKills{}, id)
	if result.Error != nil {
		res.Code = 500
		res.Msg = "删除失败"
		return result.Error
	}
	res.Code = 200
	res.Msg = "删除成功"
	return nil
}
func (*SecKillsHandler) SecKillToEdit(ctx context.Context, req *pb.SecKillDelRequest, res *pb.SecKillToEditResponse) error {
	id := req.Id
	seckill := models.SecKills{}
	result := data_source.Db.First(&seckill, id)
	if result.Error != nil {
		res.Code = 500
		res.Msg = "获取秒杀产品失败"
		return result.Error
	}
	res.Code = 200
	res.Msg = "获取秒杀产品成功"
	seckill_rep := pb.SecKill{}
	seckill_rep.Id = int64(seckill.Id)
	seckill_rep.Name = seckill.Name
	seckill_rep.Price = seckill.Price
	seckill_rep.Num = int64(seckill.Num)
	product := models.ProductModel{
		Id: seckill.PId,
	}
	data_source.Db.First(&product) // 根据ID查询产品，结果存储在product中，因为gorm中一行数据也算一个抽象表，这个表的ID就是产品的ID
	seckill_rep.Pid = int64(seckill.PId)
	seckill_rep.Pname = product.Name
	seckill_rep.StartTime = seckill.StartTime.Format("2006-01-02 15:04:05")
	seckill_rep.EndTime = seckill.EndTime.Format("2006-01-02 15:04:05")
	seckill_rep.CreateTime = seckill.CreateTime.Format("2006-01-02 15:04:05")
	res.Seckill = &seckill_rep
	return nil
}

// 秒杀活动更新
func (*SecKillsHandler) SecKillDoEdit(ctx context.Context, req *pb.SecKill, res *pb.SecKillResponse) error {
	id := req.Id
	name := req.Name
	price := req.Price
	num := req.Num
	pid := req.Pid
	start_time := req.StartTime
	end_time := req.EndTime

	time_start_time, _ := time.Parse("2006-01-02 15:04:05", start_time)
	time_end_time, _ := time.Parse("2006-01-02 15:04:05", end_time)

	seckill := models.SecKills{
		Id:         id,
		Name:       name,
		Price:      price,
		Num:        num,
		PId:        pid,
		StartTime:  time_start_time,
		EndTime:    time_end_time,
		Status:     0,
		CreateTime: time.Now(),
	}
	result := data_source.Db.Save(&seckill)

	if result.Error != nil {
		res.Code = 500
		res.Msg = "更新失败"
		return result.Error
	}
	res.Code = 200
	res.Msg = "更新成功"
	return nil

}

/*
// 秒杀活动列表
*/
func (*SecKillsHandler) FrontSecKillList(ctx context.Context, req *pb.FrontSecKillRequest, res *pb.FrontSecKillResponse) error {
	tomorrow_time := time.Now().Add(24 * time.Hour)
	//2025-01-09 10:37:33
	//tomorrow_time_str := tomorrow_time.Format("2006-01-02 15:04:05")

	currentPage := req.CurrentPage
	pagesize := req.Pagesize

	offsetNum := pagesize * (currentPage - 1)
	seckills := []models.SecKills{}
	result := data_source.Db.Where("start_time <= ?", tomorrow_time).Where("status = ?", 0).Limit(int(pagesize)).Offset(int(offsetNum)).Find(&seckills)

	if result.Error != nil {
		res.Code = 500
		res.Msg = "获取秒杀列表失败"
		return result.Error
	}

	seckills_rep := []*pb.SecKill{}

	for _, seckill := range seckills {
		seckill_rep := pb.SecKill{}
		seckill_rep.Id = int64(seckill.Id)
		seckill_rep.Name = seckill.Name
		seckill_rep.Price = seckill.Price
		seckill_rep.Num = int64(seckill.Num)
		seckill_rep.Pid = int64(seckill.PId)
		product := models.ProductModel{}
		data_source.Db.Where("id = ?", seckill.PId).Find(&product)
		seckill_rep.Pname = product.Name
		seckill_rep.Pic = product.Pic
		seckill_rep.PPrice = product.Price
		seckill_rep.Pdesc = product.Desc
		seckill_rep.StartTime = seckill.StartTime.Format("2006-01-02 15:04:05")
		seckill_rep.EndTime = seckill.EndTime.Format("2006-01-02 15:04:05")
		seckill_rep.CreateTime = seckill.CreateTime.Format("2006-01-02 15:04:05")

		seckills_rep = append(seckills_rep, &seckill_rep)
	}

	seckills_count := []models.SecKills{}
	var count int64
	data_source.Db.Where("start_time <= ?", tomorrow_time).Where("status = ?", 0).Find(&seckills_count).Count(&count)

	res.Code = 200
	res.Msg = "成功"
	res.Current = currentPage
	res.PageSize = pagesize
	res.TotalPage = (count + pagesize - 1) / pagesize
	res.SeckillList = seckills_rep

	return nil

}

func (*SecKillsHandler) FrontSecKillDetail(ctx *context.Context, req *pb.SecKillDelRequest, res *pb.FrongSecKillDetailResponse) error {

	id := req.Id

	seckill := models.SecKills{}
	result := data_source.Db.Where("id = ?", id).Find(&seckill)
	if result.Error != nil {
		res.Code = 500
		res.Msg = "获取秒杀产品失败"
		return result.Error
	}

	product := models.ProductModel{}

	data_source.Db.Where("id = ?", seckill.PId).Find(&product)
	seckill_rep := pb.SecKill{
		Id:         int64(seckill.Id),
		Name:       seckill.Name,
		Num:        int64(seckill.Num),
		Price:      seckill.Price,
		Pid:        int64(seckill.PId),
		Pname:      product.Name,
		Pic:        product.Pic,
		PPrice:     product.Price,
		Pdesc:      product.Desc,
		Unit:       product.Unit,
		StartTime:  seckill.StartTime.Format("2006-01-02 15:04:05"),
		EndTime:    seckill.EndTime.Format("2006-01-02 15:04:05"),
		CreateTime: seckill.CreateTime.Format("2006-01-02 15:04:05"),
	}
	res.Code = 200
	res.Msg = "获取秒杀产品成功"
	res.Seckill = &seckill_rep
	return nil
}
