package api

import (
	"we_a_family/we_a_family/api/image_api"
	"we_a_family/we_a_family/api/member_api"
	"we_a_family/we_a_family/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	MemberApi   member_api.MemberApi
	ImagesApi   image_api.ImageApi
}

var ApiGroupApp = new(ApiGroup)
