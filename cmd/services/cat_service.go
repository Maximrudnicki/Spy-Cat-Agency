package services

import (
	"errors"
	"log"
	"test_rudnytskyi/cmd/data/request"
	"test_rudnytskyi/cmd/data/response"
	"test_rudnytskyi/cmd/models"
	u "test_rudnytskyi/cmd/utils"

	"test_rudnytskyi/cmd/repositories"
)

type CatService interface {
	Create(req request.CreateCatRequest) error
	Update(req request.UpdateCatRequest) error
	Delete(catId int) error
	FindById(catId int) (response.CatResponse, error)
	FindAll() ([]response.CatResponse, error)
}

type CatServiceImpl struct {
	CatRepository repositories.CatRepository
}

// Create implements CatsService.
func (c *CatServiceImpl) Create(req request.CreateCatRequest) error {
	if u.ValidateBreed(req.Breed) {
		newCat := models.Cat{
			Name:       req.Name,
			Experience: req.Experience,
			Breed:      req.Breed,
			Salary:     req.Salary,
		}

		err := c.CatRepository.Save(newCat)
		if err != nil {
			return err
		}

		return nil
	} else {
		log.Printf("Service: cannot create cat")
		return errors.New("service: cannot create cat, breed invalid")
	}
}

// Delete implements CatsService.
func (c *CatServiceImpl) Delete(catId int) error {
	err := c.CatRepository.Delete(catId)
	if err != nil {
		log.Printf("Service: cannot delete cat")
		return err
	} else {
		return nil
	}
}

// FindAll implements CatsService.
func (c *CatServiceImpl) FindAll() ([]response.CatResponse, error) {
	result, err := c.CatRepository.GetAll()
	if err != nil {
		log.Printf("Service: cannot find cats")
		return nil, err
	}

	var cats []response.CatResponse
	for _, cat := range result {
		cr := response.CatResponse{
			Id:         cat.CatId,
			Name:       cat.Name,
			Experience: cat.Experience,
			Breed:      cat.Breed,
			Salary:     cat.Salary,
		}
		cats = append(cats, cr)
	}

	return cats, nil
}

// FindById implements CatsService.
func (c *CatServiceImpl) FindById(catId int) (response.CatResponse, error) {
	cat, err := c.CatRepository.Get(catId)
	if err != nil {
		log.Printf("Service: cannot find cat")
		return response.CatResponse{}, err
	}

	catResponse := response.CatResponse{
		Id:         cat.CatId,
		Name:       cat.Name,
		Experience: cat.Experience,
		Breed:      cat.Breed,
		Salary:     cat.Salary,
	}
	return catResponse, nil
}

// Update implements CatsService.
func (c *CatServiceImpl) Update(req request.UpdateCatRequest) error {
	if u.ValidateBreed(req.Breed) {
		updatedCat := models.Cat{
			CatId:      req.Id,
			Name:       req.Name,
			Experience: req.Experience,
			Breed:      req.Breed,
			Salary:     req.Salary,
		}

		err := c.CatRepository.Update(updatedCat)
		if err != nil {
			log.Printf("Service: cannot update cat")
			return err
		}
		return nil
	} else {
		log.Printf("Service: cannot create cat")
		return errors.New("service: cannot create cat, breed invalid")
	}
}

// Constructor
func NewCatServiceImpl(catRepository repositories.CatRepository) CatService {
	return &CatServiceImpl{CatRepository: catRepository}
}
