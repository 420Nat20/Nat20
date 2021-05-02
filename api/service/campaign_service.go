package service

import (
	"github.com/420Nat20/Nat20/nat-20/data/models"
)

type CampaignService struct{}

func (c CampaignService) GetCampaign(id int) (models.Campaign, error) {
	panic("implement me")
}

func (c CampaignService) GetAllCampaigns() (models.CampaignSlice, error) {
	panic("implement me")
}

func (c CampaignService) CreateCampaign() (models.Campaign, error) {
	panic("implement me")
}

func (c CampaignService) UpdateCampaign(id int, campaign models.Campaign) (models.Campaign, error) {
	panic("implement me")
}

func (c CampaignService) DeleteCampaign(id int) error {
	panic("implement me")
}
