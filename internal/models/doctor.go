package models

type Doctor struct {
	Email         string `json:"email" db:"email"`
	Name          string `json:"name" db:"name"`
	Surname       string `json:"surname" db:"name"`
	Salary        int    `json:"salary" db:"salary"`
	Phone         string `json:"phone" db:"phone"`
	CName         string `json:"cname" db:"cname"`
	Degree        string `json:"degree" db:"degree"`
	DiseaseTypeID int    `json:"id" db:"id"`
	DiseaseType   string `json:"disease_type"`
}
