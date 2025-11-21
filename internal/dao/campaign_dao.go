package dao

import (
	"context"
	"time"

	"github.com/RealImage/team-48gb/internal/dtos"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CampaignDAO struct {
	collection *mongo.Collection
}

func NewCampaignDAO(db *mongo.Database) *CampaignDAO {
	return &CampaignDAO{
		collection: db.Collection("campaigns"),
	}
}

func (d *CampaignDAO) CreateCampaign(ctx context.Context, req dtos.CreateCampaignRequest) (*dtos.CampaignResponse, error) {
	now := time.Now()

	// Prepare document for insertion
	campaign := bson.M{
		"template_id": req.TemplateID,
		"city":        req.City,
		"slots":       req.Slots,
		"days":        req.Days,
		"source_url":  req.SourceURL,
		"start_date":  req.StartDate,
		"end_date":    req.EndDate,
		"created_at":  now,
		"updated_at":  now,
	}

	// Insert into MongoDB
	result, err := d.collection.InsertOne(ctx, campaign)
	if err != nil {
		return nil, err
	}

	// Convert ObjectID to hex string
	objectID := result.InsertedID.(bson.ObjectID)

	// Return response
	return &dtos.CampaignResponse{
		ID:         objectID.Hex(),
		TemplateID: req.TemplateID,
		City:       req.City,
		Slots:      req.Slots,
		Days:       req.Days,
		SourceURL:  req.SourceURL,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		CreatedAt:  now,
		UpdatedAt:  now,
	}, nil
}
