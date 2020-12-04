package db

import (
	"regexp"

	_ "github.com/mutecomm/go-sqlcipher/v4"
	"github.com/pkg/errors"
	"github.com/xjj1/StorageReporter/repository"
)

func (r *repo) AddEmailSettings(param []string) error {
	if len(param) < 2 {
		return errors.New("Please specify at least recipient and mailserver!")
	}
	if len(param) < 6 {
		for i := len(param); i < 6; i++ {
			param = append(param, "")
		}
	}

	if !isEmailValid(param[0]) {
		return errors.New("Invalid email, not changing settings")
	}

	// first delete previous settings
	_, err := r.db.Exec("delete from email;")
	if err != nil {
		return errors.Wrap(err, "delete from email")
	}

	// convert param to []interface{} or it will not work
	tmp := make([]interface{}, len(param))
	for i, v := range param {
		tmp[i] = v
	}
	_, err = r.db.Exec(`insert into email(
				rcptto,
				mailserver,
				mailfrom,
				subject,
				username,
				password
			)
		values (?, ?, ?, ?, ?, ?)`, tmp...)

	if err != nil {
		return errors.Wrap(err, "inserting email")
	}

	return nil
}

func (r *repo) GetEmail() (*repository.Email, error) {
	rows, err := r.db.Query("select rcptto,mailserver,mailfrom,subject,username,password from email")
	if err != nil {
		return nil, errors.Wrap(err, "select email")
	}
	defer rows.Close()

	var m repository.Email

	for rows.Next() {
		_ = rows.Scan(&m.Rcptto, &m.Mailserver, &m.Mailfrom, &m.Subject, &m.Username, &m.Password)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "select email scan result")
	}
	if m.Rcptto == "" {
		return nil, errors.New("No email configured")
	}
	return &m, nil
}

// isEmailValid checks if the email provided passes the required structure and length.
func isEmailValid(e string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}
