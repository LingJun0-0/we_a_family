package config

type Download struct {
	Size int    `yaml:"size" json:"size"` //图片下载的大小
	Path string `yaml:"path" json:"path"` //图片下载的目录
}
