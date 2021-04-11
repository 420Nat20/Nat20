package service

import (
	"nat-20/data"

	"gorm.io/gorm"
)

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

func CreateLocationModel(db *gorm.DB, gameId int, item *data.LocationModel) error {
	game, err := GetGameModelByServerID(db, gameId)
	if err != nil {
		return err
	}
	err = db.Model(&game).Association("Locations").Append(item)
	return err
}

func UpdateLocationModel(db *gorm.DB, id int, item *data.LocationModel) error {
	location, err := GetLocationModelByID(db, id)
	if err != nil {
		return err
	}
	err = db.Model(&location).Updates(item).Error
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

func CreateSubLocationModel(db *gorm.DB, locationId int, item *data.SubLocationModel) error {
	location, err := GetLocationModelByID(db, locationId)
	if err != nil {
		return err
	}
	err = db.Model(&location).Association("SubLocations").Append(item)
	return err
}

func UpdateSubLocationModel(db *gorm.DB, subLocationId int, item *data.SubLocationModel) error {
	subLocation, err := GetSubLocationModelByID(db, subLocationId)
	if err != nil {
		return err
	}
	err = db.Model(&subLocation).Updates(item).Error
	return err
}

func DeleteSubLocationModel(db *gorm.DB, id int) error {
	err := db.Delete(&data.SubLocationModel{}, id).Error
	return err
}
