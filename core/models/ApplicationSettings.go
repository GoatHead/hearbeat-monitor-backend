package models

type ApplicationSettings struct {
	Id       int    `db:"id" form:"id" json:"id"`
	CycleSec      string `db:"cycleSec" form:"cycleSec" json:"cycleSec"`
}