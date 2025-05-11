package controller

import (
	"context"
	"crypto"
	"github.com/patrickmn/go-cache"
	"math/rand"
	"time"
	"user_srv/data_source"
	models "user_srv/models"
	pb "user_srv/proto/front_user"
	"user_srv/utils"
)

type FrontUserHandler struct{}

var CodeKVCache = cache.New(60*time.Second, 10*time.Second) // 60秒过期，10秒清理一次，清理过期数据

func (*FrontUserHandler) FrontUserRegister(ctx context.Context, req *pb.FrontUserRequest, res *pb.FrontUserResponse) error {
	email := req.Email
	code := req.Code
	password := req.Password
	code2, is_ok := CodeKVCache.Get(email) // 从缓存中获取验证码,判断是否过期
	if !is_ok || code2 != code {
		res.Code = 500
		res.Msg = "验证码已过期或者不正确"
		return nil
	} else {
		pwd_to_md5 := crypto.MD5.New()     // md5加密
		pwd_to_md5.Write([]byte(password)) // 写入密码
		pwd_md5 := pwd_to_md5.Sum(nil)     // 返回加密结果
		new_front_user := &models.FrontUserModel{
			Email:      email,
			Password:   string(pwd_md5),
			Status:     1,
			CreateTime: time.Now(),
		}
		data_source.Db.Create(new_front_user)
		res.Code = 200
		res.Msg = "注册成功"
	}
	return nil
}

func (*FrontUserHandler) FrontUserSendEmail(ctx context.Context, req *pb.FrontUserMailRequest, res *pb.FrontUserResponse) error {
	email := req.Email
	// 随机生成6位数验证码
	nums := "0123456789"
	var code string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		code += string(nums[rand.Intn(len(nums))])
	}
	CodeKVCache.Set(email, code, cache.DefaultExpiration) // 缓存验证码
	res.Code = 200
	res.Msg = "验证码发送成功"
	if err := utils.SendEmail(email, "注册验证码", code); err != nil {
		res.Code = 500
		res.Msg = "验证码发送失败,请检查邮箱是否正确"
	}
	return nil
}

func (*FrontUserHandler) FrontUserLogin(ctx context.Context, req *pb.FrontUserRequest, res *pb.FrontUserResponse) error {
	email := req.Email
	password := req.Password
	pwd_to_md5 := crypto.MD5.New()     // md5加密
	pwd_to_md5.Write([]byte(password)) // 写入密码
	pwd_md5 := pwd_to_md5.Sum(nil)     // 返回加密结果
	var pb models.FrontUserModel

	var count int64
	data_source.Db.Where("email = ? AND password = ?", email, string(pwd_md5)).First(&pb).Count(&count)
	if pb.Id == 0 {
		res.Code = 500
		res.Msg = "登录失败"
	} else {
		res.Code = 200
		res.Msg = "登录成功"
		res.UserName = email
	}
	return nil
}
