package models

import "gorm.io/gorm"

type Experience struct {
	ID               uint     `json:"id" gorm:"primary_key"`
	CompanyName      string   `json:"company_name"`
	Role             string   `json:"role"`
	StartDate        string   `json:"start_date"` // format: YYYY-MM
	EndDate          string   `json:"end_date"`   // format: YYYY-MM, empty if current
	IsCurrent        bool     `json:"is_current"`
	Responsibilities []string `json:"responsibilities" gorm:"serializer:json"`
}

func (p *Experience) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == 0 {
		var maxId int64
		tx.Model(&Experience{}).Select("IFNULL(MAX(id), 0)").Scan(&maxId)
		p.ID = uint(maxId + 1)
	}
	return
}
