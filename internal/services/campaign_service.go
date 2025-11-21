package services

import (
	"context"

	"github.com/RealImage/team-48gb/internal/dao"
	"github.com/RealImage/team-48gb/internal/dtos"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// CampaignService defines the interface for campaign-related operations
type CampaignService struct {
	campaignDAO *dao.CampaignDAO
}

func NewCampaignService(db *mongo.Database) *CampaignService {
	campaignDAO := dao.NewCampaignDAO(db)
	return &CampaignService{
		campaignDAO: campaignDAO,
	}
}

func (s *CampaignService) CreateCampaign(ctx context.Context, req dtos.CreateCampaignRequest) (*dtos.CampaignResponse, error) {
	// Call DAO to create campaign in MongoDB
	campaign, err := s.campaignDAO.CreateCampaign(ctx, req)
	if err != nil {
		return nil, err
	}

	return campaign, nil
}
