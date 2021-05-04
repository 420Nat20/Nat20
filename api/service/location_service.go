package service

import (
	"context"
	"database/sql"
	"github.com/420Nat20/Nat20/nat-20/data/models"
)

type LocationService struct {
	Ctx context.Context
	DB  *sql.DB
}

func (l LocationService) GetLocation(id int) (*models.Location, error) {
	panic("implement me")
}

func (l LocationService) GetAllLocations(gameId int) (models.LocationSlice, error) {
	panic("implement me")
}

func (l LocationService) CreateLocation(gameId int, location *models.Location) (*models.Location, error) {
	panic("implement me")
}

func (l LocationService) UpdateLocation(gameId int, id int, json map[string]interface{}) (int, error) {
	panic("implement me")
}

func (l LocationService) DeleteLocation(id int) error {
	panic("implement me")
}

func (l LocationService) GetSubLocation(id int) (*models.SubLocation, error) {
	panic("implement me")
}

func (l LocationService) GetAllSubLocation(gameId int, locationId int) (models.SubLocationSlice, error) {
	panic("implement me")
}

func (l LocationService) CreateSubLocation(gameId int, locationId int, location *models.SubLocation) (*models.SubLocation, error) {
	panic("implement me")
}

func (l LocationService) UpdateSubLocation(gameId int, locationId int, id int, json map[string]interface{}) (int, error) {
	panic("implement me")
}

func (l LocationService) DeleteSubLocation(id int) error {
	panic("implement me")
}
