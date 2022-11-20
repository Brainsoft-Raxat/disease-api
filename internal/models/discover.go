package models

import "time"

type Discover struct {
	CName        string     `json:"cname" db:"cname"`
	DiseaseCode  string     `json:"disease_code" db:"disease_code"`
	FirstEncDate *time.Time `json:"first_enc_date" db:"first_enc_date"`
}
