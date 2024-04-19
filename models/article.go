package models

type Article struct {
	ID int32 `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	NodeID string `json:"nodeid" gorm:"type:varchar(100)"`
	Heading string `json:"heading"`
	Description string `json:"description"`
	Error   string `json:"error"`
	Urls []Url `json:"urls" gorm:"foreignKey:ArticleID"`
}