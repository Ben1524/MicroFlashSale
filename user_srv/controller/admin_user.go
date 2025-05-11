package controller

import (
	"context"
	"user_srv/data_source"
	models "user_srv/models"
	pb "user_srv/proto/admin_user"
)

//type AdminUserHandler interface {
//	AdminUserLogin(context.Context, *AdminUserRequest, *AdminUserResponse) error
//	FrontUserList(context.Context, *FrontUsersRequest, *FrontUsersResponse) error
//}

type AdminUserHandler struct{}

func (a *AdminUserHandler) AdminUserLogin(ctx context.Context, req *pb.AdminUserRequest, rsp *pb.AdminUserResponse) error {
	UserName := req.Username
	Password := req.Password
	var count int64
	data_source.Db.Where("user_name = ? AND password = ?", UserName, Password).First(&models.AdminUserModel{}).Count(&count)
	if count > 0 {
		rsp.Msg = "登录成功"
		rsp.UserName = UserName
		rsp.Code = 200
	} else {
		rsp.Msg = "登录失败"
		rsp.Code = 500
	}
	return nil
}

func (a *AdminUserHandler) FrontUserList(ctx context.Context, req *pb.FrontUsersRequest, rsp *pb.FrontUsersResponse) error {
	currentPage := req.CurrentPage // 当前页
	pageSize := req.Pagesize       // 每页显示的条数
	var frontUsers []models.FrontUserModel
	var total int64
	data_source.Db.Model(&models.FrontUserModel{}).Count(&total)
	if err := data_source.Db.Offset(int((currentPage - 1) * pageSize)).Limit(int(pageSize)).Find(&frontUsers); err.Error != nil {
		rsp.Code = 500
		rsp.Msg = "获取前台用户列表失败"
		return err.Error
	}
	rsp.Code = 200
	rsp.Msg = "获取前台用户列表成功"
	for index, frontUser := range frontUsers {
		status_str := "正常"
		if frontUser.Status == 0 {
			status_str = "冻结"
		}
		create_time := frontUser.CreateTime.Format("2006-01-02 15:04:05")
		rsp.FrontUsers = append(rsp.FrontUsers, &pb.FrontUserMessage{
			Email:      frontUser.Email,
			Desc:       frontUser.Desc,
			Status:     status_str,
			CreateTime: create_time,
		})
		if index == len(frontUsers)-1 {
			rsp.Total = int32(total)
			rsp.Current = currentPage
			rsp.PageSize = pageSize
		}
	}
	return nil
}
