package api

import (
	"we_a_family/we_a_family/api/member_api"
	"we_a_family/we_a_family/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	MemberApi   member_api.MemberApi
}

var ApiGroupApp = new(ApiGroup)
