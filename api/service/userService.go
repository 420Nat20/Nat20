package service

import (
	"nat-20/data"

	"gorm.io/gorm"
)

func GetAllUserModels(db *gorm.DB) ([]data.UserModel, error) {
	var items []data.UserModel
	err := db.Find(&items).Error
	return items, err
}

func GetUserModelByID(db *gorm.DB, id int) (data.UserModel, error) {
	var item data.UserModel
	err := db.First(&item, id).Error
	return item, err
}

func CreateUserModel(db *gorm.DB, item *data.UserModel) error {
	err := db.Create(item).Error
	return err
}

func UpdateUserModel(db *gorm.DB, item *data.UserModel) error {
	err := db.Model(item).Updates(*item).Error
	return err
}

func DeleteUserModel(db *gorm.DB, id int) error {
	err := db.Delete(&data.UserModel{}, id).Error
	return err
}
