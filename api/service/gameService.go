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
	err := db.Delete(&data.GameModel{}, id).Error
	return err
}

// Location CRUD

func GetAllLocationModels(db *gorm.DB) ([]data.LocationModel, error) {
	var items []data.LocationModel
	err := db.Find(&items).Error
	return items, err
}

func GetLocationModelByID(db *gorm.DB, id int) (data.LocationModel, error) {
	var item data.LocationModel
	err := db.First(&item, id).Error
	return item, err
}

func CreateLocationModel(db *gorm.DB, item *data.LocationModel) error {
	err := db.Create(item).Error
	return err
}

func UpdateLocationModel(db *gorm.DB, item *data.LocationModel) error {
	err := db.Model(item).Updates(*item).Error
	return err
}

func DeleteLocationModel(db *gorm.DB, id int) error {
	err := db.Delete(&data.LocationModel{}, id).Error
	return err
}

// SubLocation CRUD

func GetAllSubLocationModels(db *gorm.DB) ([]data.SubLocationModel, error) {
	var items []data.SubLocationModel
	err := db.Find(&items).Error
	return items, err
}

func GetSubLocationModelByID(db *gorm.DB, id int) (data.SubLocationModel, error) {
	var item data.SubLocationModel
	err := db.First(&item, id).Error
	return item, err
}

func CreateSubLocationModel(db *gorm.DB, item *data.SubLocationModel) error {
	err := db.Create(item).Error
	return err
}

func UpdateSubLocationModel(db *gorm.DB, item *data.SubLocationModel) error {
	err := db.Model(item).Updates(*item).Error
	return err
}

func DeleteSubLocationModel(db *gorm.DB, id int) error {
	err := db.Delete(&data.SubLocationModel{}, id).Error
	return err
}
