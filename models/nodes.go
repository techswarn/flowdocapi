package models

import ( "time")

type Node struct {
	ID 	string `json:"id"  gorm:"type:varchar(100);unique"`
	NodeType string `json:"nodetype"`  
	Label string `json:"subject"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Article Article `json:"article" gorm:"foreignKey:NodeID"`
}