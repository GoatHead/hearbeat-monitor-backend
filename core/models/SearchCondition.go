package models

type SearchCondition struct {
	Name       string    `form:"name" json:"name"`
	Url   string `form:"url" json:"url"`
	Status    int `form:"status" json:"status"`
	StartDate string `form:"startDate" json:"startDate"`
	EndDate string `form:"endDate" json:"endDate"`
	PageStart int `form:"pageStart" json:"pageStart"`
	PageSize int `form:"pageSize" json:"pageSize"`
}
