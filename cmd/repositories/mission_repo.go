package repositories

import (
	"errors"
	"log"
	"test_rudnytskyi/cmd/models"

	"gorm.io/gorm"
)

type MissionRepository interface {
	Save(mission models.Mission) error
	Get(missionId uint32) (models.Mission, error)
	Update(mission models.Mission) error
	Delete(missionId uint32) error
	GetAll() ([]models.Mission, error)
	GetMissionByCatID(catId uint32) ([]models.Mission, error)
	AssignCatToMission(catId uint32, missionId uint32) error
	CompleteMission(missionId uint32) error

	CompleteTarget(missionId uint32, targetId uint32) error
	AddTarget(target models.Target, missionId uint32) error
	GetTargets(missionId uint32) ([]models.Target, error)
	GetTarget(targetId uint32, missionId uint32) (models.Target, error)
	RemoveTarget(targetId uint32, missionId uint32) error
	UpdateNotes(targetId uint32, note string, missionId uint32) error
	UpdateTarget(target models.Target, targetId uint32, missionId uint32) error
}

type MissionRepositoryImpl struct {
	Db *gorm.DB
}

// UpdateNotes implements MissionRepository.
func (m *MissionRepositoryImpl) UpdateNotes(targetId uint32, note string, missionId uint32) error {
	var tar models.Target
	tar, err := m.GetTarget(targetId, missionId)
	if err != nil {
		log.Printf("Repo: cannot complete target")
		return errors.New("cannot complete target")
	}

	if !tar.IsCompleted{
		target := models.Target{
			ID:          uint(targetId),
			MissionId:   int(missionId),
			Notes: note,
		}
	
		result := m.Db.Model(&target).Update("notes", note)
		if result.Error != nil {
			log.Printf("Repo: cannot complete target")
			return errors.New("cannot complete target")
		}
	
		return nil
	} else {
		log.Printf("Repo: cannot complete target")
		return errors.New("cannot complete target")
	}
}

// UpdateTarget implements MissionRepository.
func (m *MissionRepositoryImpl) UpdateTarget(target models.Target, targetId uint32, missionId uint32) error {
	var mission models.Mission
	result := m.Db.Where("id = ?", missionId).Find(&mission)
	if result.Error != nil {
		log.Printf("Repo: cannot update target")
		return errors.New("cannot update target")
	}

	tars, _ := m.GetTargets(missionId)
	mission.Targets = tars

	for _, t := range mission.Targets {
		if t.ID == uint(targetId) {
			t = target
		}
	}

	result = m.Db.Save(&mission)
	if result.Error != nil {
		log.Printf("Repo: cannot complete target")
		return errors.New("cannot complete target")
	}

	return nil
}

// RemoveTarget implements MissionRepository.
func (m *MissionRepositoryImpl) RemoveTarget(targetId uint32, missionId uint32) error {
	var target models.Target
	result := m.Db.Where("mission_id = ? AND id = ?", missionId, targetId).Delete(&target)

	if result.Error != nil {
		log.Printf("Repo: cannot remove target")
		return errors.New("cannot remove target")
	}
	return nil
}

// AddTarget implements MissionRepository.
func (m *MissionRepositoryImpl) AddTarget(target models.Target, missionId uint32) error {
	var mission models.Mission
	result := m.Db.Where("id = ?", missionId).Find(&mission)
	if result.Error != nil {
		log.Printf("Repo: cannot add target")
		return errors.New("cannot add target")
	}

	if !mission.IsCompleted {
		tars, _ := m.GetTargets(missionId)
		mission.Targets = tars

		if len(mission.Targets) < 3 {
			mission.Targets = append(mission.Targets, target)
		}

		result = m.Db.Save(&mission)
		if result.Error != nil {
			log.Printf("Repo: cannot add target")
			return errors.New("cannot add target")
		}

		return nil
	} else {
		log.Printf("Repo: cannot complete target, is_completed: %v", mission.IsCompleted)
		return errors.New("cannot complete target")
	}
}

// CompleteTarget implements MissionRepository.
func (m *MissionRepositoryImpl) CompleteTarget(missionId uint32, targetId uint32) error {
	var mission models.Mission
	result := m.Db.Where("id = ?", missionId).Find(&mission)
	if result.Error != nil {
		log.Printf("Repo: cannot complete target")
		return errors.New("cannot complete target")
	}

	tars, _ := m.GetTargets(missionId)
	mission.Targets = tars

	targets := mission.Targets
	var counter int
	for _, target := range targets {
		if target.IsCompleted {
			counter++
		}

		if target.ID == uint(targetId) {
			target.IsCompleted = true
		}
	}

	if counter == len(targets) {
		mission.IsCompleted = true

		result = m.Db.Save(&mission)
		if result.Error != nil {
			log.Printf("Repo: cannot complete target")
			return errors.New("cannot complete target")
		}

	}

	target := models.Target{
		ID:          uint(targetId),
		MissionId:   int(missionId),
		IsCompleted: true,
	}

	result = m.Db.Model(&target).Update("is_completed", true)
	if result.Error != nil {
		log.Printf("Repo: cannot complete target")
		return errors.New("cannot complete target")
	}

	return nil
}

