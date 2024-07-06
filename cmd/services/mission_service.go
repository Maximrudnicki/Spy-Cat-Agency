package services

import (
	"log"
	"test_rudnytskyi/cmd/data/request"
	"test_rudnytskyi/cmd/data/response"
	"test_rudnytskyi/cmd/models"

	"test_rudnytskyi/cmd/repositories"
)

type MissionService interface {
	Create(req request.CreateMissionRequest) error
	UpdateNameRequest(req request.UpdateNameMissionRequest) error
	Delete(missionId int) error
	FindById(missionId int) (response.MissionResponse, error)
	FindAll() ([]response.MissionResponse, error)
	FindMissionByCatId(catId int) ([]response.MissionResponse, error)
	AssignCatToMission(catId int, missionId int) error
	CompleteMission(missionId int) error
	CompleteTarget(missionId int, targetId int) error
	AddTarget(req request.AddTargetRequest) error
	RemoveTarget(targetId int, missionId int) error
	UpdateNotes(req request.UpdateNotesRequest) error
	UpdateTarget(req request.UpdateTargetRequest) error
}

type MissionServiceImpl struct {
	MissionRepository repositories.MissionRepository
}

// AddTarget implements MissionService.
func (m *MissionServiceImpl) AddTarget(req request.AddTargetRequest) error {
	newTarget := models.Target{
		Name:      req.Name,
		Country:   req.Country,
		Notes:     req.Notes,
		MissionId: req.MissionId,
	}

	err := m.MissionRepository.AddTarget(newTarget, int(req.MissionId))
	if err != nil {
		return err
	}

	return nil
}

// AssignCatToMission implements MissionService.
func (m *MissionServiceImpl) AssignCatToMission(catId int, missionId int) error {
	err := m.MissionRepository.AssignCatToMission(catId, missionId)
	if err != nil {
		return err
	}

	return nil
}

// CompleteMission implements MissionService.
func (m *MissionServiceImpl) CompleteMission(missionId int) error {
	err := m.MissionRepository.CompleteMission(missionId)
	if err != nil {
		return err
	}

	return nil
}

// CompleteTarget implements MissionService.
func (m *MissionServiceImpl) CompleteTarget(missionId int, targetId int) error {
	err := m.MissionRepository.CompleteTarget(missionId, targetId)
	if err != nil {
		return err
	}

	return nil
}

// Create implements MissionService.
func (m *MissionServiceImpl) Create(req request.CreateMissionRequest) error {
	newMission := models.Mission{
		Name: req.Name,
	}

	err := m.MissionRepository.Save(newMission)
	if err != nil {
		return err
	}

	return nil
}

// Delete implements MissionService.
func (m *MissionServiceImpl) Delete(missionId int) error {
	err := m.MissionRepository.Delete(missionId)
	if err != nil {
		log.Printf("Service: cannot delete mission")
		return err
	} else {
		return nil
	}
}

// FindAll implements MissionService.
func (m *MissionServiceImpl) FindAll() ([]response.MissionResponse, error) {
	result, err := m.MissionRepository.GetAll()
	if err != nil {
		log.Printf("Service: cannot find missions")
		return nil, err
	}

	var missions []response.MissionResponse

	for _, mission := range result {
		var targets []response.TargetResponse

		for _, target := range mission.Targets {
			targetResp := response.TargetResponse{
				Id:          int(target.ID),
				Name:        target.Name,
				Country:     target.Country,
				Notes:       target.Notes,
				IsCompleted: target.IsCompleted,
			}

			targets = append(targets, targetResp)
		}

		mr := response.MissionResponse{
			Id:          int(mission.ID),
			Name:        mission.Name,
			CatId:       mission.CatId,
			IsCompleted: mission.IsCompleted,
			Targets:     targets,
		}
		missions = append(missions, mr)
	}

	return missions, nil
}

// FindById implements MissionService.
func (m *MissionServiceImpl) FindById(missionId int) (response.MissionResponse, error) {
	result, err := m.MissionRepository.Get(missionId)
	if err != nil {
		log.Printf("Service: cannot find missions")
		return response.MissionResponse{}, err
	}

	var targets []response.TargetResponse

	for _, target := range result.Targets {
		targetResp := response.TargetResponse{
			Id:          int(target.ID),
			Name:        target.Name,
			Country:     target.Country,
			Notes:       target.Notes,
			IsCompleted: target.IsCompleted,
		}
		targets = append(targets, targetResp)
	}

	mr := response.MissionResponse{
		Id:          int(result.ID),
		Name:        result.Name,
		CatId:       result.CatId,
		IsCompleted: result.IsCompleted,
		Targets:     targets,
	}

	return mr, nil
}

// FindMissionByCatId implements MissionService.
func (m *MissionServiceImpl) FindMissionByCatId(catId int) ([]response.MissionResponse, error) {
	result, err := m.MissionRepository.GetMissionByCatID(catId)
	if err != nil {
		log.Printf("Service: cannot find missions")
		return nil, err
	}

	var missions []response.MissionResponse

	for _, mission := range result {
		var targets []response.TargetResponse

		for _, target := range mission.Targets {
			targetResp := response.TargetResponse{
				Id:          int(target.ID),
				Name:        target.Name,
				Country:     target.Country,
				Notes:       target.Notes,
				IsCompleted: target.IsCompleted,
			}

			targets = append(targets, targetResp)
		}

		mr := response.MissionResponse{
			Id:          int(mission.ID),
			Name:        mission.Name,
			CatId:       mission.CatId,
			IsCompleted: mission.IsCompleted,
			Targets:     targets,
		}
		missions = append(missions, mr)
	}

	return missions, nil
}

// RemoveTarget implements MissionService.
func (m *MissionServiceImpl) RemoveTarget(targetId int, missionId int) error {
	err := m.MissionRepository.RemoveTarget(targetId, missionId)
	if err != nil {
		return err
	}

	return nil
}

// Update implements MissionService.
func (m *MissionServiceImpl) UpdateNameRequest(req request.UpdateNameMissionRequest) error {
	updatedMission := models.Mission{
		Name: req.Name,
	}

	err := m.MissionRepository.Update(updatedMission)
	if err != nil {
		log.Printf("Service: cannot update mission")
		return err
	}
	return nil
}

// UpdateNotes implements MissionService.
func (m *MissionServiceImpl) UpdateNotes(req request.UpdateNotesRequest) error {
	err := m.MissionRepository.UpdateNotes(req.Id, req.Notes, int(req.MissionId))
	if err != nil {
		log.Printf("Service: cannot update mission")
		return err
	}
	return nil
}

// UpdateTarget implements MissionService.
func (m *MissionServiceImpl) UpdateTarget(req request.UpdateTargetRequest) error {
	updatedTarget := models.Target{
		Name:        req.Name,
		Country:     req.Country,
		Notes:       req.Notes,
		IsCompleted: req.IsCompleted,
		MissionId:   req.MissionId,
	}

	err := m.MissionRepository.UpdateTarget(updatedTarget, req.Id, int(req.MissionId))
	if err != nil {
		log.Printf("Service: cannot update mission")
		return err
	}
	return nil
}

// Constructor
func NewMissionServiceImpl(missionRepository repositories.MissionRepository) MissionService {
	return &MissionServiceImpl{MissionRepository: missionRepository}
}
