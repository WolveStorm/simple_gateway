package dto

type PanelGroupData struct {
	ServiceNum      int `json:"service_num"`
	TodayRequestNum int `json:"today_request_num"`
	CurrentQps      int `json:"current_qps"`
	AppNum          int `json:"app_num"`
}

type DashboardStatOutput struct {
	Today     []int32 `json:"today"`
	Yesterday []int32 `json:"yesterday"`
}
