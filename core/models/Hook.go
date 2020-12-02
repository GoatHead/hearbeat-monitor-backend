package models

type Hook struct {
	Id       int    `db:"id" form:"id" json:"id"`
	Url      string `db:"url" form:"url" json:"url"`
	Name     string `db:"name" form:"name" json:"name"`
	CreateDt string `db:"create_dt" form:"createDt" json:"createDt"`
	UpdateDt string `db:"update_dt" form:"updateDt" json:"updateDt"`
}