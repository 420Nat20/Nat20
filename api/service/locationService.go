package service

import (
	"github.com/420Nat20/Nat20/nat-20/data/model"

	"gorm.io/gorm"
)

// Location CRUD
func GetAllLocationModels(db *gorm.DB) ([]model.LocationModel, error) {
	var items []model.LocationModel
	err := db.Find(&items).Error
	return items, err
}

func GetLocationModelByID(db *gorm.DB, id int) (model.LocationModel, error) {
	var item model.LocationModel
	err := db.First(&item, id).Error
	return item, err
}

func CreateLocationModel(db *gorm.DB, gameId int, item *model.LocationModel) error {
	game, err := GetGameModelByServerID(db, gameId)
	if err != nil {
		return err
	}
	err = db.Model(&game).Association("Locations").Append(item)
	return err
}

func UpdateLocationModel(db *gorm.DB, id int, item *model.LocationModel) error {
	location, err := GetLocationModelByID(db, id)
	if err != nil {
		return err
	}
	err = db.Model(&location).Updates(item).Error
	return err
}

func DeleteLocationModel(db *gorm.DB, id int) error {
	err := db.Delete(&model.LocationModel{}, id).Error
	return err
}

// SubLocation CRUD
func GetAllSubLocationModels(db *gorm.DB, locationId int) ([]model.SubLocationModel, error) {
	location, err := GetLocationModelByID(db, locationId)
	if err != nil {
		return nil, err
	}

	var items []model.SubLocationModel
	err = db.Where("location_model_id = ?", location.ID).Find(&items).Error
	return items, err
}

func GetSubLocationModelByID(db *gorm.DB, id int) (model.SubLocationModel, error) {
	var item model.SubLocationModel
	err := db.First(&item, id).Error
	return item, err
}

func CreateSubLocationModel(db *gorm.DB, locationId int, item *model.SubLocationModel) error {
	location, err := GetLocationModelByID(db, locationId)
	if err != nil {
		return err
	}
	err = db.Model(&location).Association("SubLocations").Append(item)
	return err
}

func UpdateSubLocationModel(db *gorm.DB, subLocationId int, item *model.SubLocationModel) error {
	subLocation, err := GetSubLocationModelByID(db, subLocationId)
	if err != nil {
		return err
	}
	err = db.Model(&subLocation).Updates(item).Error
	return err
}

func DeleteSubLocationModel(db *gorm.DB, id int) error {
	err := db.Delete(&model.SubLocationModel{}, id).Error
	return err
}
