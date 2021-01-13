package database

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var firstRun = true

var defaultSchema = `
CREATE TABLE IF NOT EXISTS application_settings (
	id    INTEGER PRIMARY KEY autoincrement,
	cycleSec  INTEGER DEFAULT 300
);

INSERT INTO application_settings (cycleSec)
SELECT 300 WHERE NOT EXISTS(SELECT * FROM application_settings);

CREATE TABLE IF NOT EXISTS service (
	id    INTEGER PRIMARY KEY autoincrement,
	url  VARCHAR(1000) DEFAULT '',
    name VARCHAR(80)  DEFAULT '' UNIQUE ,
	status INTEGER DEFAULT NULL,
	create_dt datetime default (datetime('now', 'localtime')),
  	update_dt datetime default (datetime('now', 'localtime'))
);

CREATE TABLE IF NOT EXISTS service_history (
	id    INTEGER PRIMARY KEY autoincrement,
	url  VARCHAR(1000) DEFAULT '',
    name VARCHAR(80)  DEFAULT '',
	status INTEGER DEFAULT NULL,
	create_dt datetime default (datetime('now', 'localtime'))
);

CREATE TABLE IF NOT EXISTS hook (
	id    INTEGER PRIMARY KEY autoincrement,
	url  VARCHAR(1000) DEFAULT '',
	type  VARCHAR(10) DEFAULT 'MS_TEAMS',
    name VARCHAR(80)  DEFAULT '' UNIQUE ,
	create_dt datetime default (datetime('now', 'localtime')),
  	update_dt datetime default (datetime('now', 'localtime'))
);
`

func GetDbConnector() (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", "../db/_service_list.db")
	logger := gin.DefaultWriter
	errorLogger := gin.DefaultErrorWriter
	if err != nil {
		errorLogger.Write([]byte(err.Error()))
	}
	if firstRun {

		db.MustExec(defaultSchema)
		firstRun = false
		logger.Write([]byte("DB INITIALIZED\n"))

	}

	return db, err
}