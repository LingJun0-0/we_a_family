package Models

type Tag struct {
	Id        int    `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(100);not null"`
	OwnerId   int    //创建者ID
	PictureId int    //图片资源ID
	CreatedAt string `gorm:"column:created_at;type:datetime"`
	UpdatedAt string `gorm:"column:updated_at;type:datetime"`
	Perm      []Perm `gorm:"foreignKey:TagId"`
}
