package services

import (
	"log"
	"test_rudnytskyi/cmd/data/request"
	"test_rudnytskyi/cmd/data/response"
	"test_rudnytskyi/cmd/models"

	"test_rudnytskyi/cmd/repositories"
)

type MissionService interface {
	Create(cmr request.CreateMissionRequest) error
	UpdateNameRequest(umr request.UpdateNameMissionRequest) error
	Delete(missionId uint32) error
	FindById(missionId uint32) (response.MissionResponse, error)
	FindAll() ([]response.MissionResponse, error)
	FindMissionByCatId(catId uint32) ([]response.MissionResponse, error)
	AssignCatToMission(catId uint32, missionId uint32) error
	CompleteMission(missionId uint32) error
	CompleteTarget(missionId uint32, targetId uint32) error
	AddTarget(atr request.AddTargetRequest) error
	RemoveTarget(targetId uint32, missionId uint32) error
	UpdateNotes(unr request.UpdateNotesRequest) error
	UpdateTarget(utr request.UpdateTargetRequest) error
}

type MissionServiceImpl struct {
	MissionRepository repositories.MissionRepository
}

// AddTarget implements MissionService.
func (m *MissionServiceImpl) AddTarget(atr request.AddTargetRequest) error {
	newTarget := models.Target{
		Name:      atr.Name,
		Country:   atr.Country,
		Notes:     atr.Notes,
		MissionId: atr.MissionId,
	}

	err := m.MissionRepository.AddTarget(newTarget, uint32(atr.MissionId))
	if err != nil {
		return err
	}

	return nil
}

// AssignCatToMission implements MissionService.
func (m *MissionServiceImpl) AssignCatToMission(catId uint32, missionId uint32) error {
	err := m.MissionRepository.AssignCatToMission(catId, missionId)
	if err != nil {
		return err
	}

	return nil
}

// CompleteMission implements MissionService.
func (m *MissionServiceImpl) CompleteMission(missionId uint32) error {
	err := m.MissionRepository.CompleteMission(missionId)
	if err != nil {
		return err
	}

	return nil
}

// CompleteTarget implements MissionService.
func (m *MissionServiceImpl) CompleteTarget(missionId uint32, targetId uint32) error {
	err := m.MissionRepository.CompleteTarget(missionId, targetId)
	if err != nil {
		return err
	}

	return nil
}

// Create implements MissionService.
func (m *MissionServiceImpl) Create(cmr request.CreateMissionRequest) error {
	newMission := models.Mission{
		Name: cmr.Name,
	}

	err := m.MissionRepository.Save(newMission)
	if err != nil {
		return err
	}

	return nil
}

// Delete implements MissionService.
func (m *MissionServiceImpl) Delete(missionId uint32) error {
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
				Id:          uint32(target.ID),
				Name:        target.Name,
				Country:     target.Country,
				Notes:       target.Notes,
				IsCompleted: target.IsCompleted,
			}

			targets = append(targets, targetResp)
		}

		mr := response.MissionResponse{
			Id:          uint32(mission.ID),
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
func (m *MissionServiceImpl) FindById(missionId uint32) (response.MissionResponse, error) {
	result, err := m.MissionRepository.Get(missionId)
	if err != nil {
		log.Printf("Service: cannot find missions")
		return response.MissionResponse{}, err
	}

	var targets []response.TargetResponse

	for _, target := range result.Targets {
		targetResp := response.TargetResponse{
			Id:          uint32(target.ID),
			Name:        target.Name,
			Country:     target.Country,
			Notes:       target.Notes,
			IsCompleted: target.IsCompleted,
		}
		targets = append(targets, targetResp)
	}

	mr := response.MissionResponse{
		Id:          uint32(result.ID),
		Name:        result.Name,
		CatId:       result.CatId,
		IsCompleted: result.IsCompleted,
		Targets:     targets,
	}
	
	return mr, nil
}

// FindMissionByCatId implements MissionService.
func (m *MissionServiceImpl) FindMissionByCatId(catId uint32) ([]response.MissionResponse, error) {
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
				Id:          uint32(target.ID),
				Name:        target.Name,
				Country:     target.Country,
				Notes:       target.Notes,
				IsCompleted: target.IsCompleted,
			}

			targets = append(targets, targetResp)
		}

		mr := response.MissionResponse{
			Id:          uint32(mission.ID),
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
func (m *MissionServiceImpl) RemoveTarget(targetId uint32, missionId uint32) error {
	err := m.MissionRepository.RemoveTarget(targetId, missionId)
	if err != nil {
		return err
	}

	return nil
}

// Update implements MissionService.
func (m *MissionServiceImpl) UpdateNameRequest(umr request.UpdateNameMissionRequest) error {
	updatedMission := models.Mission{
		Name: umr.Name,
	}

	err := m.MissionRepository.Update(updatedMission)
	if err != nil {
		log.Printf("Service: cannot update mission")
		return err
	}
	return nil
}

// UpdateNotes implements MissionService.
func (m *MissionServiceImpl) UpdateNotes(unr request.UpdateNotesRequest) error {
	err := m.MissionRepository.UpdateNotes(unr.Id, unr.Notes, uint32(unr.MissionId))
	if err != nil {
		log.Printf("Service: cannot update mission")
		return err
	}
	return nil
}

// UpdateTarget implements MissionService.
func (m *MissionServiceImpl) UpdateTarget(utr request.UpdateTargetRequest) error {
	updatedTarget := models.Target{
		Name:        utr.Name,
		Country:     utr.Country,
		Notes:       utr.Notes,
		IsCompleted: utr.IsCompleted,
		MissionId:   utr.MissionId,
	}

	err := m.MissionRepository.UpdateTarget(updatedTarget, utr.Id, uint32(utr.MissionId))
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
