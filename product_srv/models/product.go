package models

import "time"

type ProductModel struct {
	Id         int64
	Name       string
	Price      float32 `gorm:"type:decimal(11,2)"`
	Num        int64
	Unit       string
	Pic        string
	Desc       string
	CreateTime time.Time

	SecKills []SecKills `gorm:"ForeignKey:PId;AssiciationForeignKey:Id"` // 有多个秒杀活动
}

func (ProductModel) TableName() string {
	return "products"

}
