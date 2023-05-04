package Models

import "time"

type Tag struct {
	Id         int       `gorm:"primary_key"`
	Name       string    `gorm:"type:varchar(100);not null"`
	OwnerId    int       //创建者ID
	PictureId  int       //图片资源ID
	Created_at time.Time `gorm:"column:created_at;type:datetime"`
	Updated_at time.Time `gorm:"column:updated_at;type:datetime"`
	Perm       []Perm    `gorm:"foreignKey:TagId"`
}
