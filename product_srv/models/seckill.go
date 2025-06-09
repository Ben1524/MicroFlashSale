package models

import "time"

type SecKills struct {
	Id        int64
	Name      string
	Price     float32 `gorm:"type:decimal(11,2)"`
	Num       int64
	PId       int64
	StartTime time.Time
	EndTime   time.Time
	// 1表示下架，0表示未下架
	Status     int
	CreateTime time.Time
}

func (SecKills) TableName() string {
	return "product_seckills"

}
