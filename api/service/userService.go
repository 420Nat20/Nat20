package service

import (
	"github.com/420Nat20/Nat20/nat-20/data/model"

	"gorm.io/gorm"
)

func GetAllUserModels(db *gorm.DB) ([]model.UserModel, error) {
	var items []model.UserModel
	err := db.Find(&items).Error
	return items, err
}

func GetUserModelByDiscordID(db *gorm.DB, id int) (model.UserModel, error) {
	var item model.UserModel
	err := db.Where("discord_id = ?", id).First(&item).Error
	return item, err
}

func CreateUserModel(db *gorm.DB, gameId int, item *model.UserModel) error {
	game, err := GetGameModelByServerID(db, gameId)
	if err != nil {
		return err
	}
	err = db.Model(&game).Association("Users").Append(item)
	return err
}

func UpdateUserModel(db *gorm.DB, gameId int, item *model.UserModel) error {
	game, err := GetGameModelByServerID(db, gameId)
	if err != nil {
		return err
	}
	err = db.Model(&game).Association("Users").Replace(item)
	return err
}

func DeleteUserModel(db *gorm.DB, id int) error {
	err := db.Delete(&model.UserModel{}, id).Error
	return err
}
