package models

// A cat is described as Name, Years of Experience, Breed, and Salary.
// Breed must be validated | Validate cat’s breed with [TheCatAPI](https://api.thecatapi.com/v1/breeds)
type Cat struct {
	CatId      int    `gorm:"type:int;primary_key"`
	Name       string `gorm:"type:varchar;not null"`
	Experience int    `gorm:"not null"`
	Breed      string `gorm:"type:varchar;not null"`
	Salary     int    `gorm:"not null"`
}
