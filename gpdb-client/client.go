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
	"io/ioutil"
	"log"
	"math/rand"
	"path"
	"regexp"
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
	DBTemplate DBTemplate
}

type DBTemplate struct {
	TemplatePath string
	SchemaTemplateFile string
	UserTemplateFile string
}


type NewUser struct {
	UserName   string `json:"user_name"`
	SchemaName string `json:"schema_name"`
	Password   string `json:"password"`
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
		log.Println(err, fmt.Sprintf("Row error: %s", rows.Err()))
		return errors.New("Load connection query failed")
	}

	return nil

}

func (driver DBDriver) InitializeDBForUser(dbname string, username string, logger *log.Logger) (n NewUser, err error){
	n.SchemaName = dbname
	n.UserName = username

	connString := driver.createConnectionString()
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Println(err)
		return NewUser{}, errors.New("Connection Failed")
	}

	/*
		Creating user based on templates
	 */
	n.Password = randStringBytes()

	filePath := path.Join(driver.DBTemplate.TemplatePath, driver.DBTemplate.UserTemplateFile)
	varMap := map[string]string {
		"schema_username": n.UserName,
		"schema_user_password": n.Password,
	}
	queryString, err := renderFileWithVariables(filePath, varMap)
	if err != nil{
		log.Println(err)
		return NewUser{}, errors.New("User creation error")
	}

	logger.Println("creating user: ", queryString)
	_, err = db.Query(queryString)
	if err != nil{
		log.Println(err)
		return NewUser{}, errors.New("Create user error")
	}

	/*
		Creating schema and grant access associated with the user
	 */
	filePath = path.Join(driver.DBTemplate.TemplatePath, driver.DBTemplate.SchemaTemplateFile)
	varMap = map[string]string {
		"schema_username": n.UserName,
		"schema_name": n.SchemaName,
	}

	queryString, err = renderFileWithVariables(filePath, varMap)
	if err != nil{
		log.Println(err)
		return NewUser{}, errors.New("User creation error")
	}

	logger.Println("creating schema: ", queryString)
	_, err = db.Query(queryString)
	if err != nil{
		log.Println(err)
		return NewUser{}, errors.New("DB creation error")
	}
	
	return
}

func (driver DBDriver) DeleteDatabase(dbname string, logger *log.Logger) error {
	connString := driver.createConnectionString()

	db, err := sql.Open("postgres", connString)
	if err != nil {
		logger.Println(err)
		return errors.New("Connection Failed")
	}

	row, err := db.Query(fmt.Sprintf("SELECT * FROM pq_database WHERE datname='%s'", dbname))
	if err != nil || row.Next() {
		logger.Println(err)
		return errors.New(fmt.Sprintf("Incorrect binding_id: %s", dbname))
	}

	_, err = db.Query(fmt.Sprintf(	"DROP DATABASE %s", dbname))
	if err != nil{
		logger.Println(err)
		return errors.New("Dropping database error")
	}
	return nil
}




func (db DBDriver) createConnectionString() string{
	connString := fmt.Sprintf("user=%s password=%s host=%s",db.User, db.Password, db.Hostname)
	if db.DatabaseName != "" {
		connString += connString + fmt.Sprintf(" dbname=%s", db.DatabaseName)
	}
	if db.SSLMode != "" {
		connString += connString + fmt.Sprintf("sslmode=%s", db.SSLMode)
	}
	log.Println(connString)
	return connString
}


func randStringBytes() string {

	b := make([]byte, helper.RandStringLength)
	for i := range b {
		b[i] = helper.LetterBytes[rand.Intn(len(helper.LetterBytes))]
	}
	return string(b)
}

func renderFileWithVariables(templateFilePath string, vars map[string]string) (string, error) {
	templateBytes, err := ioutil.ReadFile(templateFilePath)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Loading template file: ", templateFilePath))
	}

	templateData := string(templateBytes)
	for stringName, replaceWith := range vars {
		regex := regexp.MustCompile(stringName)
		templateData = regex.ReplaceAllString(templateData, replaceWith)
	}

	return templateData, nil
}