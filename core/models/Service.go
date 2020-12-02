package models

type Service struct {
	id       int    `db:"id" form:"id" json:"id"`
	url    string `db:"url" form:"url" json:"url"`
	name   string `db:"name" form:"name" json:"name"`
	status int    `db:"status" form:"status" json:"status"`
	createDt string `db:"create_dt" form:"createDt" json:"createDt"`
	updateDt string `db:"update_dt" form:"updateDt" json:"updateDt"`
}
