package api

import (
	"we_a_family/we_a_family/api/image_api"
	"we_a_family/we_a_family/api/member_api"
	"we_a_family/we_a_family/api/settings_api"
	"we_a_family/we_a_family/api/tag_api"
)

type Group struct {
	SettingsApi settings_api.SettingsApi
	MemberApi   member_api.MemberApi
	ImagesApi   image_api.ImageApi
	TagApi      tag_api.TagApi
}

var GroupApp = new(Group)
