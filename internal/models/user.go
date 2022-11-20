package models

type PublicServant struct {
	Email      string `json:"email" db:"email"`
	Name       string `json:"name" db:"name"`
	Surname    string `json:"surname" db:"name"`
	Salary     int    `json:"salary" db:"salary"`
	Phone      string `json:"phone" db:"phone"`
	CName      string `json:"cname" db:"cname"`
	Department string `json:"department" db:"department"`
}
