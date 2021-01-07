package models

type ApplicationSettings struct {
	Id       int    `db:"id" form:"id" json:"id"`
	CycleSec      int `db:"cycleSec" form:"cycleSec" json:"cycleSec"`
}