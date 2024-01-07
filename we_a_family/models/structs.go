package Models

import (
	"github.com/dgrijalva/jwt-go"
)

// Member 数据库模型 - 标签
type Member struct {
	Username  int       `gorm:"not null;primary_key"`
	Password  string    `gorm:"type:varchar(255);not null"`
	CreatedAt string    `gorm:"column:created_at;type:datetime"`
	UpdatedAt string    `gorm:"column:updated_at;type:datetime"`
	Tags      []Tag     `gorm:"many2many:member_tags"`
	Pictures  []Picture `gorm:"many2many:member_pictures"`
}

// Picture 数据库模型 - 标签
type Picture struct {
	Id        int      `gorm:"primary_key"`
	Name      string   //照片名字
	Url       string   //照片连接
	Code      string   //照片标识
	CreatedAt string   `gorm:"column:created_at;type:datetime"`
	UpdatedAt string   `gorm:"column:updated_at;type:datetime"`
	Tags      []Tag    `gorm:"many2many:tag_pictures"`
	Members   []Member `gorm:"many2many:member_pictures"`
}

// Tag 数据库模型 - 标签
type Tag struct {
	Id          int `gorm:"primary_key"`
	Name        string
	Description string
	CreatedAt   string    `gorm:"column:created_at;type:datetime"`
	UpdatedAt   string    `gorm:"column:updated_at;type:datetime"`
	Pictures    []Picture `gorm:"many2many:tag_pictures"`
	Members     []Member  `gorm:"many2many:member_tags"`
}

// Permission 对资源的权限描述
type Permission struct {
	ResourceId  int `gorm:"not null;"`
	Code        int
	Description string
}

// AuthPayload 鉴权负载结构
type AuthPayload struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

// FileResponse 文件响应结构
type FileResponse struct {
	FileName  string `json:"file_name"`  //文件名
	IsSuccess bool   `json:"is_success"` //是否成功
	Msg       string `json:"msg"`        //消息
}
