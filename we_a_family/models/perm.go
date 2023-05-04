package Models

import "time"

type Perm struct {
	Id          int       `gorm:"primary_key"`
	UserId      int       //用户ID
	PermType    string    `gorm:"not null"`
	Code        int       //权限ID
	TagId       int       //标签ID
	ResourcesId int       //图片资源ID
	Created_at  time.Time `gorm:"column:created_at;type:datetime"`
	Updated_at  time.Time `gorm:"column:updated_at;type:datetime"`
}
