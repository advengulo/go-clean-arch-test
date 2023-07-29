package models

import "time"

type Token struct {
	Token     string    `json:"access_token"`
	ExpiredAt time.Time `json:"expired_at"`
}
