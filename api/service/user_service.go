package service

import (
	"context"
	"database/sql"
	"github.com/420Nat20/Nat20/nat-20/common"
	"github.com/420Nat20/Nat20/nat-20/data/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"strconv"
)

type UserService struct {
	Ctx context.Context
	DB  *sql.DB
}

func (u UserService) GetUser(id int) (*models.User, error) {
	user, err := models.FindUser(u.Ctx, u.DB, id)
	if err != nil {
		return &models.User{}, common.NotFound.New("User not found")
	}
	return user, nil
}

func (u UserService) GetAllUsers() (models.UserSlice, error) {
	all, err := models.Users().All(u.Ctx, u.DB)
	if err != nil {
		return models.UserSlice{}, common.BadRequest.New("Failed to get all users")
	}
	return all, nil
}

func (u UserService) CreateUser(user *models.User) error {
	err := user.Insert(u.Ctx, u.DB, boil.Infer())
	if err != nil {
		return common.BadRequest.New("Failed to create user")
	}
	return nil
}

func (u UserService) UpdateUser(id int, json map[string]interface{}) (int, error) {
	existingUser, findErr := u.GetUser(id)
	if findErr != nil {
		return id, findErr
	}

	existingUser.DiscordID = json[models.UserColumns.DiscordID].(string)

	_, err := existingUser.Update(u.Ctx, u.DB, boil.Infer())
	if err != nil {
		updateFail := common.BadRequest.New("User not found")
		_ = common.AddErrorContext(updateFail, "id", strconv.Itoa(id))
		return id, updateFail
	}

	return id, nil
}

func (u UserService) DeleteUser(id int) error {
	user, err := models.FindUser(u.Ctx, u.DB, id)
	if err != nil {
		return common.NotFound.New("User not found")
	}

	_, err = user.Delete(u.Ctx, u.DB)
	if err != nil {
		deleteFail := common.BadRequest.New("User not found")
		_ = common.AddErrorContext(deleteFail, "id", strconv.Itoa(id))
		return deleteFail
	}
	return nil
}
