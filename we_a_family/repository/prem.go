package repository

import (
	"we_a_family/we_a_family/global"
	Models "we_a_family/we_a_family/models"
)

func InsertOnePermission(resourceId, code int, description string) error {
	perm := Models.Permission{ResourceId: resourceId, Code: code, Description: description}
	err := global.DB.Model(Models.Permission{}).Create(&perm).Error
	return err
}

func UpdateOnePermission(resourceId, code int, description string) error {
	err := global.DB.Model(Models.Permission{}).
		Where(Models.Permission{
			ResourceId:  resourceId,
			Description: description}).
		Updates(Models.Permission{Code: code}).Error
	return err
}

func DeleteOnePermission(resourceId, code int, description string) error {
	err := global.DB.Model(Models.Permission{}).
		Delete(&Models.Permission{},
			Models.Permission{
				ResourceId:  resourceId,
				Code:        code,
				Description: description}).Error
	return err
}

func DelPermissionsByResourceId(resourceId int) error {
	var perm Models.Permission
	err := global.DB.Model(&perm).Delete(&perm, resourceId).Error
	return err
}

func GetOnePermissionByAllArguments(resourceId, code int, description string) (Models.Permission, error) {
	var perm Models.Permission
	err := global.DB.Model(Models.Permission{}).
		Where(Models.Permission{ResourceId: resourceId, Code: code, Description: description}).
		Find(&perm).Error
	return perm, err
}

func GetOnePermissionByResourceIdAndDescription(resourceId int, description string) (Models.Permission, error) {
	var perm Models.Permission
	err := global.DB.Model(Models.Permission{}).
		Where(Models.Permission{ResourceId: resourceId, Description: description}).
		Find(&perm).Error
	return perm, err
}

func GetPermissionsByResourceId(resourceId int) ([]Models.Permission, error) {
	var perms []Models.Permission
	err := global.DB.Model(Models.Permission{}).Where(Models.Permission{ResourceId: resourceId}).Find(&perms).Error
	return perms, err
}

func GetPermissionsByCode(code int) ([]Models.Permission, error) {
	var perms []Models.Permission
	err := global.DB.Model(Models.Permission{}).
		Where(Models.Permission{Code: code}).
		Find(&perms).Error
	return perms, err
}

func GetPermissionsByDescription(description string) ([]Models.Permission, error) {
	var perms []Models.Permission
	err := global.DB.Model(Models.Permission{}).Where(Models.Permission{Description: description}).Find(&perms).Error
	return perms, err
}

func GetAllPermissions() ([]Models.Permission, error) {
	var perms []Models.Permission
	err := global.DB.Model(Models.Permission{}).Find(&perms).Error
	return perms, err
}

func SamePermission(resourceId, code int, description string) bool {
	var count int64
	global.DB.Model(Models.Permission{}).
		Where(Models.Permission{ResourceId: resourceId, Code: code, Description: description}).
		Count(&count)
	return count > 0
}
