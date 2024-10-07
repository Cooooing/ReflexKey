package model

type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt LocalTime
	UpdatedAt LocalTime
}
