package models

type Url struct {
	ID int32 `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	ArticleID int32 `json:"articleid"`
	Label string `json:"label"`
	Link string `json:"link"`
}