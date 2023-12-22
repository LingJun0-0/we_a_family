package config

type Config struct {
	Mysql     Mysql     `yaml:"mysql"`
	Logger    Logger    `yaml:"logger"`
	System    System    `yaml:"system"`
	Upload    Upload    `yaml:"upload"`
	Download  Download  `yaml:"download"`
	JwtSecret JwtSecret `yaml:"JwtSecret"`
}
