package service

import (
	"context"
	"database/sql"
	"github.com/420Nat20/Nat20/nat-20/data/models"
)

type LocationService struct {
	ctx context.Context
	db  *sql.DB
}

func (l LocationService) GetLocation(gameId int, id int) (models.Location, error) {
	panic("implement me")
}

func (l LocationService) GetAllLocations(gameId int) (models.LocationSlice, error) {
	panic("implement me")
}

func (l LocationService) CreateLocation(gameId int, location models.Location) (models.Location, error) {
	panic("implement me")
}

func (l LocationService) UpdateLocation(gameId int, id int, location models.Location) (models.Location, error) {
	panic("implement me")
}

func (l LocationService) DeleteLocation(id int) error {
	panic("implement me")
}

func (l LocationService) GetSubLocation(gameId int, locationId int, id int) (models.SubLocation, error) {
	panic("implement me")
}

func (l LocationService) GetAllSubLocation(gameId int, locationId int) (models.SubLocationSlice, error) {
	panic("implement me")
}

func (l LocationService) CreateSubLocation(gameId int, locationId int, location models.Location) (models.SubLocation, error) {
	panic("implement me")
}

func (l LocationService) UpdateSubLocation(gameId int, locationId int, id int, location models.Location) (models.SubLocation, error) {
	panic("implement me")
}

func (l LocationService) DeleteSubLocation(id int) error {
	panic("implement me")
}
