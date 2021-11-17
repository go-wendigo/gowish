package model

import uuid "github.com/satori/go.uuid"

type Wish struct {
	ID        uuid.UUID
	Title     string
	ImageLink string
	Booked    bool
}
