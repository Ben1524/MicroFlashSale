package models

import "time"

type Products struct {
	Id         int
	Name       string
	Price      float32 `gorm:"type:decimal(11,2)"`
	Num        int
	Unit       string
	Pic        string
	Desc       string
	CreateTime time.Time

	SecKills []SecKills `gorm:"ForeignKey:PId;AssiciationForeignKey:Id"` // 有多个秒杀活动
}

func (Products) TableName() string {
	return "products"

}
