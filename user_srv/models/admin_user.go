package models

import "time"

type AdminUserModel struct {
	Id         int
	UserName   string
	Password   string
	Desc       string
	Status     int
	CreateTime time.Time
}

func (AdminUserModel) TableName() string {
	return "admin_users"

}
