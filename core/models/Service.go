package models

type Service struct {
	Id       int    `db:"id" form:"id" json:"id"`
	Name   string `db:"name" form:"name" json:"name"`
	Url    string `db:"url" form:"url" json:"url"`
	Status int    `db:"status" form:"status" json:"status"`
	CreateDt string `db:"create_dt" form:"createDt" json:"createDt"`
	UpdateDt string `db:"update_dt" form:"updateDt" json:"updateDt"`
}
