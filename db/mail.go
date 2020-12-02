package db

import (
	_ "github.com/mutecomm/go-sqlcipher/v4"
	"github.com/pkg/errors"
)

func (r *Repository) AddEmail(param []string) error {
	if len(param) < 2 {
		return errors.New("Please specify at least recipient and mailserver!")
	}
	if len(param) < 6 {
		for i := len(param); i < 6; i++ {
			param = append(param, "")
		}
	}

	//stmt, err := r.db.Prepare("delete from email;")
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

	//log.Printf("%s %s added\n", param[0], param[1])

	return nil
}

type Email struct {
	Rcptto, Mailserver, Mailfrom, Subject, Username, Password string
}

func (r *Repository) GetEmail() (*Email, error) {
	rows, err := r.db.Query("select rcptto,mailserver,mailfrom,subject,username,password from email")
	if err != nil {
		return nil, errors.Wrap(err, "select email")
	}
	defer rows.Close()

	var m Email
	for rows.Next() {
		_ = rows.Scan(&m.Rcptto, &m.Mailserver, &m.Mailfrom, &m.Subject, &m.Username, &m.Password)
		// if err != nil {
		// 	return nil, errors.Wrap(err, "select email scan result")
		// }
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "select email scan result")
	}
	return &m, nil
}

/*
func listDevices() ([]Device, error) {
	var A []Device
	rows, err := DBCon.Query("select ArrayType, Cluster, Name, Friendlyname, Username, Password from Arrays")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var at DeviceType
		var Cluster, Name, Friendlyname, Username, Password string
		err = rows.Scan(&at, &Cluster, &Name, &Friendlyname, &Username, &Password)
		if err != nil {
			log.Fatalf("(listedevices) Error in DB %v", err)
		}
		var arr1, usr1, psw1 string
		arr1, err = Decrypt(Name)
		if err != nil {
			log.Fatalf("Cannot decrypt array name : %v", err)
		}
		usr1, err = Decrypt(Username)
		if err != nil {
			log.Fatalf("Cannot decrypt username : %v", err)
		}
		psw1, err = Decrypt(Password)
		if err != nil {
			log.Fatalf("Cannot decrypt password : %v", err)
		}
		A = append(A, Device{at, Cluster, arr1, Friendlyname, usr1, psw1})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return A, nil
}


func listarrays_names() ([]string, error) {
	A, err := listDevices()
	if err != nil {
		return []string{}, err
	}
	var ret []string
	for _, v := range A {
		ret = append(ret, v.Type.String()+` `+v.Name+` `+v.Friendlyname+` `+v.Cluster)
	}
	return ret, nil
}



func getemail() (rcpt_to, mailserver, mailfrom, subject, username, password string, err error) {
	rows, err := DBCon.Query("select rcpt_to,mailserver,mailfrom,subject,username,password from email")
	if err != nil {
		fmt.Println("Select failed")
		panic(err)
	}
	defer rows.Close()

	var rcpt_to1, mailserver1, mailfrom1, subject1, username1, password1 string
	for rows.Next() {
		err = rows.Scan(&rcpt_to, &mailserver, &mailfrom, &subject, &username, &password)
		if err != nil {
			log.Fatalln(err)
		}

		rcpt_to1, err = Decrypt(rcpt_to)
		if err != nil {
			log.Fatalln(err)
		}
		mailserver1, err = Decrypt(mailserver)
		if err != nil {
			log.Fatalln(err)
		}
		mailfrom1, err = Decrypt(mailfrom)
		if err != nil {
			log.Fatalln(err)
		}

		subject1, err = Decrypt(subject)
		if err != nil {
			log.Fatalln(err)
		}
		username1, err = Decrypt(username)
		if err != nil {

			log.Fatalln(err)
		}
		password1, err = Decrypt(password)
		if err != nil {
			log.Fatalln(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}
	return rcpt_to1, mailserver1, mailfrom1, subject1, username1, password1
}
*/
