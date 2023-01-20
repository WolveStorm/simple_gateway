package dto

type ServiceStatOutput struct {
	Info      ServiceInfo `json:"Info"`
	Today     []int32     `json:"today"`
	Yesterday []int32     `json:"yesterday"`
}

type ServiceInfo struct {
	ID uint `json:"id"`
}
