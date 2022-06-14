package inputs

type TaskCreateInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CategoryID  int    `json:"category_id"`
}

type TaskEditInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TaskUpdateStatus struct {
	Status bool `json:"status"`
}

type TaskUpdateCategory struct {
	CategoryID int `json:"category_id"`
}

type TaskID struct {
	ID int `uri:"id"`
}
