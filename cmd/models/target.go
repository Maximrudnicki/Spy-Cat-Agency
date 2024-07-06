package models

// Targets are created in place along with a mission,
// meaning that there will be no page to see/create all/individual targets

type Target struct {
	ID          int    `gorm:"type:int;primary_key"`
	Name        string `gorm:"type:varchar;not null"`
	Country     string `gorm:"type:varchar;not null"`
	Notes       string `gorm:"type:varchar;not null"`
	IsCompleted bool   `gorm:"default:false"`
	MissionId   int
}
