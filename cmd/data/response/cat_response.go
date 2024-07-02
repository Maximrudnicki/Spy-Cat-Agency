package response

type CatResponse struct {
	Id         uint32 `json:"id"`
	Name       string `json:"name"`
	Experience uint32 `json:"experience"`
	Breed      string `json:"breed"`
	Salary     uint32 `json:"salary"`
}
