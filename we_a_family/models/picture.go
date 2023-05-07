package Models

type Picture struct {
	Id        int `gorm:"primary_key"`
	Name      string
	Url       string
	CreatedAt string `gorm:"column:created_at;type:datetime"`
	UpdatedAt string `gorm:"column:updated_at;type:datetime"`
	Tag       []Tag  `gorm:"foreignKey:PictureId"`
	Perm      []Perm `gorm:"foreignKey:ResourcesId"`
}
