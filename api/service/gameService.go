package service

import (
	"github.com/420Nat20/Nat20/nat-20/data/model"

	"gorm.io/gorm"
)

func GetAllGameModels(db *gorm.DB) ([]model.GameModel, error) {
	var items []model.GameModel
	err := db.Find(&items).Error
	return items, err
}

func GetGameModelByServerID(db *gorm.DB, id int) (model.GameModel, error) {
	var item model.GameModel
	err := db.Where("server_id = ?", id).First(&item).Error
	return item, err
}

func CreateGameModel(db *gorm.DB, item *model.GameModel) error {
	err := db.Create(item).Error
	return err
}

func UpdateGameModel(db *gorm.DB, item *model.GameModel) error {
	err := db.Model(item).Updates(*item).Error
	return err
}

func DeleteGameModel(db *gorm.DB, id int) error {
	err := db.Where("server_id = ?", id).Delete(&model.GameModel{}).Error
	return err
}
