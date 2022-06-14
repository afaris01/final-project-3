package responses

import "time"

type CreateResponse struct {
	ID        int       `json:"id"`
	Type      string    `json:"types"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateResponse struct {
	ID        int       `json:"id"`
	Type      string    `json:"types"`
	UpdatedAt time.Time `json:"updated_at"`
}
