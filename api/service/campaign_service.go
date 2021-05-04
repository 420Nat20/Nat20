package service

import (
	"context"
	"database/sql"
	"github.com/420Nat20/Nat20/nat-20/common"
	"github.com/420Nat20/Nat20/nat-20/data/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"strconv"
)

type CampaignService struct {
	Ctx context.Context
	DB  *sql.DB
}

func (c CampaignService) GetCampaign(id int) (*models.Campaign, error) {
	campaign, err := models.FindCampaign(c.Ctx, c.DB, id)
	if err != nil {
		return &models.Campaign{}, common.NotFound.New("Campaign not found")
	}
	return campaign, nil
}

func (c CampaignService) GetAllCampaigns() (models.CampaignSlice, error) {
	all, err := models.Campaigns().All(c.Ctx, c.DB)
	if err != nil {
		return models.CampaignSlice{}, common.BadRequest.New("Failed to get campaigns")
	}
	return all, nil
}

func (c CampaignService) CreateCampaign(campaign *models.Campaign) error {
	err := campaign.Insert(c.Ctx, c.DB, boil.Infer())
	if err != nil {
		return common.BadRequest.New("Failed to create user")
	}
	return nil
}

func (c CampaignService) UpdateCampaign(id int, json map[string]interface{}) (int, error) {
	existingCampaign, findErr := c.GetCampaign(id)
	if findErr != nil {
		return id, findErr
	}

	existingCampaign.DMID = json[models.CampaignColumns.DMID].(int)

	_, err := existingCampaign.Update(c.Ctx, c.DB, boil.Infer())
	if err != nil {
		updateFail := common.BadRequest.New("Campaign not found")
		_ = common.AddErrorContext(updateFail, "id", strconv.Itoa(id))
		return id, updateFail
	}

	return id, nil
}

func (c CampaignService) DeleteCampaign(id int) error {
	campaign, err := models.FindCampaign(c.Ctx, c.DB, id)
	if err != nil {
		return common.NotFound.New("User not found")
	}

	_, err = campaign.Delete(c.Ctx, c.DB)
	if err != nil {
		deleteFail := common.BadRequest.New("Campaign not found")
		_ = common.AddErrorContext(deleteFail, "id", strconv.Itoa(id))
		return deleteFail
	}
	return nil
}