// AssignCatToMission implements MissionRepository.
func (m *MissionRepositoryImpl) AssignCatToMission(catId uint32, missionId uint32) error {
	var mission models.Mission
	result := m.Db.Where("id = ?", missionId).Find(&mission)
	if result.Error != nil {
		log.Printf("Repo: cannot assign cat to the mission")
		return errors.New("cannot assign cat to the mission")
	}

	mission.CatId = catId

	result = m.Db.Save(&mission)
	if result.Error != nil {
		log.Printf("Repo: cannot assign cat to the mission")
		return errors.New("cannot assign cat to the mission")
	}

	return nil
}

// CompleteMission implements MissionRepository.
func (m *MissionRepositoryImpl) CompleteMission(missionId uint32) error {
	var mission models.Mission
	result := m.Db.Where("id = ?", missionId).Find(&mission)
	if result.Error != nil {
		log.Printf("Repo: cannot complete mission")
		return errors.New("cannot complete mission")
	}

	tars, _ := m.GetTargets(missionId)
	mission.Targets = tars

	mission.IsCompleted = true
	targets := mission.Targets
	for _, target := range targets {
		target.IsCompleted = true
	}

	result = m.Db.Save(&mission)
	if result.Error != nil {
		log.Printf("Repo: cannot complete mission")
		return errors.New("cannot complete mission")
	}

	return nil
}

// Delete implements MissionRepository.
func (m *MissionRepositoryImpl) Delete(missionId uint32) error {
	var mission models.Mission
	result := m.Db.Where("id = ?", missionId).Delete(&mission)
	if result.Error != nil {
		log.Printf("Repo: cannot delete mission")
		return errors.New("cannot delete mission")
	}
	return nil
}

// Get implements MissionRepository.
func (m *MissionRepositoryImpl) Get(missionId uint32) (models.Mission, error) {
	var mission models.Mission
	result := m.Db.Where("id = ?", missionId).Find(&mission)

	targets, err := m.GetTargets(missionId)
	if err != nil {
		return mission, nil
	}

	mission.Targets = targets

	if result != nil {
		return mission, nil
	} else {
		log.Printf("Repo: mission is not found")
		return mission, errors.New("mission is not found")
	}
}

// GetTarget implements MissionRepository.
func (m *MissionRepositoryImpl) GetTargets(missionId uint32) ([]models.Target, error) {
	var targets []models.Target
	result := m.Db.Where("mission_id = ?", missionId).Find(&targets)

	if result != nil {
		return targets, nil
	} else {
		log.Printf("Repo: target is not found")
		return targets, errors.New("target is not found")
	}
}

// GetTarget implements MissionRepository.
func (m *MissionRepositoryImpl) GetTarget(targetId uint32, missionId uint32) (models.Target, error) {
	var target models.Target
	result := m.Db.Where("mission_id = ? AND id = ?", missionId, targetId).Find(&target)

	if result != nil {
		return target, nil
	} else {
		log.Printf("Repo: target is not found")
		return target, errors.New("target is not found")
	}
}

// GetAll implements MissionRepository.
func (m *MissionRepositoryImpl) GetAll() ([]models.Mission, error) {
	var missions []models.Mission
	result := m.Db.Find(&missions)

	for i := range missions {
		targets, err := m.GetTargets(uint32(missions[i].ID))
		if err != nil {
			return missions, nil
		}

		missions[i].Targets = targets
	}

	if result.Error != nil {
		log.Printf("Repo: cannot get all mission")
		return nil, errors.New("cannot get all mission")
	}
	return missions, nil
}

// GetMissionByCatID implements MissionRepository.
func (m *MissionRepositoryImpl) GetMissionByCatID(catId uint32) ([]models.Mission, error) {
	var missions []models.Mission
	result := m.Db.Where("cat_id = ?", catId).Find(&missions)

	for i := range missions {
		targets, err := m.GetTargets(uint32(missions[i].ID))
		if err != nil {
			return missions, nil
		}

		missions[i].Targets = targets
	}

	if result != nil {
		return missions, nil
	} else {
		return nil, errors.New("mission is not found")
	}
}

// Save implements MissionRepository.
func (m *MissionRepositoryImpl) Save(mission models.Mission) error {
	result := m.Db.Create(&mission)
	if result.Error != nil {
		log.Printf("cannot save mission")
		return errors.New("cannot save mission")
	}
	return nil
}

// Update implements MissionRepository.
func (m *MissionRepositoryImpl) Update(mission models.Mission) error {
	var updatedMission = &models.Mission{
		CatId: mission.CatId,
		// Targets:     mission.Targets,
		IsCompleted: mission.IsCompleted,
	}

	result := m.Db.Model(&mission).Where("id = ?", mission.ID).Updates(updatedMission)
	if result.Error != nil {
		log.Printf("Repo: cannot update mission")
		return errors.New("cannot update mission")
	}
	return nil
}

// Constructor
func NewMissionRepositoryImpl(Db *gorm.DB) MissionRepository {
	return &MissionRepositoryImpl{Db: Db}
}
