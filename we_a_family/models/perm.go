package Models

type Perm struct {
	Id          int    `gorm:"primary_key"`
	UserId      int    //用户ID
	PermType    string `gorm:"not null"`
	Code        int    //权限ID
	TagId       int    //标签ID
	ResourcesId int    //图片资源ID
	CreatedAt   string `gorm:"column:created_at;type:datetime"`
	UpdatedAt   string `gorm:"column:updated_at;type:datetime"`
}
