package models

/*
1) Mission consists of spying on targets and collecting data

2) One cat can only have one mission at a time, and a mission assumes a range of targets (minimum: 1, maximum: 3)

3) Cats should be able to share the collected data into the system by writing notes on a specific target.
Cats will be updating their notes from time to time and eventually mark the target as complete.
If the target is complete, notes should be frozen, i.e. cats should not be able to update them in any way.
After completing all of the targets, the mission is marked as completed.
*/
type Mission struct {
	ID          int    `gorm:"type:int;primary_key"`
	Name        string `gorm:"type:varchar;not null"`
	CatId       int
	IsCompleted bool     `gorm:"default:false"`
	Targets     []Target `gorm:"constraint:OnDelete:CASCADE;"`
}
