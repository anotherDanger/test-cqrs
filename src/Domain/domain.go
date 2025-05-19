package domain

import "github.com/google/uuid"

type Domain struct {
	Id     uuid.UUID `json:"id"`
	Author string    `json:"author"`
	Title  string    `json:"title"`
	Genre  string    `json:"genre"`
}
