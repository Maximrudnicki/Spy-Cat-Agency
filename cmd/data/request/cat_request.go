package request

type CreateCatRequest struct {
	Name       string `json:"name"`
	Experience int    `json:"experience"`
	Breed      string `json:"breed"`
	Salary     int    `json:"salary"`
}

type UpdateCatRequest struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Experience int    `json:"experience"`
	Breed      string `json:"breed"`
	Salary     int    `json:"salary"`
}
