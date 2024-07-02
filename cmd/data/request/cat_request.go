package request

type CreateCatRequest struct {
	Name       string `json:"name"`
	Experience uint32 `json:"experience"`
	Breed      string `json:"breed"`
	Salary     uint32 `json:"salary"`
}

type UpdateCatRequest struct {
	Id         uint32 `json:"id"`
	Name       string `json:"name"`
	Experience uint32 `json:"experience"`
	Breed      string `json:"breed"`
	Salary     uint32 `json:"salary"`
}
