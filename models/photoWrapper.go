package models

type PhotoWrapper struct {
	Photos []*Photo `json:"photos"`
}