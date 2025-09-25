package model

import "time"

type Password struct {
	ID int `json: "id"`
	Label string `json: "label"`
	Password string `json: "password"`
	CreatedAt time.Time `json: "createdAt"`
}