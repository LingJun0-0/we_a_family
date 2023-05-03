package Models

import "time"

type Tag struct {
	Id         int         `gorm:"primary_key"`
	Name       string      `gorm:"type:varchar(100);not null"`
	MemberId   int         `gorm:"foreignKey:MemberId"` //创建者ID
	Created_at time.Time   `gorm:"tag_created_at"`
	Updated_at time.Time   `gorm:"tag_updated_at"`
	TagMember  []TagMember `gorm:"foreignKey:TagId"`
}

type TagMember struct {
	TagId    int `json:"tag_id"`
	MemberId int `json:"member_id"` //标签管理者
}
