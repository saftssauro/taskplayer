package entities

type Report struct {
	Id         string `json:"id"`
	CreatedAt  uint64 `json:"createdAt"`
	UpdatedAt  uint64 `json:"updatedAt"`
	Name       string `json:"name"`
	Finished   bool   `json:"finished"`
	TaskAmount uint8  `json:"taskAmount"`
	TimeSpent  uint16 `json:"timeSpent"`
	User       User   `json:"user"`
	Tasks      []Task `json:"tasks"`
}
