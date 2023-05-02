package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"we_a_family/we_a_family/models/res"
)

const file = "we_a_family/models/res/err_code.josn"

type ErrMap map[res.ErrorCode]string

func main() {
	byteData, err := os.ReadFile(file)
	if err != nil {
		logrus.Error(err)
		return
	}
	var errMap = ErrMap{}
	err = json.Unmarshal(byteData, &errMap)
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println(errMap)
	fmt.Printf(errMap[res.SettingsError])
}
