package entities

type Admin struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}