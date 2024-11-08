package model

type Book struct {
	Id          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `gorm:"type:varchar(100);uniqueIndex" json:"name" binding:"required"`
	Description string  `gorm:"type:text" json:"description" binding:"required"`
	Price       float64 `gorm:"type:decimal(10,2)" json:"price" binding:"required"`
}
