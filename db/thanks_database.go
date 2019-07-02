package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const dbfile = "./thanks.db"

type ThanksRepository struct {
	db *sql.DB
}

func NewThanksRepository() (*ThanksRepository, error) {
	db, err := initdb()
	if err != nil {
		return nil, err
	}
	return &ThanksRepository{db}, nil
}

func initdb() (*sql.DB, error) {
	_, err := os.Stat(dbfile)
	// exists database
	if err == nil {
		return sql.Open("sqlite3", dbfile)
	}
	// if not exist database
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	c := `CREATE TABLE THANKS(
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		SOURCE TEXT,
		DESTINATION TEXT,
		MESSAGE TEXT,
		CREATED_AT TIMESTAMP DEFAULT (DATETIME('now','localtime'))
	)`
	if _, err := db.Exec(c); err != nil {
		return nil, err
	}
	return db, nil
}

func (thanks *ThanksRepository) Close() {
	thanks.db.Close()
}

func (thanks *ThanksRepository) Save(message string, source string, destination string) error {
	sql := fmt.Sprintf("INSERT INTO THANKS(SOURCE, DESTINATION, MESSAGE) VALUES('%v', '%v', '%v')", source, destination, message)
	if _, err := thanks.db.Exec(sql); err != nil {
		return err
	}
	return nil
}

type Summary struct {
	Count       int    `json:"count"`
	Destination string `json:"destination"`
}

type Summaries []*Summary

func (thanks *ThanksRepository) Summary(diff int) (Summaries, error) {
	var summaries Summaries
	s := `select count(1) as count
,destination
from THANKS
where date(created_at) = date('now', '-%d days')
group by destination
order by count desc limit 2;`
	sql := fmt.Sprintf(s, diff)
	rows, err := thanks.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var count int
		var destination string
		if err := rows.Scan(&count, &destination); err != nil {
			return nil, err
		}
		summaries = append(summaries, &Summary{count, destination})
	}
	return summaries, nil
}
