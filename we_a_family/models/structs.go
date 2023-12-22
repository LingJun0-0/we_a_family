package Models

import (
	"github.com/dgrijalva/jwt-go"
)

// Member 数据库模型 - 标签
type Member struct {
	Id        int    `gorm:"primary_key"`
	Username  string `gorm:"not null;unique_index"`
	Password  string `gorm:"type:varchar(255);not null"`
	Status    int    //member状态
	Deleted   bool   `gorm:"column:deleted;type:tinyint(1)"`
	CreatedAt string `gorm:"column:created_at;type:datetime"`
	UpdatedAt string `gorm:"column:updated_at;type:datetime"`
}

// Picture 数据库模型 - 标签
type Picture struct {
	Id        int    `gorm:"primary_key"`
	Name      string //照片名字
	Url       string //照片连接
	Code      string //照片标识
	CreatedAt string `gorm:"column:created_at;type:datetime"`
	UpdatedAt string `gorm:"column:updated_at;type:datetime"`
	Tags      []Tag  `gorm:"many2many:picture_tags"`
}

// Tag 数据库模型 - 标签
type Tag struct {
	Id          int `gorm:"primary_key"`
	Name        string
	Description string
	OwnerID     int
	Deleted     bool      `gorm:"column:deleted;type:tinyint(1)"`
	CreatedAt   string    `gorm:"column:created_at;type:datetime"`
	UpdatedAt   string    `gorm:"column:updated_at;type:datetime"`
	Admins      []Member  `gorm:"many2many:tag_admins;"`
	Viewers     []Member  `gorm:"many2many:tag_viewers;"`
	Pictures    []Picture `gorm:"many2many:picture_tags"`
}

// AuthPayload 鉴权负载结构
type AuthPayload struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}
