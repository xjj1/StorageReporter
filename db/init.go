package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mutecomm/go-sqlcipher/v4"
	"github.com/pkg/errors"
)

const fname = "data.db"

func InitSQLiteRepo() (*Repository, error) {
	var err error
	key := "2DD29CA851E7B56E4697B0E1F08507293D761A05CE4D1B628663F411A8086D99"
	dbname := fmt.Sprintf("%s?_pragma_key=x'%s'&_pragma_cipher_page_size=4096", fname, key)
	dbConn, err := sql.Open("sqlite3", dbname)
	if err != nil {
		return nil, errors.Wrap(err, "sql")
	}
	dbConn.SetMaxIdleConns(100)

	// Name = ip address or fqdn
	// FriendlyName = if fqdn is not configured we need some kind of readable name
	sqlStmt := `create table if not exists Arrays (
		ArrayType integer,
		Cluster varchar(1024),
		Name varchar(1024),
		FriendlyName varchar(1024),
		Username varchar(1024),
		Password varchar(1024)
	);`

	if _, err = dbConn.Exec(sqlStmt); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("%q: %s\n", err, sqlStmt))
	}

	sqlStmt = `create table if not exists email (
		RcptTo varchar(1024),
		MailServer varchar(1024),
		MailFrom varchar(1024),
		Subject varchar(1024),
		Username varchar(1024),
		Password varchar(1024)
	);`

	if _, err = dbConn.Exec(sqlStmt); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("%q: %s\n", err, sqlStmt))
	}

	sqlStmt = `create table if not exists history (
		Datetime integer,
		Array varchar(1024),
		Disktype varchar(3),
		Allsize integer,
		Freesize integer,
		EstFreeSize integer,
		UsedPerc float,
		Snapshots integer,
		PresentedSize integer
	);
	`
	_, err = dbConn.Exec(sqlStmt)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("%q: %s\n", err, sqlStmt))
	}
	return &Repository{
		dbConn,
	}, nil
}
