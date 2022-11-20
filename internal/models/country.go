package models

type Country struct {
	Cname      string `json:"cname" db:"cname"`
	Population int64  `json:"population" db:"population"`
}
