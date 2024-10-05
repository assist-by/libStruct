package service

import "time"

type Service struct {
	Name          string `json:"name"`
	Address       string `json:"address"`
	LastHeartbeat time.Time
}
