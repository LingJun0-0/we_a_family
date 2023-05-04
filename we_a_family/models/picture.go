package Models

import "time"

type Picture struct {
	Id         int `gorm:"primary_key"`
	Name       string
	Url        string
	Created_at time.Time
	Updated_at time.Time
	Tag        []Tag  `gorm:"column:created_at;type:datetime"` //创建者ID
	Perm       []Perm `gorm:"column:updated_at;type:datetime"` //创建者ID
}
