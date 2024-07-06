package response

type MissionResponse struct {
	Id          int              `json:"id"`
	Name        string           `json:"name"`
	CatId       int              `json:"cat_id"`
	IsCompleted bool             `json:"is_completed"`
	Targets     []TargetResponse `json:"targets"`
}

type TargetResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Country     string `json:"country"`
	Notes       string `json:"notes"`
	IsCompleted bool   `json:"is_completed"`
}
