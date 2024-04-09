package models

type Edge struct {
	ID 					int `json:"id"  gorm:"unique"`
	Source 				string `json:"source"`
	Target 				string `json:"target"`
	Type				string `json:"type"`
	Animated            bool   `json:"animated"`
}