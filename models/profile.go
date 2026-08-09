package models

type Profile struct {
	ID                uint   `json:"id" gorm:"primary_key"`
	Summary           string `json:"summary"`
	YearsOfExperience int    `json:"years_of_experience"`
}
