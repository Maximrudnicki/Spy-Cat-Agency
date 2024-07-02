package response

type MissionResponse struct {
	Id          uint32 `json:"id"`
	Name        string `json:"name"`
	CatId       uint32 `json:"cat_id"`
	IsCompleted bool   `json:"is_completed"`
	Targets     []TargetResponse `json:"targets"`
}

type TargetResponse struct {
	Id          uint32 `json:"id"`
	Name        string `json:"name"`
	Country     string `json:"country"`
	Notes       string `json:"notes"`
	IsCompleted bool   `json:"is_completed"`
}
