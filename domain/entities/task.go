package entities

type Task struct {
	Id          string `json:"id"`
	CreatedAt   uint64 `json:"createdAt"`
	UpdatedAt   uint64 `json:"updatedAt"`
	Name        string `json:"name"`
	Description string `json:"description"`
	FinishedAt  uint64 `json:"finishedAt"`
	Paused      bool   `json:"paused"`
	TimeSpent   uint16 `json:"timeSpent"`
	Report      Report `json:"report"`
}
