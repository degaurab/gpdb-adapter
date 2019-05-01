package gpdb_client

/*
* dbname - The name of the database to connect to
* user - The user to sign in as
* password - The user's password
* host - The host to connect to. Values that start with / are for unix
  domain sockets. (default is localhost)
* port - The port to bind to. (default is 5432)
* sslmode - Whether or not to use SSL (default is require, this is not
  the default for libpq)
* fallback_application_name - An application_name to fall back to if one isn't provided.
* connect_timeout - Maximum wait for connection, in seconds. Zero or
  not specified means wait indefinitely.
* sslcert - Cert file location. The file must contain PEM encoded data.
* sslkey - Key file location. The file must contain PEM encoded data.
* sslrootcert - The location of the root certificate file. The file
  must contain PEM encoded data.
 */
import (
	"database/sql"
	"fmt"
	"github.com/degaurab/gbdb-adapter/helper"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"log"
	"math/rand"
)


type DBDriver struct {
	User string
	Password string
	Port int
	Hostname string
	DatabaseName string
	SSLMode string
	SSLCertPath string
	SSLKeyPath string
	SSLRootCertPath string
	ConnectionTimeout int
}


type NewUser struct {
	UserName string `json:"user_name"`
	DBName string `json:"db_name"`
	Password string `json:"password"`
}


func (driver DBDriver) TestConnection(logger *log.Logger) error{
	connString := driver.createConnectionString()
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Println(err)
		return errors.New("Connection Failed")
	}

	rows, err := db.Query("SHOW DATABASES")
	if err != nil || rows.Err() != nil {
		log.Println(err, fmt.Sprint("Row error: %s", rows.Err()))
		return errors.New("Load connection query failed")
	}

	return nil

}

func (driver DBDriver) InitializeDBForUser(dbname string, username string, logger *log.Logger) (n NewUser, err error){
	n.DBName = dbname
	n.UserName = username

	connString := driver.createConnectionString()
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Println(err)
		return NewUser{}, errors.New("Connection Failed")
	}

	_, err = db.Query(fmt.Sprint("CREATE DATABSE %s", dbname))
	if err != nil{
		log.Println(err)
		return NewUser{}, errors.New("DB creation error")
	}

	n.Password = randStringBytes()

	_, err = db.Query(fmt.Sprint("CREATE USER %s with encrypted password '%s'", n.UserName, n.Password))
	if err != nil{
		log.Println(err)
		return NewUser{}, errors.New("Create user error")
	}

	/*
	 TODO: Additional grant permission that can be added for each user creation
	       this can be moved to seperate binary or rule engine later on

	GRANT ALL PRIVILEGES ON DATABASE yourdbname TO youruser;
	 */


	_, err = db.Query(fmt.Sprint(	"	GRANT ALL PRIVILEGES ON DATABASE %s TO %s", n.DBName, n.UserName))
	if err != nil{
		log.Println(err)
		return NewUser{}, errors.New("Granting access to user to database error")
	}
	return
}

func (driver DBDriver) DeleteDatabase(dbname string) error {
	connString := driver.createConnectionString()

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Println(err)
		return errors.New("Connection Failed")
	}

	row, err := db.Query(fmt.Sprint("SELECT * FROM pq_database WHERE datname='%s'", dbname))
	if err != nil || row.Next() {
		log.Println(err)
		return errors.New("Incorrect DB name")
	}

	_, err = db.Query(fmt.Sprint(	"DROP DATABASE %s", dbname))
	if err != nil{
		log.Println(err)
		return errors.New("Granting access to user to database error")
	}
	return nil
}




func (db DBDriver) createConnectionString() string{
	connString := fmt.Sprint("user=%s password=%s host=%s dbname=%s",db.User, db.Password, db.Hostname, db.DatabaseName)
	if db.SSLMode != "" {
		connString += connString + fmt.Sprint("sslmode=%s", db.SSLMode)
	}

	return connString
}


func randStringBytes() string {

	b := make([]byte, helper.RandStringLength)
	for i := range b {
		b[i] = helper.LetterBytes[rand.Intn(len(helper.LetterBytes))]
	}
	return string(b)
}