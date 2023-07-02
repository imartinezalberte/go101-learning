// It will print varl, init, init2, init3
//
// If there are more than one init function inside one package, Go will follow alphabetical order.
package mistakes

import (
	"database/sql"
	"fmt"
	"os"
)

type (
	DBExecutioner interface {
		Ping() error
	}

	DBExecutionDef struct {}

	DBConn struct {
		actualDB DBExecutioner
	}
)

var GlobalDBConn *DBConn

var a = func() int {
	fmt.Println("var")
	return 0
}

func init() {
	fmt.Println("init")
}

func init() {
	fmt.Println("init2")
}

func init() {
	fmt.Println("init3")
}

func init() {
	GlobalDBConn = &DBConn{
		actualDB: DBExecutionDef{},
	}
}

func DatabasePool() {
	const mysqlDataSourceName = "MYSQL_DATA_SOURCE_NAME"
	var globalDBConn *sql.DB

	// init function, it's only executed once
	func() {
		db, err := sql.Open("mysql", os.Getenv(mysqlDataSourceName))
		if err != nil {
			panic(err)
		}

		if err := db.Ping(); err != nil {
			panic(err)
		}

		globalDBConn = db
	}()

	// Creating a function is better than using init function to modify a global variable
	createDBClient := func(dataSourceName string) (*sql.DB, error) {
		db, err := sql.Open("mysql", dataSourceName)
		if err != nil {
			return nil, err
		}
		return db, db.Ping()
	}

	_, _ = createDBClient("")
	_ = globalDBConn
}

func (DBExecutionDef) Ping() error {
	return nil
}

func (d DBConn) Execute() error {
	if err := d.actualDB.Ping(); err != nil {
		return err
	}

	fmt.Println("Executing")
	return nil
}
