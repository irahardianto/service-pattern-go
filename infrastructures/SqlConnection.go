package infrastructures

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// SQLconnection properties //
type SqlConnection struct {
	User     string
	Password string
	Host     string
	Port     string
	DbName   string
	Charset  string
	Db       *gorm.DB
}

// InitDB, initialize connection to database //
func (sql *SqlConnection) InitDB() error {

	var err error

	// open a db connection //
	sql.Db, err = gorm.Open("sqlite3", "/var/tmp/gorm.db")
	if err != nil {
		fmt.Println("Failed to connect database : ", err.Error())
	}
	sql.Db.LogMode(true)

	return err
}

// GetDB, get database connection //
func (sql *SqlConnection) GetDB() *gorm.DB {
	return sql.Db
}
