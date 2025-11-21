package dtos

import "time"

type CreateCampaignRequest struct {
	TemplateID string    `json:"template_id" validate:"required,uuid"`
	City       string    `json:"city" validate:"required"`
	Slots      []string  `json:"slots" validate:"required,min=1"`
	Days       []string  `json:"days" validate:"required,min=1"`
	SourceURL  string    `json:"source_url" validate:"required,url"`
	StartDate  time.Time `json:"start_date" validate:"required"`
	EndDate    time.Time `json:"end_date" validate:"required,gtfield=StartDate"`
}

type CampaignResponse struct {
	ID         string    `json:"_id"`
	TemplateID string    `json:"template_id"`
	City       string    `json:"city"`
	Slots      []string  `json:"slots"`
	Days       []string  `json:"days"`
	SourceURL  string    `json:"source_url"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
