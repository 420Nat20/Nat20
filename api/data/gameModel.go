package data

import "gorm.io/gorm"

type GameModel struct {
	gorm.Model
	ServerID int `gorm:"not null;unique"`

	DM        int
	Users     []UserModel
	Locations []LocationModel
}

type LocationModel struct {
	gorm.Model
	Name             string
	EventDescription string
	SubLocations     []SubLocationModel
	Visited          bool
	GameModelID      uint
}

type SubLocationModel struct {
	gorm.Model
	Name             string
	EventDescription string
	Visited          bool
	LocationModelID  uint
}
