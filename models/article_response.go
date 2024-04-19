package models

type ArticleResponse struct {
	ID int32 `json:"id"`
	NodeId string `json:"nodeid"`
	Heading string `json:"heading"`
	Description string `json:"description"`
	Error   string `json:"error"`
	Urls []Url `json:"urls" gorm:"foreignKey:ArticleID"`
}