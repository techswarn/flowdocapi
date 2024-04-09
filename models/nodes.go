package models

type Node struct {
	ID 	string `json:"id"  gorm:"unique"`
	Type string `json:"type"`
	Label string `json:"subject"`
}