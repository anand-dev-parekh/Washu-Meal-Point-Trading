package models

type Report struct {
	ID          uint64 `json:"id"`
	ReportID    uint64 `json:"reportID"`
	UserID      uint64 `json:"userID"`
	ReportEmail string `json:"reportEmail"`
	UserEmail   string `json:"userEmail"`
	Message     string `json:"message"`
}
