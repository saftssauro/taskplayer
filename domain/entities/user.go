package entities

type User struct {
	Id        string   `json:"id"`
	CreatedAt uint64   `json:"createdAt"`
	UpdatedAt uint64   `json:"updatedAt"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Reports   []Report `json:"reports"`
}
