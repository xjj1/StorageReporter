package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mutecomm/go-sqlcipher/v4"
	"github.com/pkg/errors"
)

const fname = "data.db"

func InitDB() (*sql.DB, error) {
	var err error
	key := "2DD29CA851E7B56E4697B0E1F08507293D761A05CE4D1B628663F411A8086D99"
	dbname := fmt.Sprintf("%s?_pragma_key=x'%s'&_pragma_cipher_page_size=4096", fname, key)
	DBCon, err := sql.Open("sqlite3", dbname)
	if err != nil {
		return nil, errors.Wrap(err, "sql")
	}
	DBCon.SetMaxIdleConns(100)

	sqlStmt := `create table if not exists Arrays (
		ArrayType integer,
		Cluster varchar(1024),
		Name varchar(1024),
		Friendlyname varchar(1024),
		Username varchar(1024),
		Password varchar(1024)
	);`

	if _, err = DBCon.Exec(sqlStmt); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("%q: %s\n", err, sqlStmt))
		//return
	}

	sqlStmt = `create table if not exists email (
		rcpt_to varchar(1024),
		mailserver varchar(1024),
		mailfrom varchar(1024),
		subject varchar(1024),
		username varchar(1024),
		password varchar(1024)
	);`

	if _, err = DBCon.Exec(sqlStmt); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("%q: %s\n", err, sqlStmt))
		//return
	}

	sqlStmt = `create table if not exists history (
		datetime integer,
		array varchar(1024),
		disktype varchar(3),
		allsize integer,
		freesize integer,
		est_free_size integer,
		used_perc integer,
		snapshots integer,
		presented_size integer
	);
	`

	_, err = DBCon.Exec(sqlStmt)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("%q: %s\n", err, sqlStmt))
	}
	return DBCon, nil
}
