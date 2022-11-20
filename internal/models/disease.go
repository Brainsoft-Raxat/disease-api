package models

type DiseaseType struct {
	ID          int    `json:"id,omitempty" db:"id"`
	Description string `json:"description" db:"description"`
}

type Disease struct {
	Code          string `json:"code" db:"code"`
	Pathogen      string `json:"pathogen" db:"pathogen"`
	Description   string `json:"description" db:"description"`
	DiseaseTypeID int    `json:"id" db:"id"`
	DiseaseType   string `json:"disease_type"`
}
