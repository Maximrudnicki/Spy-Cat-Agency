package repositories

import (
	"errors"
	"log"
	"test_rudnytskyi/cmd/models"

	"gorm.io/gorm"
)

// From agency perspective, they regularly hire new spy cats and
// so should be able to add them to  and visualize in the system
type CatRepository interface {
	Save(cat models.Cat) error
	Get(catId int) (models.Cat, error)
	Update(cat models.Cat) error
	Delete(catId int) error
	GetAll() ([]models.Cat, error)
}

type CatRepositoryImpl struct {
	Db *gorm.DB
}

// Delete implements CatRepository.
func (c *CatRepositoryImpl) Delete(catId int) error {
	var cat models.Cat
	result := c.Db.Where("cat_id = ?", catId).Delete(&cat)
	if result.Error != nil {
		log.Printf("Repo: cannot delete cat")
		return errors.New("cannot delete cat")
	}
	return nil
}

// Get implements CatRepository.
func (c *CatRepositoryImpl) Get(catId int) (models.Cat, error) {
	var cat models.Cat
	result := c.Db.Where("cat_id = ?", catId).Find(&cat)
	if result != nil {
		return cat, nil
	} else {
		log.Printf("Repo: cat is not found")
		return cat, errors.New("cat is not found")
	}
}

// GetAll implements CatRepository.
func (c *CatRepositoryImpl) GetAll() ([]models.Cat, error) {
	var cats []models.Cat
	result := c.Db.Find(&cats)
	if result.Error != nil {
		log.Printf("Repo: cannot get all cats")
		return nil, errors.New("cannot get all cats")
	}
	return cats, nil
}

// Save implements CatRepository.
func (c *CatRepositoryImpl) Save(cat models.Cat) error {
	result := c.Db.Create(&cat)
	if result.Error != nil {
		log.Printf("Repo: cannot save cat")
		return errors.New("cannot save cat")
	}
	return nil
}

// Update implements CatRepository.
func (c *CatRepositoryImpl) Update(cat models.Cat) error {
	var updatedCat = &models.Cat{
		Name:       cat.Name,
		Experience: cat.Experience,
		Breed:      cat.Breed,
		Salary:     cat.Salary,
	}

	result := c.Db.Model(&cat).Where("cat_id = ?", cat.CatId).Updates(updatedCat)
	if result.Error != nil {
		log.Printf("Repo: cannot update cat")
		return errors.New("cannot update cat")
	}
	return nil
}

// Constructor
func NewCatRepositoryImpl(Db *gorm.DB) CatRepository {
	return &CatRepositoryImpl{Db: Db}
}
