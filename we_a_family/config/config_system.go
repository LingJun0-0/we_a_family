package config

import (
	"strconv"
)

type System struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

func (s *System) Addr() string {
	return s.Host + ":" + strconv.Itoa(s.Port) + s.Env
}
