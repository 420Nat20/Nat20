package data

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	GameModelID uint

	DiscordID string `gorm:"unique"`

	Name       string `gorm:"not null"`
	Class      string `gorm:"not null"`
	Background string `gorm:"not null"`
	PlayerName string `gorm:"not null"`
	Race       string `gorm:"not null"`
	Alignment  string `gorm:"not null"`

	Strength     int `gorm:"not null"`
	Dexterity    int `gorm:"not null"`
	Constitution int `gorm:"not null"`
	Intelligence int `gorm:"not null"`
	Wisdom       int `gorm:"not null"`
	Charisma     int `gorm:"not null"`

	TraitOne string
	TraitTwo string
	Ideal    string
	Bond     string
	Flaw     string
}
