package models

type Record struct {
	Email         string `json:"email" db:"email"`
	CName         string `json:"cname" db:"cname"`
	DiseaseCode   string `json:"disease_code" db:"disease_code"`
	TotalDeaths   int    `json:"total_deaths" db:"total_deaths"`
	TotalPatients int    `json:"total_patients" db:"total_patients"`
}
