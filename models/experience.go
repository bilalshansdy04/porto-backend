package models

type Experience struct {
	ID               uint     `json:"id" gorm:"primary_key"`
	CompanyName      string   `json:"company_name"`
	Role             string   `json:"role"`
	StartDate        string   `json:"start_date"` // format: YYYY-MM
	EndDate          string   `json:"end_date"`   // format: YYYY-MM, empty if current
	IsCurrent        bool     `json:"is_current"`
	Responsibilities []string `json:"responsibilities" gorm:"serializer:json"`
}
