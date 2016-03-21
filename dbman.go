package goapp

import (
	"github.com/go-ozzo/ozzo-dbx"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
)

type DBMan struct {
	write *dbx.DB
	read  []*dbx.DB
}

type DBConfig struct {
	driver string
	dsn    string
}

type DBManConfig struct {
	write *DBConfig
	read  []*DBConfig
}

func NewDBMan(config *DBManConfig) *DBMan {
	var err error
	dbman := new(DBMan)
	dbman.write, err = dbx.MustOpen(config.write.driver, config.write.dsn)
	dbman.write.LogFunc = log.Printf
	if err != nil {
		panic(err)
	}
	dbman.read = make([]*dbx.DB, 0)
	for _, c := range config.read {
		db, _ := dbx.MustOpen(c.driver, c.dsn)
		db.LogFunc = log.Printf
		dbman.read = append(dbman.read, db)
	}
	return dbman
}

func (d *DBMan) Write() *dbx.DB {
	return d.write
}

func (d *DBMan) Read() *dbx.DB {
	return d.read[rand.Intn(len(d.read))]
}
