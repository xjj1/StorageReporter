package repository

type Email struct {
	Rcptto, Mailserver, Mailfrom, Subject, Username, Password string
}

type Repository interface {
	AddEmailSettings(param []string) error
	GetEmail() (*Email, error)
	Close()
}

// func (r *Repository) Close() {
// 	r.db.Close()
// }
