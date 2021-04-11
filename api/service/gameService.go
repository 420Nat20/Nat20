package service

import (
	"nat-20/data"

	"gorm.io/gorm"
)

func GetAllGameModels(db *gorm.DB) ([]data.GameModel, error) {
	var items []data.GameModel
	err := db.Find(&items).Error
	return items, err
}

func GetGameModelByServerID(db *gorm.DB, id int) (data.GameModel, error) {
	var item data.GameModel
	err := db.Where("server_id = ?", id).First(&item).Error
	return item, err
}

func CreateGameModel(db *gorm.DB, item *data.GameModel) error {
	err := db.Create(item).Error
	return err
}

func UpdateGameModel(db *gorm.DB, item *data.GameModel) error {
	err := db.Model(item).Updates(*item).Error
	return err
}

func DeleteGameModel(db *gorm.DB, id int) error {
	err := db.Where("server_id = ?", id).Delete(&data.GameModel{}).Error
	return err
}
