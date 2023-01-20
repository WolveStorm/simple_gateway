package dto

type UserStatOutput struct {
	Name      string  `json:"name"`
	Today     []int32 `json:"today"`
	Yesterday []int32 `json:"yesterday"`
}
