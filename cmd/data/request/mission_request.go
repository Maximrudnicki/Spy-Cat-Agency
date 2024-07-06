package request

type CreateMissionRequest struct {
	Name string `json:"name"`
}

type UpdateNameMissionRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type AssignCatToMissionRequest struct {
	MissionId int `json:"mission_id"`
	CatId     int `json:"cat_id"`
}

type CompleteMissionRequest struct {
	Id          int  `json:"id"`
	IsCompleted bool `json:"is_completed"`
}

type UpdateTargetRequest struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Country     string `json:"country"`
	Notes       string `json:"notes"`
	IsCompleted bool   `json:"is_completed"`
	MissionId   int    `json:"mission_id"`
}

type UpdateNotesRequest struct {
	Id        int    `json:"id"`
	MissionId int    `json:"mission_id"`
	Notes     string `json:"notes"`
}

type CompleteTargetRequest struct {
	Id          int  `json:"id"`
	MissionId   int  `json:"mission_id"`
	IsCompleted bool `json:"is_completed"`
}

type AddTargetRequest struct {
	Name      string `json:"name"`
	Country   string `json:"country"`
	Notes     string `json:"notes"`
	MissionId int    `json:"mission_id"`
}

type FindMissionByCatId struct {
	CatId int `json:"cat_id"`
}

type RemoveTarget struct {
	TargetId  int `json:"target_id"`
	MissionId int `json:"mission_id"`
}
