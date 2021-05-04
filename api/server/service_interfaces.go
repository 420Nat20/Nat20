package server

import (
	"github.com/420Nat20/Nat20/nat-20/data/models"
)

type CampaignService interface {
	GetCampaign(id int) (*models.Campaign, error)
	GetAllCampaigns() (models.CampaignSlice, error)
	CreateCampaign(campaign *models.Campaign) (*models.Campaign, error)
	UpdateCampaign(id int, json map[string]interface{}) (int, error)
	DeleteCampaign(id int) error
}

type UserService interface {
	GetUser(id int) (*models.User, error)
	GetAllUsers() (models.UserSlice, error)
	CreateUser(user *models.User) error
	UpdateUser(id int, json map[string]interface{}) (int, error)
	DeleteUser(id int) error
}

type LocationService interface {
	GetLocation(id int) (*models.Location, error)
	GetAllLocations(gameId int) (models.LocationSlice, error)
	CreateLocation(gameId int, location *models.Location) (*models.Location, error)
	UpdateLocation(gameId int, id int, json map[string]interface{}) (int, error)
	DeleteLocation(id int) error

	GetSubLocation(id int) (*models.SubLocation, error)
	GetAllSubLocation(gameId int, locationId int) (models.SubLocationSlice, error)
	CreateSubLocation(gameId int, locationId int, location *models.SubLocation) (*models.SubLocation, error)
	UpdateSubLocation(gameId int, locationId int, id int, json map[string]interface{}) (int, error)
	DeleteSubLocation(id int) error
}
