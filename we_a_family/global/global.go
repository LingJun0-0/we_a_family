package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"we_a_family/we_a_family/config"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
)
