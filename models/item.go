package models

type Item struct {
	ID    uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
