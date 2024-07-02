package models

// A cat is described as Name, Years of Experience, Breed, and Salary.
// Breed must be validated | Validate cat’s breed with [TheCatAPI](https://api.thecatapi.com/v1/breeds)
type Cat struct {
	CatId      uint32 `gorm:"type:int;primary_key"`
	Name       string `gorm:"type:varchar;not null"`
	Experience uint32 `gorm:"not null"`
	Breed      string `gorm:"type:varchar;not null"`
	Salary     uint32 `gorm:"not null"`
}
