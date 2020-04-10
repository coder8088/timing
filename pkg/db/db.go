package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var SQLite *sql.DB

func init() {
	var err error
	SQLite, err = sql.Open("sqlite3", "./timing.db")
	if err != nil {
		log.Fatalf("Fail to open db './timing.db': %v", err)
	}
	//创建表
	createActionGroupTableSql := "CREATE TABLE IF NOT EXISTS action_group(id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE);"
	createActionTableSql := "CREATE TABLE IF NOT EXISTS action(id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE, group_id INT NOT NULL DEFAULT 0);"
	createTimingTableSql := `CREATE TABLE IF NOT EXISTS timing(id INTEGER PRIMARY KEY AUTOINCREMENT, action TEXT NOT NULL, duration_seconds INT NOT NULL DEFAULT 0, start_time DATE NOT NULL, stop_time DATE NULL, dt TEXT NOT NULL);`
	_, err = SQLite.Exec(createActionGroupTableSql)
	if err != nil {
		log.Fatalf("Fail to create table 'action_group': %v", err)
	}
	_, err = SQLite.Exec(createActionTableSql)
	if err != nil {
		log.Fatalf("Fail to create table 'action': %v", err)
	}
	_, err = SQLite.Exec(createTimingTableSql)
	if err != nil {
		log.Fatalf("Fail to create table 'timing': %v", err)
	}
}
