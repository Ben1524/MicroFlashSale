package models

import "time"

type FrontUserModel struct {
	Id         int
	Email      string
	Password   string
	Desc       string
	Status     int
	CreateTime time.Time
}

func (FrontUserModel) TableName() string {
	return "front_users"

}
