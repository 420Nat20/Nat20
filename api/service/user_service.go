package service

import (
	"context"
	"database/sql"
	"github.com/420Nat20/Nat20/nat-20/data/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserService struct {
	Ctx context.Context
	DB  *sql.DB
}

func (u UserService) GetUser(id int) (*models.User, error) {
	user, err := models.FindUser(u.Ctx, u.DB, id)
	if err != nil {
		return &models.User{}, err
	}
	return user, nil
}

func (u UserService) GetAllUsers() (models.UserSlice, error) {
	all, err := models.Users().All(u.Ctx, u.DB)
	if err != nil {
		return models.UserSlice{}, err
	}
	return all, nil
}

func (u UserService) CreateUser(user *models.User) error {
	err := user.Insert(u.Ctx, u.DB, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (u UserService) UpdateUser(id int, json map[string]interface{}) (int, error) {
	existingUser, err := u.GetUser(id)
	if err != nil {
		return id, err
	}

	existingUser.DiscordID = json[models.UserColumns.DiscordID].(string)

	_, err = existingUser.Update(u.Ctx, u.DB, boil.Infer())
	if err != nil {
		return id, err
	}

	return id, nil
}

func (u UserService) DeleteUser(id int) error {
	user, err := models.FindUser(u.Ctx, u.DB, id)
	if err != nil {
		return err
	}

	_, err = user.Delete(u.Ctx, u.DB)
	if err != nil {
		return err
	}
	return nil
}
