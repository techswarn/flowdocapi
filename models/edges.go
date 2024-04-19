package models

import ("time")

type Edge struct {
	ID 					string `json:"id"  gorm:"unique"`
	Source 				string `json:"source"`
	Target 				string `json:"target"`
	Type				string `json:"type"`
	Animated            bool   `json:"animated"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}