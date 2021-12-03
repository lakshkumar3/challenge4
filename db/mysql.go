package db

import (
	"database/sql"
	"fmt"
	"github.com/apex/log"
	"github.com/cakemarketing/snowbank/stores"
	"github.com/cakemarketing/snowbank/stores/aurora"
	"sync"
)

type Mysql struct {
	*aurora.Aurora
	mu *sync.RWMutex
}

type EquationSaver interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

func (db *Mysql) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.DB.Exec(query, args...)

}
func Connect(host, port, database, user, password string) (*Mysql, error) {
	url := fmt.Sprintf("%s:%s", host, port)
	conn, err := aurora.NewConnection(url, user, password, database, 1000)

	if err != nil {
		log.Error(fmt.Sprintf("error while coonecting %v ", err))
		return nil, fmt.Errorf("error while getting connection")
	}

	mysqlDb := &Mysql{
		Aurora: conn,
		mu:     &sync.RWMutex{},
	}
	err = mysqlDb.checkSchema()
	if err != nil {
		return nil, err
	}
	return mysqlDb, nil
}
func (db *Mysql) checkSchema() error {
	ok := db.TableExists("Equations")
	if !ok {
		query := "CREATE TABLE Equations (    ID int AUTO_INCREMENT,    UserName varchar(255) NOT NULL,    Expression varchar(255) NOT NULL,    Result varchar(255) NOT NULL,    Eqtimestamp datetime Not Null,    PRIMARY KEY (ID))"
		result, err := db.Exec(query)
		if err != nil {
			log.Error(fmt.Sprintf("error at creating schema %v", err))
			return err
		}
		log.Info(fmt.Sprintf("result of query %v", result))
		return nil
	}
	return nil
}

func (db *Mysql) GetModuleType() stores.ModuleType {
	return stores.ModuleType_DB
}

//
func (db *Mysql) GetDatabaseType() stores.DatabaseType {
	return stores.DatabaseType_Aurora
}

func GetModuleType() stores.ModuleType {
	return stores.ModuleType_DB
}

func (db *Mysql) Healthy() error {
	return db.Healthy()
}
