package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"we_a_family/we_a_family/config"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
)
