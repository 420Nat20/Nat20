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

func CreateUserModel(db *gorm.DB, gameId int, item *data.UserModel) error {
	game, err := GetGameModelByServerID(db, gameId)
	if err != nil {
		return err
	}
	err = db.Model(&game).Association("Users").Append(item)
	return err
}

func UpdateUserModel(db *gorm.DB, gameId int, item *data.UserModel) error {
	game, err := GetGameModelByServerID(db, gameId)
	if err != nil {
		return err
	}
	err = db.Model(&game).Association("Users").Replace(item)
	return err
}

func DeleteUserModel(db *gorm.DB, id int) error {
	err := db.Delete(&data.UserModel{}, id).Error
	return err
}
